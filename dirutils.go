package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
// fileext: ".txt"
func FindFilesInDir(dir, fileext string) []string {
	if len(dir) == 0 {
		return []string{}
	}
	log.Printf("[DBG]Find in %s...\n", dir)
	allres := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) != fileext {
			return nil
		}

		allres = append(allres, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return allres
}

// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
func FindAllFilesInDir(dir string) []string {
	allres := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		allres = append(allres, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return allres
}

func FindFileCountInDir(dir string) int {
	count := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if strings.Contains(err.Error(), "Access is denied") {
				//log.Printf("Cannot read %s: %v\n", path, err)
				return nil
			}
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		count++
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return count
}

// https://golangr.com/rename-file/
func renameFile(src, dst string) {
	// rename file
	os.Rename(src, dst)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func GetFileLength(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return info.Size()
}

func PureFileName(filename string) string {
	return filepath.Base(filename)
}
