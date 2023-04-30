package dokr

import (
	"archive/tar"
	"context"
	"fmt"
	"io"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/gabrieldiaziv/wghidra/app/system"
	"github.com/labstack/gommon/log"
)

func (r *runner) Run(ctx context.Context, def bo.TaskDefinition) []bo.TaskResult {

	res := make([]bo.TaskResult, len(def.Tasks))
	resCh := make(chan bo.TaskResult, len(def.Tasks))
	readers := system.GetReaders(def.Exe, len(def.Tasks))

	for i := range def.Tasks {
		go func(tsk bo.UnitTask, rdr io.Reader) {
			r.runTask(ctx, tsk, resCh, rdr)
		}(def.Tasks[i], readers[i])
	}

	for i := range def.Tasks {
		res[i] = <-resCh
	}

	return res
}

func (r *runner) runTask(ctx context.Context, task bo.UnitTask, resCh chan<- bo.TaskResult, exeStream io.Reader) {

	fmt.Println("preparing tasks - ", task.Name)
	if err := r.containerManager.PullImage(ctx); err != nil {
		resCh <- bo.TaskFailed(task, 0, "PULL_IMAGE")
		return
	}

	id, err := r.containerManager.CreateContainer(ctx, task, exeStream)
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
