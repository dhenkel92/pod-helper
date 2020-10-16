package config

import "github.com/urfave/cli/v2"

func NewConfigFromCliContext(c *cli.Context) Config {
	kubeconfig := c.String("kubeconfig")
	namespace := c.String("namespace")
	labels := c.StringSlice("labels")
	command := c.String("command")
	if c.Bool("all-namespaces") {
		namespace = ""
	}

	return Config{
		Kubeconfig: kubeconfig,
		Namespace:  namespace,
		Labels:     labels,
		Command:    command,
	}
}
