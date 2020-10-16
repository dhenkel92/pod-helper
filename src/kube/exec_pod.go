package kube

import (
	"bytes"
	"github.com/dhenkel92/pod-exec/src/config"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (podExec *PodExecutor) Exec(c chan ExecResult, conf *config.Config, pod *v1.Pod) {
	req := podExec.Clientset.
		CoreV1().
		RESTClient().
		Post().
		Resource("pods").
		Name(pod.Name).
		Namespace(pod.Namespace).
		SubResource("exec")

	option := &v1.PodExecOptions{
		Command: conf.RunConfig.Command,
		Stdin:   false,
		Stdout:  true,
		Stderr:  true,
	}

	req.VersionedParams(option, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(podExec.Config, "POST", req.URL())
	if err != nil {
		c <- ExecResult{Error: err}
		return
	}

	var stdout, stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		c <- ExecResult{Error: err}
		return
	}

	c <- ExecResult{StdOut: stdout.String(), StdErr: stderr.String()}
}
