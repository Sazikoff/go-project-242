// Package main implements the hexlet-path-size CLI application.
package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},

			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories",
			},

			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				return fmt.Errorf("path is required")
			}

			path := cmd.Args().Get(0)
			size, _ := code.GetPathSize(path, cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))

			fmt.Printf("%s	%s\n", size, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
