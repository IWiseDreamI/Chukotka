package db

import (
    "os"
    "fmt"
    "sync"
    "gorm.io/gorm"
    "chukotka/models"
    "gorm.io/driver/postgres"
)

var (
    DB   *gorm.DB
    once sync.Once
)

func connect() {
    var err error
    host := os.Getenv("DB_HOST")
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASS")
    databaseName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(fmt.Sprintf("Failed to connect to the database: %v", err))
    } else {
        fmt.Println("Successfully connected to the database")
    }
}

func InitDatabase() {
	once.Do(connect)

	err := DB.AutoMigrate(
		&models.Term{},
		&models.Admin{},
		&models.Village{},
		&models.District{},
		&models.Material{},
		&models.AboutPage{},
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	SeedAdmin()
    SeedAbout()
}
