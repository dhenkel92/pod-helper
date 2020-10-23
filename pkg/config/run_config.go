package config

import (
	"strings"

	"github.com/urfave/cli/v2"
)

func newRunConfig(c *cli.Context) RunConfig {
	entrypoint := c.String("entrypoint")
	command := c.String("command")

	return RunConfig{
		Entrypoint: strings.Split(entrypoint, " "),
		Command:    command,
	}
}
