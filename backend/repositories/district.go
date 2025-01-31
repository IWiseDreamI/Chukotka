package repositories

import (
	"errors"
	"gorm.io/gorm"

	"chukotka/core/db"
	"chukotka/models"
)


// GetAllDistricts возвращает все районы
func GetAllDistricts() ([]models.District, error) {
	var districts []models.District
	if err := db.DB.Preload("Villages").Find(&districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

// GetDistrictByID возвращает район по ID
func GetDistrictByID(id string) (*models.District, error) {
	var district models.District
	if err := db.DB.Preload("Villages").First(&district, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &district, nil
}

// CreateDistrict создает новый район
func CreateDistrict(district *models.District) error {
	return db.DB.Create(district).Error
}

// UpdateDistrict обновляет данные района
func UpdateDistrict(id string, updatedDistrict *models.District) error {
	var district models.District
	if err := db.DB.First(&district, id).Error; err != nil {
		return err
	}
	return db.DB.Model(&district).Updates(updatedDistrict).Error
}

// DeleteDistrict удаляет район по ID
func DeleteDistrict(id string) error {
	var district models.District
	if err := db.DB.First(&district, id).Error; err != nil {
		return err
	}
	return db.DB.Unscoped().Delete(&district).Error
}