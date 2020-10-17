package commands

import (
	"github.com/dhenkel92/pod-helper/src/config"
	"github.com/dhenkel92/pod-helper/src/kube"
	"github.com/dhenkel92/pod-helper/src/types"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
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
		c := make(chan kube.ExecResult)
		go podExec.Run(c, &cliConf, &pod)
		res := <-c

		result := types.Result{ExecResult: res, Pod: pod}
		result.Print()
	}

	return nil
}
