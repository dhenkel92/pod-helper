package config

import "github.com/urfave/cli/v2"

func NewConfigFromCliContext(c *cli.Context) Config {
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
		RunConfig:  runConfig,
		LogsConfig: logsConfig,
	}
}
