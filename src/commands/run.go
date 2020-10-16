package commands

import (
	"github.com/dhenkel92/pod-exec/src/config"
	"github.com/dhenkel92/pod-exec/src/kube"
	"github.com/dhenkel92/pod-exec/src/log"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	v1 "k8s.io/api/core/v1"
)

type Result struct {
	ExecResult kube.ExecResult
	Pod        *v1.Pod
}

func (result *Result) print() {
	log.Info.Println("----------------------------------------")
	log.Info.Println(Green(result.Pod.Name))
	log.Info.Println(Green("Successful"))
	log.Info.Printf("Result:\n%s", result.ExecResult.StdOut)
	log.Info.Println("----------------------------------------")
}

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

	var results []Result
	for _, pod := range pods.Items {
		c := make(chan kube.ExecResult)
		go podExec.Exec(c, cliConf.Command, &pod)
		res := <-c

		results = append(results, Result{ExecResult: res, Pod: &pod})
	}

	for _, res := range results {
		res.print()
	}

	return nil
}
