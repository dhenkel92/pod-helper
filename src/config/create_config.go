package config

import "github.com/urfave/cli/v2"

func NewConfigFromCliContext(c *cli.Context) Config {
	containerIndex := c.Int64("container-index")
	container := c.String("container")

	kubeconfig := c.String("kubeconfig")
	namespaces := c.StringSlice("namespace")
	labels := c.StringSlice("labels")

	runConfig := newRunConfig(c)
	logsConfig := newLogsConfig(c)

	return Config{
		Kubeconfig: kubeconfig,
		Namespaces: namespaces,
		Labels:     labels,

		ContainerIndex: containerIndex,
		Container:      container,

		RunConfig:  runConfig,
		LogsConfig: logsConfig,
	}
}
