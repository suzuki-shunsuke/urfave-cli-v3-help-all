package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/suzuki-shunsuke/urfave-cli-v3-help-all/helpall"
	"github.com/urfave/cli/v3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return helpall.With(&cli.Command{
		Name: "hello",
		Commands: []*cli.Command{
			{
				Name:        "foo",
				Usage:       "foo command",
				Description: "This is a foo command",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(cmd.Name)
					return nil
				},
			},
			{
				Name:        "bar",
				Usage:       "bar command",
				Description: "This is a bar command",
			},
		},
	}, nil).Run(context.Background(), os.Args)
}
