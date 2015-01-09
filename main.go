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
