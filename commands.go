package main

import (
	"github.com/codegangsta/cli"
)

var (
	replaceCmd = cli.Command{
		Name:   "replace",
		Usage:  "String replace on all files in the directory tree",
		Action: cliReplace,
		Flags: []cli.Flag{
			pathFlag,
		},
	}

	pullCmd = cli.Command{
		Name:   "pull",
		Usage:  "Swap paths, pull changes, swap back",
		Action: cliPull,
		Flags:  []cli.Flag{},
	}

	checkoutCmd = cli.Command{
		Name:   "checkout",
		Usage:  "Checkout a git branch across many dirs",
		Action: cliCheckout,
		Flags:  []cli.Flag{},
	}

	pathFlag = cli.StringFlag{
		Name:  "path, p",
		Usage: "specify the path to act upon",
		Value: ".",
	}

	depthFlag = cli.IntFlag{
		Name:  "depth, d",
		Usage: "specify the recursion depth",
	}
)
