package main

import (
	"fmt"
	"os"

	"github.com/brad-jones/hashdir"
	"github.com/urfave/cli/v2"
)

// Injected by ldflags
// see: https://stackoverflow.com/questions/11354518
var (
	version = "0.0.0"
)

func main() {
	if err := (&cli.App{
		Name:      "hashdir",
		Usage:     "get the hash of a directory",
		UsageText: "hashdir [global options] <dir>",
		Description: "A cli tool that will recursively calculate an idempotent\n" +
			"hash of a given directory, ignoring all metadata such as\n" +
			"permission bits, timestamps, etc...",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "Brad Jones",
				Email: "brad@bjc.id.au",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "alg",
				Usage: "the hash algorithm to use [values: md5, sha1, sha256, sha512]",
				Value: "sha256",
			},
		},
		Action: func(c *cli.Context) error {
			dir := c.Args().First()
			if dir == "" {
				fmt.Fprint(os.Stderr, "Error: missing <dir> argument\n\n")
				cli.ShowAppHelpAndExit(c, 1)
			}

			dirHash, err := hashdir.Make(dir, c.String("alg"))
			if err != nil {
				return fmt.Errorf("failed to calc hash: %w", err)
			}

			fmt.Fprint(os.Stdout, dirHash)
			return nil
		},
	}).Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
