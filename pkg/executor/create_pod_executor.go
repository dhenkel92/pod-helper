package executor

import (
	"github.com/dhenkel92/pod-helper/pkg/kube"
)

func NewPodExecutor(kubeconfigPath string) (*PodExecutor, error) {
	clientset, err := kube.NewClientset(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	config, err := kube.NewRestclient(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return &PodExecutor{Clientset: clientset, Config: config}, nil
}
