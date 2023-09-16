package presentation

import (
	"log"
	"os"

	"github.com/HungTP-Play/kompo/application"
	"github.com/HungTP-Play/kompo/entity"
	"github.com/urfave/cli/v2"
)

type KompoCli struct{}

func (k *KompoCli) Start() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			{
				Name: "convert",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "mode",
						Aliases:     []string{"mod", "m"},
						DefaultText: "apply",
					},
					&cli.StringFlag{
						Name:        "input",
						Aliases:     []string{"in", "i"},
						DefaultText: "./docker-compose.yaml",
					},
					&cli.StringFlag{
						Name:        "output",
						Aliases:     []string{"out", "o"},
						DefaultText: "./kompo_output",
					},
					&cli.StringFlag{
						Name:        "version",
						Aliases:     []string{"ver", "v"},
						DefaultText: "latest",
					},
					&cli.BoolFlag{
						Name:    "reserve",
						Aliases: []string{"rev", "r"},
					},
				},
				Action: func(cCtx *cli.Context) error {
					mode := cCtx.String("mode")
					input := cCtx.String("input")
					output := cCtx.String("output")
					version := cCtx.String("version")
					reverse := cCtx.Bool("reserve")

					kompoOption := entity.KompoOption{
						In:         input,
						Out:        output,
						K8sVersion: version,
						Mode:       mode,
						Reserve:    reverse,
					}

					app := application.Converter{}
					err := app.Convert(&kompoOption)

					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
