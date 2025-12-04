package router

import (
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc(moviesRoute, getMovies)

	r.HandleFunc(movieRoute, getMovie)

	r.HandleFunc(movieRoute, createMovie)

	r.HandleFunc(movieRoute, deleteMovie)

	r.HandleFunc(movieRoute, updateMovie)
}
