package repositories

import (
	"chukotka/core/db"
	"chukotka/models"
	"gorm.io/gorm"
)

// GetAllVillages - возвращает все деревни
func GetAllVillages() ([]models.Village, error) {
	var villages []models.Village
	if err := db.DB.Find(&villages).Error; err != nil {
		return nil, err
	}
	return villages, nil
}

// GetVillageByID - возвращает деревню по ID
func GetVillageByID(id string) (*models.Village, error) {
	var village models.Village
	if err := db.DB.First(&village, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &village, nil
}

// CreateVillage - создает новую деревню
func CreateVillage(village *models.Village) error {
	return db.DB.Create(village).Error
}

// UpdateVillage - обновляет данные деревни
func UpdateVillage(village *models.Village) error {
	return db.DB.Save(village).Error
}

// DeleteVillage - удаляет деревню по ID
func DeleteVillage(id string) error {
	var village models.Village
	if err := db.DB.First(&village, id).Error; err != nil {
		return err
	}
	return db.DB.Unscoped().Delete(&village).Error
}