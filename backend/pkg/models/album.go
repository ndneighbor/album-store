package models

type Album struct {
	ID     string  `json:"id"`
	Image  string  `json:"image"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
