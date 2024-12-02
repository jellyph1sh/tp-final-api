package config

import (
	"dapi-tpfinal-s2/database"
	"dapi-tpfinal-s2/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	CatEntryRepository       dbmodel.CatEntryRepository
	VisitEntryRepository     dbmodel.VisitEntryRepository
	TreatmentEntryRepository dbmodel.TreatmentEntryRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(sqlite.Open("vet-clinic.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	config.CatEntryRepository = dbmodel.NewCatEntryRepository(databaseSession)
	config.VisitEntryRepository = dbmodel.NewVisitEntryRepository(databaseSession)
	config.TreatmentEntryRepository = dbmodel.NewTreatmentEntryRepository(databaseSession)

	return &config, nil
}
