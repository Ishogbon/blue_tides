package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
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
	FrameBenchmarksChannel chan time.Duration
}

func (m *Movie) ReadMovieMeta(movieMetaPath string) MovieMeta {

	movieMetaFile, err := os.ReadFile(movieMetaPath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read meta for Movie at %s", movieMetaPath))
	}
	var meta MovieMeta
	err = json.Unmarshal(movieMetaFile, &meta)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse JSON: %v", err))
	}
	return meta
}

func (m *Movie) PlayMovie(movieDir string, movieName string, streamMovie func([]byte, time.Duration)) {
	movieLoc := filepath.Join(movieDir, movieName)
	frameLoadStartTime := time.Now()
	// asssume meta file is .meta.json
	movieMeta := m.ReadMovieMeta(movieLoc + ".meta.json")
	m.path = movieLoc
	m.ReadFile(int(movieMeta.Scale), func(b []byte) {
		defer func() {
			frameLoadStartTime = time.Now()
		}()
		frameLoadTime := time.Since(frameLoadStartTime)
		streamMovie(b, frameLoadTime)
	})
}
