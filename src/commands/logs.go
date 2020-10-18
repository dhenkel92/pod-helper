package commands

import (
	"github.com/dhenkel92/pod-helper/src/config"
	"github.com/dhenkel92/pod-helper/src/kube"
	"github.com/dhenkel92/pod-helper/src/types"
	"github.com/dhenkel92/pod-helper/src/utils"
	"github.com/urfave/cli/v2"
)

func Logs(c *cli.Context) error {
	cliConf := config.NewConfigFromCliContext(c)

	clientset, err := kube.NewClientset(cliConf.Kubeconfig)
	if err != nil {
		return err
	}

	podExec, err := kube.NewPodExecutor(cliConf.Kubeconfig)
	if err != nil {
		return err
	}

	pods, err := kube.ListPods(clientset, cliConf.Namespace, cliConf.Labels)
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		containers, err := utils.FilterContainers(&pod.Spec.Containers, &cliConf)
		if err != nil {
			result := types.Result{ExecResult: kube.ExecResult{Error: err, StdOut: err.Error()}, Pod: pod}
			result.Print()
			continue
		}
		for _, container := range containers {
			c := make(chan kube.ExecResult)
			go podExec.Logs(c, &cliConf, &pod, &container)
			res := <-c

			result := types.Result{ExecResult: res, Pod: pod, Container: container}
			result.Print()
		}
	}

	return nil
}
