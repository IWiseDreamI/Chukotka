package models

import "gorm.io/gorm"

type District struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	Population string `gorm:"type:text;" json:"population"`
	Information string `gorm:"type:text;" json:"information"`
	Map string `gorm:"type:text;" json:"map"`
	Villages []Village `gorm:"foreignKey:DistrictID" json:"villages"`
}
