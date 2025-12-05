package cmd

import (
	"github.com/MostafaSensei106/GoMovieAPI/internal/movies_repo"
	"github.com/MostafaSensei106/GoMovieAPI/internal/router"
)

func Execute() {
	movies_repo.InitMovies()
	router.ExtractRoutes()
}
