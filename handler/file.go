package handler

import "os"

type File struct {
	path string
}

func (file *File) ReadFile() {
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
