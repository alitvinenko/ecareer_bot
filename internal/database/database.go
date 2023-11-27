package database

import (
	cmModel "github.com/alitvinenko/ecareer_bot/internal/repository/club_member/model"
	pModel "github.com/alitvinenko/ecareer_bot/internal/repository/profile/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("error on init DB. Path: %s. Error: %v", path, err)
	}

	err = db.AutoMigrate(&cmModel.ClubMember{})
	err = db.AutoMigrate(&pModel.Profile{})
	if err != nil {
		log.Fatalf("error on automigrate: %v", err)
	}

	return db
}
