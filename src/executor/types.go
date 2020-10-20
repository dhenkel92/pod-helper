package executor

import (
	"github.com/dhenkel92/pod-helper/src/config"
	"github.com/dhenkel92/pod-helper/src/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type CommandStrategy func(chan bool, *PodExecutor, *config.Config, *types.Result)

type PodExecutor struct {
	Clientset *kubernetes.Clientset
	Config    *rest.Config
}
