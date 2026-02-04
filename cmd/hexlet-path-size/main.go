package main

import (
	"code/internal/pathsize"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				return fmt.Errorf("path is required")
			}

			path := cmd.Args().Get(0)
			size,_ := pathsize.GetSize(path)

			fmt.Printf("%dB	%s\n", size, path)
			// cwd,_ := os.Getwd()
			// fmt.Println(cwd)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}