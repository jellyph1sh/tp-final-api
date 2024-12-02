package treatment

import (
	"dapi-tpfinal-s2/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	treatmentConfig := New(configuration)
	router := chi.NewRouter()

	router.Get("/", treatmentConfig.GetAllTreatmentsHandler)
	router.Get("/{id}", treatmentConfig.GetTreatmentByIdHandler)
	router.Post("/", treatmentConfig.CreateTreatmentHandler)
	router.Put("/{id}", treatmentConfig.UpdateTreatmentHandler)
	router.Delete("/{id}", treatmentConfig.DeleteTreatmentHandler)

	return router

}
