package handler

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// This struct is suboptimal to it's true purpose, I should make the fi part of the struct itself, however I would need to find a solution(one option is context) to close the file at the end of the caller's scope end
type File struct {
	path              string
	fileBufferChannel chan []byte
}

func (file *File) ReadFile(bytesize int, c func([]byte)) error {
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
			c(bytesBuffer[:n])
		}

		// Returns an EOF as an error hence
		if err != nil {
			break
		}
	}
	return nil
}

func (file *File) AppendToFile(filePath string, data []byte) error {
	// Ensure the parent directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Open the file in append mode, create it if it doesn't exist
	fi, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer fi.Close()

	// Write the data to the file
	if _, err := fi.Write(data); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
