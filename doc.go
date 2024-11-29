package main

import (
	"fmt"
	"io/ioutil"
)

func ListSubDirs(dirname string) []string {
	fi, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Printf("WARN: Cannot read dir %s: %v!\n", dirname, err)
		return []string{}
	}

	alldir := []string{}

	for _, f := range fi {
		if f.IsDir() {
			alldir = append(alldir, f.Name())
			//fmt.Println(f.Name())
		}
	}

	return alldir
}

type FolderTree struct {
	dir2count map[string]int
}
