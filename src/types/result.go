package types

import (
	"strings"

	"github.com/dhenkel92/pod-helper/src/log"
	"github.com/logrusorgru/aurora"
	v1 "k8s.io/api/core/v1"
)

type Result struct {
	ExecResult ExecResult
	Pod        v1.Pod
	Container  v1.Container
}

func (result *Result) printString(str string, isError bool) {
	podName := aurora.Green(result.Pod.Name)
	if isError {
		podName = aurora.Red(result.Pod.Name)
	}

	lineByLine := strings.Split(str, "\n")
	for _, line := range lineByLine {
		if line == "" {
			continue
		}
		log.Raw.Printf("%s/%s: %s", podName, aurora.Gray(14, result.Container.Name), line)
	}
}

func (result *Result) Print() {
	if result.ExecResult.Error == nil {
		result.printString(result.ExecResult.StdOut, false)
	} else {
		result.printString("Failed:\n", true)
		result.printString(result.ExecResult.Error.Error(), true)
		if result.ExecResult.StdOut != "" {
			result.printString(result.ExecResult.StdOut, true)
		}
		if result.ExecResult.StdErr != "" {
			result.printString(result.ExecResult.StdErr, true)
		}
	}
	// Add a line break after all the lines were printed to the screen
	// so that the results / container are easier to differntiate.
	log.Raw.Println("")
}
