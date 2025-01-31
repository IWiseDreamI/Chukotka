package models

import "gorm.io/gorm"

type Term struct {
	gorm.Model
	Title       	string `gorm:"size:255;not null;unique" json:"title"`
	Content     	string `gorm:"type:text;not null" json:"content"`
}
