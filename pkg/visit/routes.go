package visit

import (
	"dapi-tpfinal-s2/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	visitConfig := New(configuration)
	router := chi.NewRouter()

	router.Get("/", visitConfig.GetAllVisitsHandler)
	router.Get("/{id}", visitConfig.GetVisitByIdHandler)
	router.Post("/", visitConfig.CreateVisitHandler)
	router.Put("/{id}", visitConfig.UpdateVisitHandler)
	router.Delete("/{id}", visitConfig.DeleteVisitHandler)
	router.Get("/{id}/treatments", visitConfig.GetTreatmentsByVisitHandler)
	router.Get("/{id}/treatments/{treatmentId}", visitConfig.GetTreatmentByVisitHandler)

	return router

}
