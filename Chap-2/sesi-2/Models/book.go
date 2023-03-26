package models

type Book struct {
	BookID int    `json:"bookId"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Bookdata = []Book{}
