package types

import (
	"github.com/dhenkel92/pod-helper/src/kube"
	"github.com/dhenkel92/pod-helper/src/log"
	"github.com/logrusorgru/aurora"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type Result struct {
	ExecResult kube.ExecResult
	Pod        v1.Pod
	Container  v1.Container
}

func (result *Result) Print() {
	var text string
	var resText string
	if result.ExecResult.Error == nil {
		text = aurora.Sprintf(aurora.Green("Success on %s; Container %s"), result.Pod.Name, result.Container.Name)
		resText = result.ExecResult.StdOut
	} else {
		text = aurora.Sprintf(aurora.Red("Failed on %s; Container %s"), result.Pod.Name, result.Container.Name)
		resText = result.ExecResult.StdOut
	}

	lineByLine := strings.Split(resText, "\n")

	log.Info.Println("----------------------------------------")
	log.Info.Println(text)
	log.Info.Printf("Result:\n\n\t%s", strings.Join(lineByLine, "\n\t"))
	log.Info.Println("----------------------------------------")
}
