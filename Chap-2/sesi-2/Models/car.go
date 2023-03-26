package models

type Car struct {
	CarID int    `json:"carid"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var Cardata = []Car{}
