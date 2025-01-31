package db

import (
    "os"
    "fmt"
    "gorm.io/gorm"
	"chukotka/models"
    "gorm.io/driver/postgres" 
)

var DB *gorm.DB

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
        panic(err)
    } else {
        fmt.Println("Successfully connected to the database")
    }
}

func InitDatabase() {
	connect()
	
    DB.AutoMigrate(&models.Term{})
	DB.AutoMigrate(&models.Admin{})
    DB.AutoMigrate(&models.Village{})
	DB.AutoMigrate(&models.District{})
	
    SeedAdmin()
}