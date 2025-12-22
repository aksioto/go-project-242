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
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Value:       false,
				DefaultText: "false",
				Usage:       "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Value:       false,
				DefaultText: "false",
				Usage:       "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Value:       false,
				DefaultText: "false",
				Usage:       "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("path is required")
			}

			path := c.Args().Get(0)
			recursive := c.Bool("recursive")
			human := c.Bool("human")
			all := c.Bool("all")

			result, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", result, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
