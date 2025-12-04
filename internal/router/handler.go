package router

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/MostafaSensei106/GoMovieAPI/internal/constants"
	"github.com/MostafaSensei106/GoMovieAPI/internal/data/models"
	"github.com/MostafaSensei106/GoMovieAPI/internal/movies_repo"
)

// --------------------------------
// Error Struct
// --------------------------------
type NetworkError struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

// --------------------------------
// Helpers
// --------------------------------
func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func jsonError(w http.ResponseWriter, err NetworkError) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}

// --------------------------------
// Handlers
// --------------------------------
func getMovies(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, movies_repo.GetAll())
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)
	movie := movies_repo.Get(parms["id"])

	if movie == nil {
		jsonError(w, NetworkError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "MOVIE_NOT_FOUND",
			Message:    "Movie not found",
		})
		return
	}

	jsonResponse(w, http.StatusOK, movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		jsonError(w, NetworkError{
			StatusCode: http.StatusBadRequest,
			ErrorCode:  "INVALID_JSON",
			Message:    "Invalid JSON body",
		})
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies_repo.Create(movie)
	jsonResponse(w, http.StatusCreated, movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)
	deleted := movies_repo.Delete(parms["id"])

	if !deleted {
		jsonError(w, NetworkError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "MOVIE_NOT_FOUND",
			Message:    "Movie not found",
		})
		return
	}

	jsonResponse(w, http.StatusOK, movies_repo.GetAll())
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		jsonError(w, NetworkError{
			StatusCode: http.StatusBadRequest,
			ErrorCode:  "INVALID_JSON",
			Message:    "Invalid JSON body",
		})
		return
	}

	updated := movies_repo.Update(parms["id"], movie)
	if !updated {
		jsonError(w, NetworkError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "MOVIE_NOT_FOUND",
			Message:    "Movie not found",
		})
		return
	}

	jsonResponse(w, http.StatusOK, movie)
}
