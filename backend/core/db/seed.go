package db

import (
	"chukotka/models"
	"chukotka/utils"
	_ "embed"
	"encoding/json"
	"log"
	"os"
)

//go:embed seed_geo.json
var seedGeoJSON []byte

type seedGeoFile struct {
	Districts []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"districts"`
	Villages []struct {
		Name     string `json:"name"`
		District string `json:"district"`
		Cx       string `json:"cx"`
		Cy       string `json:"cy"`
	} `json:"villages"`
}

// SeedMapData создаёт районы и сёла с именами, совпадающими с frontend/src/assets/data.json (ключи map.districts и map.villages).
// Без этого GET /districts и GET /villages пустые — на главной не рисуются path и точки.
func SeedMapData() {
	var parsed seedGeoFile
	if err := json.Unmarshal(seedGeoJSON, &parsed); err != nil {
		log.Fatalf("seed map geo: parse json: %v", err)
	}

	for _, d := range parsed.Districts {
		var cnt int64
		DB.Model(&models.District{}).Where("name = ?", d.Name).Count(&cnt)
		if cnt > 0 {
			continue
		}
		rec := models.District{
			Name:        d.Name,
			Description: d.Description,
		}
		if err := DB.Create(&rec).Error; err != nil {
			log.Fatalf("seed map geo: create district %q: %v", d.Name, err)
		}
		log.Printf("Seeded district: %s", d.Name)
	}

	districtIDs := make(map[string]uint)
	for _, d := range parsed.Districts {
		var row models.District
		res := DB.Where("name = ?", d.Name).Limit(1).Find(&row)
		if res.Error != nil {
			log.Fatalf("seed map geo: district %q: %v", d.Name, res.Error)
		}
		if res.RowsAffected == 0 {
			log.Fatalf("seed map geo: district %q missing after seed", d.Name)
		}
		districtIDs[d.Name] = row.ID
	}

	for _, v := range parsed.Villages {
		var cnt int64
		DB.Model(&models.Village{}).Where("name = ?", v.Name).Count(&cnt)
		if cnt > 0 {
			continue
		}
		did, ok := districtIDs[v.District]
		if !ok {
			log.Fatalf("seed map geo: unknown district %q for village %q", v.District, v.Name)
		}
		vv := models.Village{
			Name:        v.Name,
			DistrictID:  did,
			Coordinates: v.Cx + "," + v.Cy,
		}
		if err := DB.Create(&vv).Error; err != nil {
			log.Fatalf("seed map geo: create village %q: %v", v.Name, err)
		}
		log.Printf("Seeded village: %s", v.Name)
	}
	log.Println("Map geo seed finished.")
}

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
