package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <folder> <outhtml>\n", os.Args[0])
		return
	}

	workdir := os.Args[1]
	tofile := os.Args[2]

	if !checkFileExists(workdir) {
		fmt.Printf("WARN: %s is not exist!\n", workdir)
		return
	}

	ConvertChart(workdir, tofile)
}
