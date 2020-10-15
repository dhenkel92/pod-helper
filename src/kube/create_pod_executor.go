package kube

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewRestclient(kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

func NewPodExecutor(kubeconfigPath string) (*PodExecutor, error) {
	clientset, err := NewClientset(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	config, err := NewRestclient(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return &PodExecutor{Clientset: clientset, Config: config}, nil
}
