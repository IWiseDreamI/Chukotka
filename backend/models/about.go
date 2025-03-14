package models

import "gorm.io/gorm"

type AboutPage struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
}