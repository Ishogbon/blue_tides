package handler

import (
	"os"
	"path/filepath"
)

// There is always a minor performance letup using interfaces, use with sense.
type MovieBuffer interface {
	ReadFile(bytesize int) error
}

type MoviesReader struct {
	moviesDir    string
	moviesBuffer map[string]MovieBuffer
}

func (m *MoviesReader) StreamMovies() {
	dir, err := os.ReadDir(m.moviesDir)
	if err != nil {
		panic("Unable to read dir for movies")
	}
	for _, file := range dir {
		if !file.IsDir() && filepath.Ext(file.Name()) != ".mp4" {
			m.moviesBuffer[file.Name()] = &File{path: filepath.Join(m.moviesDir, file.Name())}
		}
	}

}
