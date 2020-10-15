package main

import (
	"os"
	"sort"

	"github.com/dhenkel92/pod-exec/src/commands"
	"github.com/dhenkel92/pod-exec/src/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pod-exec",
		Usage: "do smth",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "namespace",
				Aliases:  []string{"n"},
				Usage:    "select namespace",
				Value:    "",
				Required: false,
			},
			&cli.StringSliceFlag{
				Name:     "labels",
				Aliases:  []string{"l"},
				Value:    &cli.StringSlice{},
				Required: false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "runs smth",
				Action: commands.Run,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "command",
						Aliases:  []string{"c"},
						Required: true,
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Error.Fatal(err)
	}
}
