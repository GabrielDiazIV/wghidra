package dokr

import (
	"context"
	"fmt"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

func (r *runner) Run(ctx context.Context, def bo.TaskDefinition) []bo.TaskResult {

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
		fmt.Println(err)
		resCh <- bo.TaskFailed(task, 1, "CREATE_CONTAINER")
		return
	}

	fmt.Println("starting task - ", task.Name)
	if err = r.containerManager.StartContainer(ctx, id); err != nil {
		fmt.Println(err)
		resCh <- bo.TaskFailed(task, 2, "START_CONTAINER")
		return
	}

	defer r.containerManager.RemoveContainer(ctx, id)

	fmt.Println("waiting task - ", task.Name)
	statusSucess, err := r.containerManager.WaitForContainer(ctx, id)
	if err != nil {
		fmt.Println(err)
		resCh <- bo.TaskFailed(task, 3, "WAIT_CONTAINER")
		return
	}

	if statusSucess {
		fmt.Println("sucessed task - ", task.Name)
		stream, err := r.containerManager.CopyTarOutput(ctx, id)

		if err != nil {
			fmt.Println(err)
			resCh <- bo.TaskFailed(task, 4, "COPY_TAR")
			return
		}

		resCh <- bo.TaskResult{
			Name:      task.Name,
			TarStream: stream,
			Error:     nil,
		}
	}
}
