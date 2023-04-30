package dokr

import (
	"archive/tar"
	"context"
	"fmt"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/labstack/gommon/log"
)

// func getReaders(source io.Reader, count int) ([]io.Reader, io.Closer) {
//   readers := make([]io.Reader, 0, count)
//   pipeWriters := make([]io.Writer, 0, count)
//   pipeClosers := make([]io.Closer, 0, count)
// 	for i := 0; i < count-1; i++ {
//
//     pr, pw := io.Pipe()
//     readers = append(readers, pr)
//     pipeWriters = append(pipeWriters, pw)
//     pipeClosers = append(pipeClosers, pw)
//   }  multiWriter := io.MultiWriter(pipeWriters...)
//
//   teeReader := io.TeeReader(source, multiWriter)  // append teereader so it populates data to the rest of the readers
//   readers = append([]io.Reader{teeReader}, readers...)
//   return readers, NewMultiCloser(pipeClosers)
// }

func (r *runner) Run(ctx context.Context, def bo.TaskDefinition, src io.Reader) []bo.TaskResult {

	res := make([]bo.TaskResult, len(def.Tasks))
	resCh := make(chan bo.TaskResult, len(def.Tasks))

	for _, task := range def.Tasks {
		go r.runTask(ctx, task, resCh)
	}

	for i := range def.Tasks {
		res[i] = <-resCh
	}

	return res
}

func (r *runner) runTask(ctx context.Context, task bo.UnitTask, resCh chan<- bo.TaskResult) {

	fmt.Println("preparing tasks - ", task.Name)
	if err := r.containerManager.PullImage(ctx, task.Runner); err != nil {
		resCh <- bo.TaskFailed(task, 0, "PULL_IMAGE")
		return
	}

	id, err := r.containerManager.CreateContainer(ctx, task)
	if err != nil {
		log.Errorf("creating task: %v", err)
		resCh <- bo.TaskFailed(task, 1, "CREATE_CONTAINER")
		return
	}

	fmt.Println("starting task - ", task.Name)
	if err = r.containerManager.StartContainer(ctx, id); err != nil {
		log.Errorf("starting task: %v", err)
		resCh <- bo.TaskFailed(task, 2, "START_CONTAINER")
		return
	}

	defer r.containerManager.RemoveContainer(ctx, id)

	fmt.Println("waiting task - ", task.Name)
	statusSucess, err := r.containerManager.WaitForContainer(ctx, id)
	if err != nil {
		log.Errorf("wait for task: %v", err)
		resCh <- bo.TaskFailed(task, 3, "WAIT_CONTAINER")
		return
	}

	if statusSucess {
		fmt.Println("sucessed task - ", task.Name)
		stream, err := r.containerManager.CopyTarOutput(ctx, id)

		if err != nil {
			log.Errorf("copy tar: %v", err)
			resCh <- bo.TaskFailed(task, 4, "COPY_TAR")
			return
		}

		defer stream.Close()

		tarStream := tar.NewReader(stream)
		header, err := tarStream.Next()
		if err != nil {
			log.Errorf("stream next: %v", err)
			resCh <- bo.TaskFailed(task, 4, "HEADER_TAR")
		}

		data := make([]byte, header.Size)
		if _, err = io.ReadFull(tarStream, data); err != nil {
			log.Errorf("could not read: %v", err)
			resCh <- bo.TaskFailed(task, 4, "READ_TAR")
		}

		resCh <- bo.TaskResult{
			Name:   task.Name,
			Output: string(data),
			Error:  nil,
		}
	}
}
