package treatment

import (
	"dapi-tpfinal-s2/config"
	"dapi-tpfinal-s2/database/dbmodel"
	"dapi-tpfinal-s2/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TreatmentConfig struct {
	*config.Config
}

func New(configuration *config.Config) *TreatmentConfig {
	return &TreatmentConfig{configuration}
}

func (config *TreatmentConfig) CreateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	treatmentEntry := &dbmodel.TreatmentEntry{
		VisitID:   req.VisitID,
		Medicine:  req.Medicine,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		DoctorTip: req.DoctorTip,
	}
	config.TreatmentEntryRepository.Create(treatmentEntry)

	render.JSON(w, r, config.TreatmentEntryRepository.ToModel(treatmentEntry))
}

func (config *TreatmentConfig) GetAllTreatmentsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.TreatmentEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve all treatments", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.TreatmentResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.TreatmentEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

func (config *TreatmentConfig) GetTreatmentByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.TreatmentEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve treatment", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.TreatmentEntryRepository.ToModel(entry))
}

func (config *TreatmentConfig) UpdateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.TreatmentEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to update treatment", http.StatusInternalServerError)
		return
	}

	req := &model.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	entry.VisitID = req.VisitID
	entry.Medicine = req.Medicine
	entry.StartDate = req.StartDate
	entry.EndDate = req.EndDate
	entry.DoctorTip = req.DoctorTip

	config.TreatmentEntryRepository.Update(entry)

	render.JSON(w, r, config.TreatmentEntryRepository.ToModel(entry))
}

func (config *TreatmentConfig) DeleteTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if _, err := config.VisitEntryRepository.Delete(id); err != nil {
		http.Error(w, "Failed to delete a treatment on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"message": "Treatment deleted"})
}
