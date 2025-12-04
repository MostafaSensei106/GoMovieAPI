package models

// Director
type Director struct {
	ID       string `json:"id"`
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
}

// Movie
type Movie struct {
	ID               string    `json:"id"`
	Year             int       `json:"year"`
	Isdn             string    `json:"isdn"`
	Title            string    `json:"title"`
	Overview         string    `json:"overview"`
	OriginalLanguage string    `json:"original_language"`
	ReleaseDate      string    `json:"release_date"`
	Director         *Director `json:"director"`
}
