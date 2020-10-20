package executor

import (
	"github.com/dhenkel92/pod-helper/src/config"
	"github.com/dhenkel92/pod-helper/src/kube"
	"github.com/dhenkel92/pod-helper/src/types"
	"github.com/dhenkel92/pod-helper/src/utils"
	v1 "k8s.io/api/core/v1"
)

func createResults(conf *config.Config, pods []v1.Pod) []types.Result {
	var results []types.Result
	for _, pod := range pods {
		containers, err := utils.FilterContainers(&pod.Spec.Containers, conf)
		if err != nil {
			result := types.Result{ExecResult: types.ExecResult{Error: err, StdOut: err.Error()}, Pod: pod}
			result.Print()
			continue
		}
		for _, container := range containers {
			results = append(results, types.Result{Pod: pod, Container: container})
		}
	}
	return results
}

func Execute(conf *config.Config, strategy CommandStrategy) error {
	clientset, err := kube.NewClientset(conf.Kubeconfig)
	if err != nil {
		return err
	}

	podExec, err := NewPodExecutor(conf.Kubeconfig)
	if err != nil {
		return err
	}

	pods, err := kube.ListPods(clientset, conf.Namespaces, conf.Labels)
	if err != nil {
		return err
	}

	results := createResults(conf, pods)
	divided := types.BatchResults(conf.BatchSize, results)

	for _, chunk := range divided {
		c := make(chan bool)
		for idx, _ := range chunk {
			go strategy(c, podExec, conf, &chunk[idx])
		}
		for _, _ = range chunk {
			<-c
		}
		for _, result := range chunk {
			result.Print()
		}
	}

	return nil
}
