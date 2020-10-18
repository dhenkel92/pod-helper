package config

import "github.com/urfave/cli/v2"

func NewConfigFromCliContext(c *cli.Context) Config {
	containerIndex := c.Int64("container-index")
	container := c.String("container")

	kubeconfig := c.String("kubeconfig")
	namespace := c.String("namespace")
	labels := c.StringSlice("labels")
	if c.Bool("all-namespaces") {
		namespace = ""
	}

	runConfig := newRunConfig(c)
	logsConfig := newLogsConfig(c)

	return Config{
		Kubeconfig: kubeconfig,
		Namespace:  namespace,
		Labels:     labels,

		ContainerIndex: containerIndex,
		Container:      container,

		RunConfig:  runConfig,
		LogsConfig: logsConfig,
	}
}
