package config

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func newRunConfig(c *cli.Context) RunConfig {
	command := c.String("command")

	commands := strings.Split(command, " ")
	fmt.Println(commands)

	return RunConfig{Command: commands}
}
