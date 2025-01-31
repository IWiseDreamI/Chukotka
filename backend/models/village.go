package models

import "gorm.io/gorm"

type Village struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	Contacts string `gorm:"type:text;" json:"contacts"`
	Nationality string `gorm:"type:text;" json:"nationality"`
	Population string `gorm:"type:text;" json:"population"`
	Organizations string `gorm:"type:text;" json:"organizations"`
	Transport string `gorm:"type:text;" json:"transport"`
	Connection string `gorm:"type:text;" json:"сonnection"`
	Activity string `gorm:"type:text;" json:"activity"`
	Information string `gorm:"type:text;" json:"information"`
	People string `gorm:"type:text;" json:"people"`
	Culture string `gorm:"type:text;" json:"culture"`
	Socials string `gorm:"type:text;" json:"socials"`
	DistrictID uint `json:"district_id"`
	Coordinates   string `gorm:"type:text;not null" json:"category"`
}