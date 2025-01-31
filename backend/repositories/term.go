package repositories

import (
	"chukotka/core/db"
	"chukotka/models"
	"gorm.io/gorm"
)

// GetAllTerms - возвращает все термины
func GetAllTerms() ([]models.Term, error) {
	var terms []models.Term
	if err := db.DB.Find(&terms).Error; err != nil {
		return nil, err
	}
	return terms, nil
}

// GetTermByID - возвращает термин по ID
func GetTermByID(id string) (*models.Term, error) {
	var term models.Term
	if err := db.DB.First(&term, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &term, nil
}

// CreateTerm - создает новый термин
func CreateTerm(term *models.Term) error {
	return db.DB.Create(term).Error
}

// UpdateTerm - обновляет данные термина
func UpdateTerm(term *models.Term) error {
	return db.DB.Save(term).Error
}

// DeleteTerm - удаляет термин по ID
func DeleteTerm(id string) error {
	var term models.Term
	if err := db.DB.First(&term, id).Error; err != nil {
		return err
	}
	return db.DB.Unscoped().Delete(&term).Error
}
