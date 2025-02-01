package handler

import (
	"bufio"
	"os"
)

type File struct {
	path string
}

func (file *File) ReadFile(r bufio.Reader) {
	fi, err := os.Open(file.path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
}
