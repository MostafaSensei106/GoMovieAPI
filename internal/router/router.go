package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ExtractRoutes() {
	r := mux.NewRouter()

	r.HandleFunc(moviesRoute, getMovies).Methods(GET)
	r.HandleFunc(movieRoute, getMovie).Methods(GET)
	r.HandleFunc(movieRoute, createMovie).Methods(POST)
	r.HandleFunc(movieRoute, deleteMovie).Methods(DELETE)
	r.HandleFunc(movieRoute, updateMovie).Methods(PUT)

	fmt.Printf("Starting http server at port 8080\n")

	http.ListenAndServe(":8080", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
