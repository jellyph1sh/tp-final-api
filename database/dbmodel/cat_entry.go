package dbmodel

import (
	"dapi-tpfinal-s2/pkg/model"

	"gorm.io/gorm"
)

type CatEntry struct {
	gorm.Model
	Name   string  `gorm:"column:name"`
	Age    int     `gorm:"column:age"`
	Race   string  `gorm:"column:race"`
	Gender string  `gorm:"column:gender"`
	Weight float32 `gorm:"column:weight"`

	Visits []*VisitEntry `gorm:"foreignKey:CatID"`
}

type CatEntryRepository interface {
	Create(entry *CatEntry) (*CatEntry, error)
	FindAll() ([]*CatEntry, error)
	FindById(id int) (*CatEntry, error)
	Update(entry *CatEntry) (*CatEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *CatEntry) *model.CatResponse
}

type catEntryRepository struct {
	db *gorm.DB
}

func NewCatEntryRepository(db *gorm.DB) CatEntryRepository {
	return &catEntryRepository{db: db}
}

func (r *catEntryRepository) Create(entry *CatEntry) (*CatEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *catEntryRepository) FindAll() ([]*CatEntry, error) {
	var entries []*CatEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *catEntryRepository) FindById(id int) (*CatEntry, error) {
	var entry *CatEntry
	if err := r.db.Where("id = ?", id).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *catEntryRepository) Update(entry *CatEntry) (*CatEntry, error) {
	if err := r.db.Where("id = ?", entry.ID).Updates(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *catEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&CatEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *catEntryRepository) ToModel(entry *CatEntry) *model.CatResponse {
	return &model.CatResponse{
		ID:     int(entry.ID),
		Name:   entry.Name,
		Age:    entry.Age,
		Race:   entry.Race,
		Gender: entry.Gender,
		Weight: entry.Weight,
	}
}
