package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

// plop the config or genesis defaults into current dir
func cliReplace(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		log.Fatal("Please enter a string to replace and its replacement")
	}
	oldS := args[0]
	newS := args[1]
	dir := c.String("path")
	exit(replace(dir, oldS, newS))
}

func cliPull(c *cli.Context) {
}

func cliCheckout(c *cli.Context) {

}

func ifExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exit(err error) {
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
