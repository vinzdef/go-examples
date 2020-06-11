package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

type ListingError struct {
	Path string
}

func (le ListingError) Error() string {
	return fmt.Sprint("Couldn't list", le.Path)
}

func ls(path string) (files []os.FileInfo, err error) {	
	if exists := pathExists(path); !exists {
		err = ListingError{path}
		return
	}

	return ioutil.ReadDir(path)
}

func main() {
	files, err := ls("/tmp/foo/bar")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	if err != nil {
		panic(err)
	}

	for _, file := range(files) {
		fmt.Println(file)	
	}
}

