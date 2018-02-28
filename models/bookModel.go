package models

type Book struct {
	Id     int    `json: "id"`
	Title  string `json: "title"`
	Price  int64  `json: "price"`
	Status string `json: "status"`
}
