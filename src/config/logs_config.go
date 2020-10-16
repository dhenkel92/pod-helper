package config

import "github.com/urfave/cli/v2"

func newLogsConfig(c *cli.Context) LogsConfig {
	containerIndex := c.Int64("container-index")

	return LogsConfig{ContainerIndex: containerIndex}
}
