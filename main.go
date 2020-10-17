package main

import (
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"sort"

	"github.com/dhenkel92/pod-helper/src/commands"
	"github.com/dhenkel92/pod-helper/src/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pod-exec",
		Usage: "do smth",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "kubeconfig",
				Aliases: []string{"config"},
				Value:   filepath.Join(homedir.HomeDir(), ".kube", "config"),
			},
			&cli.StringFlag{
				Name:     "namespace",
				Aliases:  []string{"n"},
				Usage:    "select namespace",
				Value:    "default",
				Required: false,
			},
			&cli.StringSliceFlag{
				Name:     "labels",
				Aliases:  []string{"l"},
				Value:    &cli.StringSlice{},
				Required: false,
			},
			&cli.BoolFlag{
				Name:    "all-namespaces",
				Aliases: []string{"all", "a"},
				Value:   false,
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
			{
				Name:   "logs",
				Usage:  "logs",
				Action: commands.Logs,
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:    "container-index",
						Aliases: []string{"ci"},
						Value:   -1,
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
