package handler

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	path string
}

func (file *File) ReadFile(bytesSize int) error {
	fi, err := os.Open(file.path)
	if err != nil {
		return err
	}

	defer fi.Close()

	bytesBuffer := make([]byte, bytesSize)

	reader := bufio.NewReader(fi)

	for {
		n, err := reader.Read(bytesBuffer)

		if n > 0 {
			fmt.Print(string(bytesBuffer[:n])) // Print the read bytes
		}

		// Check for EOF (End of File)
		if err != nil {
			break
		}
	}
	return nil
}
