package kube

import (
	"bytes"
	"github.com/dhenkel92/pod-helper/src/config"
	"io"
	v1 "k8s.io/api/core/v1"
)

func (podExec *PodExecutor) Logs(c chan ExecResult, conf *config.Config, pod *v1.Pod, container *v1.Container) {
	options := v1.PodLogOptions{
		Container: container.Name,
	}

	if conf.LogsConfig.Tail >= 0 {
		options.TailLines = &conf.LogsConfig.Tail
	}

	req := podExec.Clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &options)
	podLogs, err := req.Stream()
	if err != nil {
		c <- ExecResult{Error: err}
		return
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		c <- ExecResult{Error: err}
		return
	}

	c <- ExecResult{StdOut: buf.String()}
}
