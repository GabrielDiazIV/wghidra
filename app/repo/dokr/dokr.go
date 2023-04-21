package dokr

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gabrieldiaziv/wghidra/app/bo"
)

func (r *runner) Run(ctx context.Context, doneCh chan<- bool) {
	taskDoneCh := make(chan bool)
	for _, task := range r.def.Tasks {
		go r.runTask(ctx, task, taskDoneCh)
	}

	tasksCompleted := 0
	for {

		if <-taskDoneCh {
			tasksCompleted++
		}

		if tasksCompleted == len(r.def.Tasks) {
			doneCh <- true
			return
		}
	}
}

func (r *runner) runTask(ctx context.Context, task bo.UnitTask, taskDoneCh chan<- bool) {
	defer func() {
		taskDoneCh <- true
	}()

	fmt.Println("preparing tasks - ", task.Name)
	if err := r.containerManager.PullImage(ctx, task.Runner); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("creating task - ", task.Name)
	if task.Exe == nil {
		panic("reader is emptyp in dokr")
	}

	id, err := r.containerManager.CreateContainer(ctx, task)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("starting task - ", task.Name)
	if err = r.containerManager.StartContainer(ctx, id); err != nil {
		fmt.Println(err)
		return
	}

	defer r.containerManager.RemoveContainer(ctx, id)

	fmt.Println("waiting task - ", task.Name)
	statusSucess, err := r.containerManager.WaitForContainer(ctx, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if statusSucess {
		fmt.Println("sucessed task - ", task.Name)
		stream, err := r.containerManager.CopyTarOutput(ctx, id)

		if err != nil {
			fmt.Println("could not parse tar", err)
			return
		}

		defer stream.Close()
		tr := tar.NewReader(stream)

		if _, err := tr.Next(); err != nil {
			panic("read err")
		}

		io.Copy(os.Stdout, stream)
	}
}
