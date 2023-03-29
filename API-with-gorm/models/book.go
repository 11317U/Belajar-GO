package models

import (
	"time"
)

type Books struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null;type:varchar(225)" json:"title"`
	Author    string    `gorm:"not null;type:varchar(225)" json:"author"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

// func MigrateBooks(db*gorm.DB).error  {
// 	err := db.AutoMigrate(&Books{})
// 	return err
// }
