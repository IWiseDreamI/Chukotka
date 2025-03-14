package repositories

import (
	"chukotka/core/db"
	"chukotka/models"
	"errors"
)

func GetAllMaterials() ([]models.Material, error) {
	var materials []models.Material
	result := db.DB.Find(&materials)
	return materials, result.Error
}

func GetMaterialByID(id uint) (models.Material, error) {
	var material models.Material
	result := db.DB.First(&material, id)
	if result.Error != nil {
		return material, errors.New("material not found")
	}
	return material, nil
}

func CreateMaterial(material *models.Material) error {
	return db.DB.Create(material).Error
}

func UpdateMaterial(material *models.Material) error {
	return db.DB.Save(material).Error
}

func DeleteMaterial(id uint) error {
	result := db.DB.Delete(&models.Material{}, id)
	if result.RowsAffected == 0 {
		return errors.New("material not found")
	}
	return nil
}
