package router

import (
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc(moviesRoute, getMovies).Methods(GET)
	r.HandleFunc(movieRoute, getMovie).Methods(GET)
	r.HandleFunc(movieRoute, createMovie).Methods(POST)
	r.HandleFunc(movieRoute, deleteMovie).Methods(DELETE)
	r.HandleFunc(movieRoute, updateMovie).Methods(PUT)

}
