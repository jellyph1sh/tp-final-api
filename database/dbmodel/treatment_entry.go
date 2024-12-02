package dbmodel

import (
	"dapi-tpfinal-s2/pkg/model"

	"gorm.io/gorm"
)

type TreatmentEntry struct {
	gorm.Model
	VisitID   int `gorm:"column:visit_id"`
	Visit     VisitEntry
	Medicine  string `gorm:"column:medicine"`
	StartDate string `gorm:"column:start_date"`
	EndDate   string `gorm:"column:end_date"`
	DoctorTip string `gorm:"column:doctor_tip"`
}

type TreatmentEntryRepository interface {
	Create(entry *TreatmentEntry) (*TreatmentEntry, error)
	FindAll() ([]*TreatmentEntry, error)
	FindById(id int) (*TreatmentEntry, error)
	FindAllByVisitId(visitId int) ([]*TreatmentEntry, error)
	Update(entry *TreatmentEntry) (*TreatmentEntry, error)
	Delete(id int) (bool, error)
	ToModel(entry *TreatmentEntry) *model.TreatmentResponse
}

type treatmentEntryRepository struct {
	db *gorm.DB
}

func NewTreatmentEntryRepository(db *gorm.DB) TreatmentEntryRepository {
	return &treatmentEntryRepository{db: db}
}

func (r *treatmentEntryRepository) Create(entry *TreatmentEntry) (*TreatmentEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *treatmentEntryRepository) FindAll() ([]*TreatmentEntry, error) {
	var entries []*TreatmentEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *treatmentEntryRepository) FindById(id int) (*TreatmentEntry, error) {
	var entry *TreatmentEntry
	if err := r.db.Where("id = ?", id).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *treatmentEntryRepository) FindAllByVisitId(visitId int) ([]*TreatmentEntry, error) {
	var entries []*TreatmentEntry
	if err := r.db.Where("visit_id = ?", visitId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *treatmentEntryRepository) Update(entry *TreatmentEntry) (*TreatmentEntry, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *treatmentEntryRepository) Delete(id int) (bool, error) {
	if err := r.db.Delete(&TreatmentEntry{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *treatmentEntryRepository) ToModel(entry *TreatmentEntry) *model.TreatmentResponse {
	return &model.TreatmentResponse{
		ID:        int(entry.ID),
		VisitID:   entry.VisitID,
		Medicine:  entry.Medicine,
		StartDate: entry.StartDate,
		EndDate:   entry.EndDate,
		DoctorTip: entry.DoctorTip,
	}
}
