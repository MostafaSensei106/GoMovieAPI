package movies_repo

import "github.com/MostafaSensei106/GoMovieAPI/internal/data/models"

// Global slice
var Movies []models.Movie

func GetAll() []models.Movie {
	return Movies
}

func Get(id string) *models.Movie {
	for _, movie := range Movies {
		if movie.ID == id {
			return &movie
		}
	}
	return nil
}

func Delete(id string) bool {
	for index, movie := range Movies {
		if movie.ID == id {
			Movies = append(Movies[:index], Movies[index+1:]...)
			return true
		}
	}
	return false
}

func Create(movie models.Movie) {
	Movies = append(Movies, movie)
}

func Update(id string, movie models.Movie) bool {
	for index, m := range Movies {
		if m.ID == id {
			Movies[index] = movie
			return true
		}
	}
	return false
}

func InitMovies() {
	Movies = append(Movies,
		models.Movie{
			ID:               "1",
			Year:             1994,
			Isdn:             "123456789",
			Title:            "The Shawshank Redemption",
			Overview:         "Two imprisoned men bond over a number of years...",
			OriginalLanguage: "en",
			ReleaseDate:      "1994-09-23",
			Director: &models.Director{
				ID:        "1",
				FirstName: "Frank",
				LastName:  "Darabont",
			},
		},
		models.Movie{
			ID:               "2",
			Year:             1999,
			Isdn:             "987654321",
			Title:            "Fight Club",
			Overview:         "An insomniac office worker and a devil-may-care soap maker...",
			OriginalLanguage: "en",
			ReleaseDate:      "1999-10-15",
			Director: &models.Director{
				ID:        "2",
				FirstName: "David",
				LastName:  "Fincher",
			},
		},
		models.Movie{
			ID:               "3",
			Year:             2006,
			Isdn:             "123456789",
			Title:            "The Prestige",
			Overview:         "A mystery set in the world of illusion...",
			OriginalLanguage: "en",
			ReleaseDate:      "2006-10-20",
			Director: &models.Director{
				ID:        "3",
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
		models.Movie{
			ID:               "4",
			Year:             2008,
			Isdn:             "987654321",
			Title:            "The Dark Knight",
			Overview:         "When the menace known as the Joker wreaks havoc...",
			OriginalLanguage: "en",
			ReleaseDate:      "2008-07-18",
			Director: &models.Director{
				ID:        "4",
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
		models.Movie{
			ID:               "5",
			Year:             2010,
			Isdn:             "123456789",
			Title:            "Inception",
			Overview:         "A thief who steals corporate secrets...",
			OriginalLanguage: "en",
			ReleaseDate:      "2010-07-16",
			Director: &models.Director{
				ID:        "5",
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
		models.Movie{
			ID:               "6",
			Year:             2014,
			Isdn:             "987654321",
			Title:            "Interstellar",
			Overview:         "A team of explorers travels through a wormhole...",
			OriginalLanguage: "en",
			ReleaseDate:      "2014-11-05",
			Director: &models.Director{
				ID:        "6",
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
		models.Movie{
			ID:               "7",
			Year:             2015,
			Isdn:             "123456789",
			Title:            "Mad Max: Fury Road",
			Overview:         "In a post-apocalyptic wasteland...",
			OriginalLanguage: "en",
			ReleaseDate:      "2015-05-15",
			Director: &models.Director{
				ID:        "7",
				FirstName: "George",
				LastName:  "Miller",
			},
		},
		models.Movie{
			ID:               "8",
			Year:             2018,
			Isdn:             "987654321",
			Title:            "Avengers: Infinity War",
			Overview:         "The Avengers and their allies must sacrifice all...",
			OriginalLanguage: "en",
			ReleaseDate:      "2018-04-27",
			Director: &models.Director{
				ID:        "8",
				FirstName: "Anthony",
				LastName:  "Russo",
			},
		},
		models.Movie{
			ID:               "9",
			Year:             2019,
			Isdn:             "123456789",
			Title:            "Parasite",
			Overview:         "A poor family, the Kims, con a rich family...",
			OriginalLanguage: "ko",
			ReleaseDate:      "2019-05-21",
			Director: &models.Director{
				ID:        "9",
				FirstName: "Bong",
				LastName:  "Joon-ho",
			},
		},
	)

}
