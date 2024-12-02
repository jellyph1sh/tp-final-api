package dbmodel

import (
	"dapi-tpfinal-s2/pkg/model"

	"gorm.io/gorm"
)

type VisitEntry struct {
	gorm.Model
	CatID      int `gorm:"column:cat_id"`
	Cat        CatEntry
	Date       string            `gorm:"column:date"`
	Reason     string            `gorm:"column:reason"`
	Doctor     string            `gorm:"column:doctor"`
	Treatments []*TreatmentEntry `gorm:"foreignKey:VisitID"`
}

type VisitEntryRepository interface {
	Create(entry *VisitEntry) (*VisitEntry, error)
	FindAll(date string, doctor string, reason string) ([]*VisitEntry, error)
	FindById(id int) (*VisitEntry, error)
	FindAllByCatId(catId int, date string, doctor string, reason string) ([]*VisitEntry, error)
	Update(entry *VisitEntry) (*VisitEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *VisitEntry) *model.VisitResponse
	ToHistoryModel(entry *VisitEntry) *model.VisitHistoryResponse
}

type visitEntryRepository struct {
	db *gorm.DB
}

func NewVisitEntryRepository(db *gorm.DB) VisitEntryRepository {
	return &visitEntryRepository{db: db}
}

func (r *visitEntryRepository) Create(entry *VisitEntry) (*VisitEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *visitEntryRepository) FindAll(date string, doctor string, reason string) ([]*VisitEntry, error) {
	query := r.db

	if date != "" {
		query = query.Where("date = ?", date)
	}
	if doctor != "" {
		query = query.Where("doctor = ?", doctor)
	}
	if reason != "" {
		query = query.Where("reason = ?", reason)
	}

	var entries []*VisitEntry
	if err := query.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *visitEntryRepository) FindById(id int) (*VisitEntry, error) {
	var entry *VisitEntry
	if err := r.db.Where("id = ?", id).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *visitEntryRepository) FindAllByCatId(catId int, date string, doctor string, reason string) ([]*VisitEntry, error) {
	query := r.db.Where("cat_id = ?", catId)

	if date != "" {
		query = query.Where("date = ?", date)
	}
	if doctor != "" {
		query = query.Where("doctor = ?", doctor)
	}
	if reason != "" {
		query = query.Where("reason = ?", reason)
	}

	var entries []*VisitEntry
	if err := query.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *visitEntryRepository) Update(entry *VisitEntry) (*VisitEntry, error) {
	if err := r.db.Where("id = ?", entry.ID).Updates(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *visitEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&VisitEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *visitEntryRepository) ToModel(entry *VisitEntry) *model.VisitResponse {
	return &model.VisitResponse{
		ID:     int(entry.ID),
		CatID:  entry.CatID,
		Date:   entry.Date,
		Reason: entry.Reason,
		Doctor: entry.Doctor,
	}
}

func (r *visitEntryRepository) ToHistoryModel(entry *VisitEntry) *model.VisitHistoryResponse {
	return &model.VisitHistoryResponse{
		ID:         int(entry.ID),
		CatID:      entry.CatID,
		Date:       entry.Date,
		Reason:     entry.Reason,
		Doctor:     entry.Doctor,
		Treatments: nil,
	}
}
