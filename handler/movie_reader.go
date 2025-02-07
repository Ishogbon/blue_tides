package handler

type MovieBuffer interface {
	ReadFile()
}

type MoviesReader struct {
	moviesDir    string
	moviesBuffer map[string][]*MovieBuffer
}

func (m *MoviesReader) StreamMovies() {

}
