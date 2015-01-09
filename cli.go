package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func checkArgs(c *cli.Context, n int) cli.Args {
	args := c.Args()
	if len(args) < n {
		exit(fmt.Errorf("Not enough arguments: require %d", n))
	}
	return args
}

// plop the config or genesis defaults into current dir
func cliReplace(c *cli.Context) {
	args := checkArgs(c, 2)
	oldS := args[0]
	newS := args[1]
	dir := c.String("path")
	depth := c.Int("depth")
	exit(replace(dir, oldS, newS, depth))
}

func cliPull(c *cli.Context) {
	args := checkArgs(c, 2)
	remote := args[0]
	branch := args[1]
	remotePath, err := resolveRemoteRepo(remote)
	ifExit(err)
	wd, err := os.Getwd()
	ifExit(err)
	localPath, err := resolveLocalRepo(wd)
	ifExit(err)
	localFullPath := path.Join(GoSrc, localPath)

	ifExit(replace(localFullPath, localPath, remotePath, -1))
	addCommit("change to upstream paths")
	gitPull(remote, branch)
	ifExit(replace(localFullPath, remotePath, localPath, -1))
}

func cliCheckout(c *cli.Context) {
	args := checkArgs(c, 1)
	branch := args[0]
	var repos []string
	if len(args) > 1 {
		repos = args[1:]
	}

	dir, _ := os.Getwd()

	if len(repos) == 0 {
		dirFiles, err := ioutil.ReadDir(dir)
		ifExit(err)
		for _, f := range dirFiles {
			name := f.Name()
			if strings.HasPrefix(name, ".") {
				continue
			}
			p := path.Join(dir, name)
			if f.IsDir() {
				gitCheckout(p, branch)
			}
		}
	}
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
