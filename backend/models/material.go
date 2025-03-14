package models

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Title          string `gorm:"size:255;not null" json:"title"`
	Content    string `gorm:"type:text" json:"content"`
	Category       string `gorm:"size:255;not null" json:"category"`
	Author         string `gorm:"size:255" json:"author"`
	PublicationDate string `gorm:"size:100" json:"publication_date"`
}