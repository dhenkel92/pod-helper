package executor

import (
	"bytes"

	"github.com/dhenkel92/pod-helper/pkg/config"
	"github.com/dhenkel92/pod-helper/pkg/types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func RunStrategy(c chan bool, podExecutor *PodExecutor, conf *config.Config, result *types.Result) {
	req := podExecutor.Clientset.
		CoreV1().
		RESTClient().
		Post().
		Resource("pods").
		Name(result.Pod.Name).
		Namespace(result.Pod.Namespace).
		SubResource("exec")

	option := &v1.PodExecOptions{
		Container: result.Container.Name,
		Command:   append(conf.RunConfig.Entrypoint, conf.RunConfig.Command),
		Stdin:     false,
		Stdout:    true,
		Stderr:    true,
	}

	req.VersionedParams(option, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(podExecutor.Config, "POST", req.URL())
	if err != nil {
		result.ExecResult = types.ExecResult{Error: err}
		c <- true
		return
	}

	var stdout, stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		result.ExecResult = types.ExecResult{Error: err, StdErr: stderr.String(), StdOut: stdout.String()}
		c <- true
		return
	}

	result.ExecResult = types.ExecResult{StdOut: stdout.String(), StdErr: stderr.String()}
	c <- true
}
