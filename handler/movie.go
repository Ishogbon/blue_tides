package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type MovieMeta struct {
	Width  uint32 `"json:width"`
	Height uint32 `"json:height"`
	Scale  uint32 `"json:scale"`
}

// There is always a minor performance letup using interfaces, use with sense.
type MovieBuffer interface {
	ReadFile(bytesize int) error
}

type Movie struct {
	File
}

func (m *Movie) ReadMovieMeta(movieMetaPath string) MovieMeta {

	movieMetaFile, err := os.ReadFile(movieMetaPath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read meta for Movie as Movie at %s", movieMetaFile))
	}
	var meta MovieMeta
	err = json.Unmarshal(movieMetaFile, &meta)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse JSON: %v", err))
	}
	return meta
}

func (m *Movie) PlayMovie(movieDir, movieName string) {
	if m.fileBufferChannel != nil {
		close(m.fileBufferChannel)
	}

	movieLoc := filepath.Join(movieDir, movieName)
	// asssume meta file is .meta.json
	movieMeta := m.ReadMovieMeta(movieLoc + ".meta.json")

	m.ReadFile(int(movieMeta.Scale))
}
