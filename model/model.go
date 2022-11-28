package model

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Genre    string `json:"genre"`
	AuthorID string `json:"-"`
}
