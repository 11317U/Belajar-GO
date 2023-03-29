package controller

import (
	"gorm.io/gorm"
)

type Controllers struct {
	masterDB *gorm.DB
}

func New(db *gorm.DB) *Controllers {
	return &Controllers{
		masterDB: db,
	}
}
