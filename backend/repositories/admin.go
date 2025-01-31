package repositories

import (
	"chukotka/core/db"
	"chukotka/models"
	"gorm.io/gorm"
)

// CreateAdmin создает нового администратора
func CreateAdmin(admin *models.Admin) error {
	return db.DB.Create(admin).Error
}

// FindAdminByUsername ищет администратора по имени пользователя
func FindAdminByUsername(username string) (*models.Admin, error) {
	var admin models.Admin
	if err := db.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}
