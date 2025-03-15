package repositories

import (
	"chukotka/core/db"
	"chukotka/models"
	"gorm.io/gorm"
)

func GetAboutPage() (models.AboutPage, error) {
	var about models.AboutPage
	result := db.DB.First(&about)
	return about, result.Error
}

// Если запись не найдена, можно создать новую.
func UpdateAboutPage(about *models.AboutPage) error {
	var existing models.AboutPage
	result := db.DB.First(&existing)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Если запись не найдена, создаём новую запись
			return db.DB.Create(about).Error
		}
		return result.Error
	}
	existing.Content = about.Content
	return db.DB.Save(&existing).Error
}
