package cm

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/labstack/gommon/log"

	api "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/pkg/errors"
)

func (cm *containerManager) PullImage(ctx context.Context) error {
	out, err := cm.cli.ImagePull(ctx, image_name, api.ImagePullOptions{})
	if err != nil {
		return errors.Wrap(err, "DOCKER PULL")
	}

	defer func() {
		if err := out.Close(); err != nil {
			log.Error(err)
		}
	}()

	fd := json.NewDecoder(out)
	var status *bo.ImagePullStatus

	for {
		if err := fd.Decode(&status); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}

		if status.Error != "" {
			return errors.Wrap(errors.New(status.Error), "DOCKER PULL")
		}
	}
	return nil
}

func (cm *containerManager) CreateContainer(ctx context.Context, task bo.UnitTask, inputStream io.Reader) (string, error) {
	config := &container.Config{
		Image:        image_name,
		Cmd:          task.Task.Cmd(),
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
	}

	log.Infof("CREATING CMD: %v", task.Task.Cmd())

	res, err := cm.cli.ContainerCreate(ctx, config, &container.HostConfig{}, nil, nil, task.ID)
	if err != nil {
		return "", nil
	}

	if err := cm.cli.CopyToContainer(
		ctx, res.ID, "/container/input/", inputStream,
		api.CopyToContainerOptions{AllowOverwriteDirWithFile: true}); err != nil {
		log.Errorf("cannot copy file: ", err)
		return "", err
	}
	return res.ID, nil

}

func (cm *containerManager) StartContainer(ctx context.Context, id string) error {
	go func() {
		reader, err := cm.cli.ContainerLogs(context.Background(), id, api.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
			Timestamps: false,
		})
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	return cm.cli.ContainerStart(ctx, id, api.ContainerStartOptions{})
}

func (cm *containerManager) CopyTarOutput(ctx context.Context, id string) (io.ReadCloser, error) {
	tarStream, _, err := cm.cli.CopyFromContainer(ctx, id, "/container/output/output.json")
	if err != nil {
		return nil, err
	}

	return tarStream, nil
}

func (cm *containerManager) WaitForContainer(ctx context.Context, id string) (bool, error) {
	if _, err := cm.cli.ContainerInspect(ctx, id); err != nil {
		return true, nil
	}

	wait, errC := cm.cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)

	select {
	case status := <-wait:
		if status.StatusCode == 0 {
			return true, nil
		}
		return false, nil
	case err := <-errC:
		return false, errors.Wrap(err, "DOCKER_WAIT")
		// case <-ctx.Done():
		// 	return false, ctx.Err()
	}
}

func (cm *containerManager) RemoveContainer(ctx context.Context, id string) error {
	return cm.cli.ContainerRemove(ctx, id, api.ContainerRemoveOptions{})
}
