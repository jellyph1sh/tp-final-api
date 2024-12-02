package visit

import (
	"dapi-tpfinal-s2/config"
	"dapi-tpfinal-s2/database/dbmodel"
	"dapi-tpfinal-s2/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VisitConfig struct {
	*config.Config
}

func New(configuration *config.Config) *VisitConfig {
	return &VisitConfig{configuration}
}

func (config *VisitConfig) CreateVisitHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	visitEntry := &dbmodel.VisitEntry{Date: req.Date, Reason: req.Reason, CatID: req.CatID, Doctor: req.Doctor}
	config.VisitEntryRepository.Create(visitEntry)

	render.JSON(w, r, config.VisitEntryRepository.ToModel(visitEntry))
}

func (config *VisitConfig) GetAllVisitsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.VisitEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve all visits", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.VisitResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.VisitEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

func (config *VisitConfig) GetVisitByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.VisitEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve a visit on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.VisitEntryRepository.ToModel(entry))
}

func (config *VisitConfig) UpdateVisitHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.VisitEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve a visit on this id", http.StatusInternalServerError)
		return
	}

	req := &model.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	entry.Date = req.Date
	entry.Reason = req.Reason
	entry.CatID = req.CatID
	entry.Doctor = req.Doctor
	config.VisitEntryRepository.Update(entry)

	render.JSON(w, r, config.VisitEntryRepository.ToModel(entry))
}

func (config *VisitConfig) DeleteVisitHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if _, err := config.VisitEntryRepository.Delete(id); err != nil {
		http.Error(w, "Failed to delete a visit on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"message": "Visit deleted"})
}

func (config *VisitConfig) GetTreatmentsByVisitHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entries, err := config.TreatmentEntryRepository.FindAllByVisitId(id)
	if err != nil {
		http.Error(w, "Failed to retrieve treatments for this visitId", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.TreatmentResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.TreatmentEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

func (config *VisitConfig) GetTreatmentByVisitHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	treatmentIdStr := chi.URLParam(r, "treatmentId")
	treatmentId, err := strconv.Atoi(treatmentIdStr)
	if err != nil || treatmentId < 0 {
		http.Error(w, "Invalid treatmentId parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.TreatmentEntryRepository.FindById(treatmentId)
	if err != nil {
		http.Error(w, "Failed to retrieve a visit on this id", http.StatusInternalServerError)
		return
	}

	if entry.VisitID != id {
		http.Error(w, "Treatment does not belong to this visit", http.StatusUnauthorized)
		return
	}

	render.JSON(w, r, config.TreatmentEntryRepository.ToModel(entry))
}
