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

// replace a line of text in every file with another
func cliReplace(c *cli.Context) {
	args := checkArgs(c, 2)
	oldS := args[0]
	newS := args[1]
	dir := c.String("path")
	depth := c.Int("depth")
	exit(replace(dir, oldS, newS, depth))
}

// replace import paths with host, pull, replace back
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

// checkout a branch across every repository in a directory
func cliCheckout(c *cli.Context) {
	args := checkArgs(c, 1)
	branch := args[0]
	var repos []string
	if len(args) > 1 {
		repos = args[1:]
	}

	var nonColon bool
	repoMap := make(map[string]string)
	for _, r := range repos {
		sp := strings.Split(r, ":")
		repo := sp[0]
		var b string
		if len(sp) != 2 {
			nonColon = true
			b = branch
			//ifExit(fmt.Errorf("Additional arguments must be of the form <repo>:<branch>"))
		} else {
			b = sp[1]
		}
		repoMap[repo] = b
	}

	dir, _ := os.Getwd()

	// if nonColon, we only loop through dirs in the repoMap
	if nonColon {
		for r, b := range repoMap {
			p := path.Join(dir, r)
			f, err := os.Stat(p)
			if err != nil {
				log.Println("Unknown repo:", r)
				continue
			}
			if !f.IsDir() {
				log.Println(r, " is not a directory")
			}
			gitCheckout(p, b)
		}
		exit(nil)
	}

	// otherwise, we loop through all dirs in the current one
	dirFiles, err := ioutil.ReadDir(dir)
	ifExit(err)
	for _, f := range dirFiles {
		name := f.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		p := path.Join(dir, name)
		if f.IsDir() {
			if b, ok := repoMap[name]; ok {
				gitCheckout(p, b)
			} else {
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
