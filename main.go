package main

import (
	"github.com/MostafaSensei106/GoMovieAPI/internal/movies_repo"
	"github.com/MostafaSensei106/GoMovieAPI/internal/router"
)

func main() {
	movies_repo.InitMovies()
	router.ExtractRoutes()
}
