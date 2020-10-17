package config

import (
	"github.com/urfave/cli/v2"
	"strings"
)

func newRunConfig(c *cli.Context) RunConfig {
	command := c.String("command")
	commands := strings.Split(command, " ")

	return RunConfig{Command: commands}
}
