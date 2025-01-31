package models

import "gorm.io/gorm"

// Admin - модель администратора
type Admin struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
}