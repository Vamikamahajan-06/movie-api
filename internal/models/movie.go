package models

type Movie struct {
	ID          string  `json: "id"`
	Title       string  `json: "Title"`
	Genre       string  `json: "Genre"`
	Year        int     `json: "Year"`
	Rating      float64 `json: "Rating"`
	Description string  `json: "Description"`
}
