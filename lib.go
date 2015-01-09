package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func replace(dir, oldS, newS string) error {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range dirFiles {
		name := f.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		fmt.Println(dir, name)
		p := path.Join(dir, name)
		if f.IsDir() {
			if err := replace(p, oldS, newS); err != nil {
				return err
			}
		}
		if strings.HasSuffix(name, ".go") {
			b, err := ioutil.ReadFile(p)
			if err != nil {
				return err
			}
			w := strings.Replace(string(b), oldS, newS, -1)
			if err = ioutil.WriteFile(p, []byte(w), f.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}
