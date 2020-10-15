package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type PodExecutor struct {
	Clientset *kubernetes.Clientset
	Config    *rest.Config
}

type ExecResult struct {
	StdOut string
	StdErr string
	Error  error
}
