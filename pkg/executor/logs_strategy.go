package executor

import (
	"bytes"
	"io"

	"github.com/dhenkel92/pod-helper/pkg/config"
	"github.com/dhenkel92/pod-helper/pkg/types"
	v1 "k8s.io/api/core/v1"
)

func LogsStrategy(c chan bool, podExec *PodExecutor, conf *config.Config, result *types.Result) {
	options := v1.PodLogOptions{
		Container: result.Container.Name,
	}

	if conf.LogsConfig.Tail >= 0 {
		options.TailLines = &conf.LogsConfig.Tail
	}

	req := podExec.Clientset.CoreV1().Pods(result.Pod.Namespace).GetLogs(result.Pod.Name, &options)
	podLogs, err := req.Stream()
	if err != nil {
		result.ExecResult = types.ExecResult{Error: err}
		c <- true
		return
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		result.ExecResult = types.ExecResult{Error: err}
		c <- true
		return
	}

	result.ExecResult = types.ExecResult{StdOut: buf.String()}
	c <- true
}
