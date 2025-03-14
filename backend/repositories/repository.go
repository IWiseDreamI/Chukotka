package repository

import (
	"chukotka/core/db"
	"chukotka/models"
)

func GetAboutPage() (models.AboutPage, error) {
	var about models.AboutPage
	result := db.DB.First(&about)
	return about, result.Error
}

func UpdateAboutPage(about *models.AboutPage) error {
	var existing models.AboutPage
	db.DB.First(&existing)
	existing.Content = about.Content
	return db.DB.Save(&existing).Error
}
