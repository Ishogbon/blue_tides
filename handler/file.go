package handler

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	path              string
	fileBufferChannel chan []byte
}

func (file *File) ReadFile(bytesize int) error {
	// file.fileBufferChannel = make(chan []byte)

	fi, err := os.Open(file.path)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file at %s", file.path))
	}

	defer fi.Close()

	bytesBuffer := make([]byte, bytesize)

	reader := bufio.NewReader(fi)
	for {
		n, err := reader.Read(bytesBuffer)

		if n > 0 {
			fmt.Print(bytesBuffer[:n])
		}

		if err != nil {
			break
		}
	}
	return nil
}
