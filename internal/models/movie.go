package models

type Movie struct {
	MovieID     string  `json:"movieId"`
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}
