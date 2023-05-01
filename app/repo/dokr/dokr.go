package dokr

import (
	"archive/tar"
	"context"
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
			r.runTask(context.Background(), tsk, resCh, rdr)
		}(def.Tasks[i], readers[i])
	}

	for i := range def.Tasks {
		log.Infof("waiting for task: %d\n", i)
		res[i] = <-resCh
	}

	return res
}

func (r *runner) runTask(ctx context.Context, task bo.UnitTask, resCh chan<- bo.TaskResult, inputStream io.Reader) {

	// log.Infof("preparing tasks - ", task.Name)
	// if err := r.containerManager.PullImage(ctx); err != nil {
	// 	log.Errorf("Pull IMAGE")
	// 	resCh <- bo.TaskFailed(task, 0, "PULL_IMAGE")
	// 	return
	// }

	log.Infof("creating task - ", task.Name)
	id, err := r.containerManager.CreateContainer(ctx, task, inputStream)
	if err != nil {
		log.Errorf("could not create: %v", err)
		resCh <- bo.TaskFailed(task, 1, "CREATE_CONTAINER")
		return
	}

	log.Infof("starting task - ", task.Name)
	if err = r.containerManager.StartContainer(ctx, id); err != nil {
		log.Errorf("could not start task: %v", err)
		resCh <- bo.TaskFailed(task, 2, "START_CONTAINER")
		return
	}

	defer r.containerManager.RemoveContainer(ctx, id)

	log.Infof("waiting task - ", task.Name)
	statusSucess, err := r.containerManager.WaitForContainer(ctx, id)
	if err != nil {
		log.Errorf("wait for failed: %v", err)
		resCh <- bo.TaskFailed(task, 3, "WAIT_CONTAINER")
		return
	}

	if !statusSucess {
		log.Errorf("task failed")
		resCh <- bo.TaskFailed(task, 3, "TASK_FAILED")
		return
	}

	log.Infof("sucessed task - ", task.Name)
	stream, err := r.containerManager.CopyTarOutput(ctx, id)
	if err != nil {
		log.Errorf("copy tar: %v", err)
		resCh <- bo.TaskFailed(task, 4, "COPY_TAR")
		return
	}

	tarStream := tar.NewReader(stream)
	output, err := getOutput(tarStream)
	if err != nil {
		resCh <- bo.TaskFailed(task, 4, "COPY_TAR")
		return
	}

	resCh <- bo.TaskResult{
		Name:   task.Name,
		Output: output,
		Error:  nil,
	}
}

func getOutput(tarStream *tar.Reader) (map[string]interface{}, error) {
	_, err := tarStream.Next()
	if err != nil {
		log.Errorf("stream next: %v", err)
		return nil, err
	}

	// data := make([]byte, header.Size)
	// if _, err = io.ReadFull(tarStream, data); err != nil {
	// 	log.Errorf("could not read: %v", err)
	// 	resCh <- bo.TaskFailed(task, 4, "READ_TAR")
	// }

	output, err := system.Decode[map[string]interface{}](tarStream)
	if err != nil {
		log.Errorf("could not decode json: %v", err)
		return nil, err
	}

	return output, nil
}
