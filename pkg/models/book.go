package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id          int     `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}
