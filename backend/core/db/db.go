package db

import (
	"fmt"
	"os"
	"sync"

	"chukotka/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, databaseName, port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	fmt.Println("Successfully connected to the database")
}

func Connect() {
	once.Do(connect)
}

func Migrate() error {
	err := DB.AutoMigrate(
		&models.Term{},
		&models.Admin{},
		&models.District{},
	)
	if err != nil {
		return fmt.Errorf("migrate base tables: %w", err)
	}

	err = DB.AutoMigrate(
		&models.Village{},
		&models.Material{},
		&models.AboutPage{},
	)
	if err != nil {
		return fmt.Errorf("migrate dependent tables: %w", err)
	}
	return nil
}

// RunAllSeeds выполняет все сиды (админ, about, карта). Нужен вызов Connect ранее.
func RunAllSeeds() {
	SeedAdmin()
	SeedAbout()
	SeedMapData()
}

func InitDatabase() {
	Connect()
	if err := Migrate(); err != nil {
		panic(fmt.Sprintf("Failed to migrate: %v", err))
	}
	RunAllSeeds()
}
