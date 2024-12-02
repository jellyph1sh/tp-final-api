package cat

import (
	"dapi-tpfinal-s2/config"
	"dapi-tpfinal-s2/database/dbmodel"
	"dapi-tpfinal-s2/helper"
	"dapi-tpfinal-s2/pkg/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CatConfig struct {
	*config.Config
}

func New(configuration *config.Config) *CatConfig {
	return &CatConfig{configuration}
}

func (config *CatConfig) CreateCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	catEntry := &dbmodel.CatEntry{Name: req.Name, Age: req.Age, Race: req.Race, Gender: req.Gender, Weight: req.Weight}
	config.CatEntryRepository.Create(catEntry)

	render.JSON(w, r, config.CatEntryRepository.ToModel(catEntry))
}

func (config *CatConfig) GetAllCatsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve all cats", http.StatusInternalServerError)
		return
	}

	responseEntries := make([]*model.CatResponse, len(entries))

	for i, entry := range entries {
		responseEntries[i] = config.CatEntryRepository.ToModel(entry)
	}

	render.JSON(w, r, responseEntries)
}

func (config *CatConfig) GetCatByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.CatEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve a cat on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.CatEntryRepository.ToModel(entry))
}

func (config *CatConfig) UpdateCatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.CatEntryRepository.FindById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve a cat on this id", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Cannot decode body", http.StatusInternalServerError)
		return
	}

	helper.ApplyChanges(data, entry)

	entry, err = config.CatEntryRepository.Update(entry)
	if err != nil {
		http.Error(w, "Failed to update cat on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, config.CatEntryRepository.ToModel(entry))
}

func (config *CatConfig) DeleteCatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	valid, err := config.CatEntryRepository.Delete(id)
	if err != nil {
		http.Error(w, "Failed to retrieve visits for this catId", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Cat does not exist", http.StatusNotFound)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Cat deleted"})
}

func (config *CatConfig) GetHistoryByCatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	date := r.URL.Query().Get("date")
	doctor := r.URL.Query().Get("doctor")
	reason := r.URL.Query().Get("reason")

	visitEntries, err := config.VisitEntryRepository.FindAllByCatId(id, date, doctor, reason)
	if err != nil {
		http.Error(w, "Failed to retrieve visits for this catId", http.StatusInternalServerError)
		return
	}

	VisitHistoryResponseEntries := make([]*model.VisitHistoryResponse, len(visitEntries))

	for i, visitEntry := range visitEntries {
		VisitHistoryResponseEntries[i] = config.VisitEntryRepository.ToHistoryModel(visitEntry)
	}

	for i, entry := range VisitHistoryResponseEntries {
		treatmentEntries, err := config.TreatmentEntryRepository.FindAllByVisitId(int(entry.ID))
		if err != nil {
			http.Error(w, "Failed to retrieve treatments for this visitId", http.StatusInternalServerError)
			return
		}

		treatmentResponseEntries := make([]*model.TreatmentResponse, len(treatmentEntries))

		for i, treatmentEntry := range treatmentEntries {
			treatmentResponseEntries[i] = config.TreatmentEntryRepository.ToModel(treatmentEntry)
		}

		VisitHistoryResponseEntries[i].Treatments = treatmentResponseEntries
	}

	render.JSON(w, r, VisitHistoryResponseEntries)
}

func (config *CatConfig) GetVisitsByCatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	date := r.URL.Query().Get("date")
	doctor := r.URL.Query().Get("doctor")
	reason := r.URL.Query().Get("reason")

	visitEntries, err := config.VisitEntryRepository.FindAllByCatId(id, date, doctor, reason)
	if err != nil {
		http.Error(w, "Failed to retrieve visits for this catId", http.StatusInternalServerError)
		return
	}

	VisitHistoryResponseEntries := make([]*model.VisitHistoryResponse, len(visitEntries))

	for i, visitEntry := range visitEntries {
		VisitHistoryResponseEntries[i] = config.VisitEntryRepository.ToHistoryModel(visitEntry)
	}

	for i, entry := range VisitHistoryResponseEntries {
		treatmentEntries, err := config.TreatmentEntryRepository.FindAllByVisitId(int(entry.ID))
		if err != nil {
			http.Error(w, "Failed to retrieve treatments for this visitId", http.StatusInternalServerError)
			return
		}

		treatmentResponseEntries := make([]*model.TreatmentResponse, len(treatmentEntries))

		for i, treatmentEntry := range treatmentEntries {
			treatmentResponseEntries[i] = config.TreatmentEntryRepository.ToModel(treatmentEntry)
		}

		VisitHistoryResponseEntries[i].Treatments = treatmentResponseEntries
	}

	render.JSON(w, r, VisitHistoryResponseEntries)
}

func (config *CatConfig) GetVisitByCatHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	visitIdStr := chi.URLParam(r, "visitId")
	visitId, err := strconv.Atoi(visitIdStr)
	if err != nil || visitId < 0 {
		http.Error(w, "Invalid visitId parameter", http.StatusBadRequest)
		return
	}

	entry, err := config.VisitEntryRepository.FindById(visitId)
	if err != nil {
		http.Error(w, "Failed to retrieve a visit on this id", http.StatusInternalServerError)
		return
	}

	if entry.CatID != id {
		http.Error(w, "Visit does not belong to this cat", http.StatusUnauthorized)
		return
	}

	historyEntry := config.VisitEntryRepository.ToHistoryModel(entry)

	treatmentEntries, err := config.TreatmentEntryRepository.FindAllByVisitId(int(historyEntry.ID))
	if err != nil {
		http.Error(w, "Failed to retrieve treatments for this visitId", http.StatusInternalServerError)
		return
	}

	treatmentResponseEntries := make([]*model.TreatmentResponse, len(treatmentEntries))

	for i, treatmentEntry := range treatmentEntries {
		treatmentResponseEntries[i] = config.TreatmentEntryRepository.ToModel(treatmentEntry)
	}

	historyEntry.Treatments = treatmentResponseEntries

	render.JSON(w, r, historyEntry)
}

func (config *CatConfig) GetTreatmentsByCatByVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitIdStr := chi.URLParam(r, "visitId")
	visitId, err := strconv.Atoi(visitIdStr)
	if err != nil || visitId < 0 {
		http.Error(w, "Invalid visitId parameter", http.StatusBadRequest)
		return
	}

	entries, err := config.TreatmentEntryRepository.FindAllByVisitId(visitId)
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

func (config *CatConfig) GetTreatmentByCatByVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitIdStr := chi.URLParam(r, "visitId")
	visitId, err := strconv.Atoi(visitIdStr)
	if err != nil || visitId < 0 {
		http.Error(w, "Invalid visitId parameter", http.StatusBadRequest)
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

	if entry.VisitID != visitId {
		http.Error(w, "Treatment does not belong to this visit", http.StatusUnauthorized)
		return
	}

	render.JSON(w, r, config.TreatmentEntryRepository.ToModel(entry))
}
