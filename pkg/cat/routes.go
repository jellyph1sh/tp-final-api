package cat

import (
	"dapi-tpfinal-s2/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	catConfig := New(configuration)
	router := chi.NewRouter()

	router.Get("/", catConfig.GetAllCatsHandler)
	router.Get("/{id}", catConfig.GetCatByIdHandler)
	router.Post("/", catConfig.CreateCatHandler)
	router.Put("/{id}", catConfig.UpdateCatHandler)
	router.Delete("/{id}", catConfig.DeleteCatHandler)
	router.Get("/{id}/history", catConfig.GetHistoryByCatHandler)
	router.Get("/{id}/visits", catConfig.GetVisitsByCatHandler)
	router.Get("/{id}/visits/{visitId}", catConfig.GetVisitByCatHandler)
	router.Get("/{id}/visits/{visitId}/treatments", catConfig.GetTreatmentsByCatByVisitHandler)
	router.Get("/{id}/visits/{visitId}/treatments/{treatmentId}", catConfig.GetTreatmentByCatByVisitHandler)

	return router
}
