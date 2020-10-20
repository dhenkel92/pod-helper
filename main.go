package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/dhenkel92/pod-helper/src/config"

	"github.com/dhenkel92/pod-helper/src/executor"

	"k8s.io/client-go/util/homedir"

	"github.com/dhenkel92/pod-helper/src/log"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	app := &cli.App{
		Name:    "pod-exec",
		Usage:   "A tool to easily operate on mutliple pods at the same time.",
		Version: fmt.Sprintf("%s, built on %s (%s)", version, date, commit),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "kubeconfig",
				Aliases: []string{"config"},
				Usage:   "path to the kubeconfig file that will be used to authenticate to your cluster.",
				Value:   filepath.Join(homedir.HomeDir(), ".kube", "config"),
			},
			&cli.StringSliceFlag{
				Name:     "namespace",
				Aliases:  []string{"n"},
				Usage:    "the namespaces that are used for discovering the pods. If none is set it will use all of them.",
				Required: false,
				Value:    cli.NewStringSlice(""),
			},
			&cli.StringSliceFlag{
				Name:     "labels",
				Usage:    "set of labels which are used to filter the pods.",
				Aliases:  []string{"l"},
				Value:    &cli.StringSlice{},
				Required: false,
			},
			&cli.Int64Flag{
				Name:    "container-index",
				Usage:   "many pods do have more than one container, but often you don't know the specific container name or you want to execute the command always on the first one. With this flag you can define the index (beginning at 0) which should be used to get the container.",
				Aliases: []string{"ci"},
				Value:   -1,
			},
			&cli.StringFlag{
				Name:    "container",
				Usage:   "define a container name which should be searched for within the pod. If the pod doesn't have a container with the given name, it will return an error.",
				Aliases: []string{"con"},
				Value:   "",
			},
			&cli.IntFlag{
				Name:    "batch-size",
				Usage:   "WIP",
				Aliases: []string{"batch", "b"},
				Value:   5,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Runs a command in all the containers of all found pods.",
				Action: func(c *cli.Context) error {
					conf := config.NewConfigFromCliContext(c)
					return executor.Execute(&conf, executor.RunStrategy)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "entrypoint",
						Aliases:  []string{"e", "entry"},
						Required: false,
						Value:    "/bin/sh -c",
						Usage:    "by default every command will be executed in the /bin/sh shell. If you want to use a different one (e.g. /bin/bash) you can set it here.",
					},
					&cli.StringFlag{
						Name:     "command",
						Aliases:  []string{"c"},
						Required: true,
						Usage:    "the command that should be executed. If it contains a whitespace, it should be quoted.",
					},
				},
			},
			{
				Name:  "logs",
				Usage: "Returns the logs of all containers of all the found pods.",
				Action: func(c *cli.Context) error {
					conf := config.NewConfigFromCliContext(c)
					return executor.Execute(&conf, executor.LogsStrategy)
				},
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:    "tail",
						Aliases: []string{"t"},
						Value:   -1,
						Usage:   "will return only the last x log entries of the containers. -1 means return all available logs.",
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
