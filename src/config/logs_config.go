package config

import "github.com/urfave/cli/v2"

func newLogsConfig(c *cli.Context) LogsConfig {
	tail := c.Int64("tail")

	return LogsConfig{
		Tail: tail,
	}
}
