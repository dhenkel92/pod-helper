package kube

import (
	"bytes"
	"github.com/dhenkel92/pod-exec/src/config"
	"io"
	v1 "k8s.io/api/core/v1"
)

func (podExec *PodExecutor) Logs(c chan ExecResult, conf *config.Config, pod *v1.Pod) {
	options := v1.PodLogOptions{}
	if conf.LogsConfig.ContainerIndex >= 0 {
		options.Container = pod.Spec.Containers[conf.LogsConfig.ContainerIndex].Name
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

	str := buf.String()
	c <- ExecResult{StdOut: str}
}
