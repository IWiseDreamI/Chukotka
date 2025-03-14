package db

import (
	"chukotka/models"
	"chukotka/utils"
	"log"
	"os"
)

func SeedAdmin() {
	var count int64
	DB.Model(&models.Admin{}).Where("username = ?", "admin").Count(&count)
	if count > 0 {
		log.Println("Admin already exists. Skipping seeding.")
		return
	}

	hashedPassword, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin := models.Admin{
		Username: "admin",
		Password: hashedPassword,
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to seed admin: %v", err)
	}

	log.Println("Admin seeded successfully.")
}


func SeedAbout() {
	var count int64
	DB.Model(&models.AboutPage{}).Count(&count)
	if count > 0 {
		log.Println("About page already exists. Skipping seeding.")
		return
	}

	about := models.AboutPage{
		Content: "Здесь можно разместить информацию о проекте.",
	}

	if err := DB.Create(&about).Error; err != nil {
		log.Fatalf("Failed to seed about page: %v", err)
	}

	log.Println("About seeded successfully.")
}
