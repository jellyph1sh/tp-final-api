package database

import (
	"dapi-tpfinal-s2/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&dbmodel.CatEntry{})
	db.AutoMigrate(&dbmodel.VisitEntry{})
	db.AutoMigrate(&dbmodel.TreatmentEntry{})
	log.Println("Database migrated successfully")
}
