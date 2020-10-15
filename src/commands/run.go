package commands

import (
	"path/filepath"
	"strings"

	"github.com/dhenkel92/pod-exec/src/kube"
	"github.com/dhenkel92/pod-exec/src/log"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/homedir"
)

type Result struct {
	ExecResult kube.ExecResult
	Pod        *v1.Pod
}

func printResult(result *Result) {
	log.Info.Println("----------------------------------------")
	log.Info.Println(Green(result.Pod.Name))
	log.Info.Println(Green("Successful"))
	log.Info.Printf("Result:\n%s\n\n", result.ExecResult.StdOut)
	log.Info.Println("----------------------------------------")
}

func Run(c *cli.Context) error {
	namespace := c.String("namespace")
	labels := c.StringSlice("labels")
	command := c.String("command")

	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	clientset, err := kube.NewClientset(kubeconfig)
	if err != nil {
		return err
	}

	podExec, err := kube.NewPodExecutor(kubeconfig)
	if err != nil {
		return err
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: strings.Join(labels, ",")})
	if err != nil {
		return err
	}
	log.Trace.Printf("%d pods in ns %s\n", len(pods.Items), namespace)

	var results []Result
	for _, pod := range pods.Items {
		c := make(chan kube.ExecResult)
		go podExec.Exec(c, command, &pod)
		res := <-c

		results = append(results, Result{ExecResult: res, Pod: &pod})
	}

	for _, res := range results {
		printResult(&res)
	}

	return nil
}
