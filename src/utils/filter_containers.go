package utils

import (
	"errors"
	"fmt"

	"github.com/dhenkel92/pod-helper/src/config"
	v1 "k8s.io/api/core/v1"
)

func FilterContainers(containers *[]v1.Container, conf *config.Config) ([]v1.Container, error) {
	resultContainers := *containers

	if conf.Container != "" {
		resultContainers = []v1.Container{}
		for _, container := range *containers {
			if container.Name == conf.Container {
				resultContainers = append(resultContainers, container)
			}
		}
		if len(resultContainers) == 0 {
			return nil, errors.New(fmt.Sprintf("Pod has no container with name %s", conf.Container))
		}
	}

	if conf.ContainerIndex >= 0 {
		if int(conf.ContainerIndex) >= len(*containers) {
			return nil, errors.New(fmt.Sprintf("Pod has only %d containers, so it cannot use index %d", len(*containers), conf.ContainerIndex))
		}
		resultContainers = []v1.Container{(*containers)[conf.ContainerIndex]}
	}

	return resultContainers, nil
}
