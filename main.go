package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path"
	"runtime"
)

var (
	GoPath = os.Getenv("GOPATH")
	GoSrc  = path.Join(GoPath, "src")
)

func main() {
	app := cli.NewApp()
	app.Name = "got"
	app.Usage = ""
	app.Version = "0.1.0"
	app.Author = "Ethan Buchman"
	app.Email = "ethan@erisindustries.com"

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		replaceCmd,
		pullCmd,
		checkoutCmd,
		branchCmd,
		depCmd,
	}

	run(app)
}

// so we can catch panics
func run(app *cli.App) {
	defer func() {
		if r := recover(); r != nil {
			trace := make([]byte, 2048)
			count := runtime.Stack(trace, true)
			fmt.Printf("Panic: ", r)
			fmt.Printf("Stack of %d bytes: %s", count, trace)
		}
	}()

	app.Run(os.Args)
}

var (
	replaceCmd = cli.Command{
		Name:   "replace",
		Usage:  "String replace on all files in the directory tree",
		Action: cliReplace,
		Flags: []cli.Flag{
			pathFlag,
			depthFlag,
		},
	}

	pullCmd = cli.Command{
		Name:   "pull",
		Usage:  "Swap paths, pull changes, swap back",
		Action: cliPull,
		Flags:  []cli.Flag{},
	}

	branchCmd = cli.Command{
		Name:   "branch",
		Usage:  "List the branch every directory is on",
		Action: cliBranch,
		Flags:  []cli.Flag{},
	}

	checkoutCmd = cli.Command{
		Name:   "checkout",
		Usage:  "Checkout a git branch across all repos in the current dir. Add arguments like <repo>:<branch> to specify excpetions and <repo> to specify which repos to run checkout in, if not all.",
		Action: cliCheckout,
		Flags:  []cli.Flag{},
	}

	depCmd = cli.Command{
		Name:   "dep",
		Usage:  "Toggle the import statement for a godep managed dependency",
		Action: cliGodep,
		Flags: []cli.Flag{
			localPathFlag,
			vendorPathFlag,
			depthFlag,
			pathFlag,
		},
	}

	pathFlag = cli.StringFlag{
		Name:  "path, p",
		Usage: "specify the path to act upon",
		Value: ".",
	}

	depthFlag = cli.IntFlag{
		Name:  "depth, d",
		Usage: "specify the recursion depth",
		Value: -1,
	}

	localPathFlag = cli.BoolFlag{
		Name:  "local",
		Usage: "set the import path to the proper $GOPATH location",
	}

	vendorPathFlag = cli.BoolFlag{
		Name:  "vendor",
		Usage: "set the import path to the vendored location (a mirror of $GOPATH within the current repo)",
	}
)
