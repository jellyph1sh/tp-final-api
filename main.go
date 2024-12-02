package main

import (
	"dapi-tpfinal-s2/config"
	"dapi-tpfinal-s2/pkg/cat"
	"dapi-tpfinal-s2/pkg/treatment"
	"dapi-tpfinal-s2/pkg/visit"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/api/v1/cats", cat.Routes(configuration))
	router.Mount("/api/v1/visits", visit.Routes(configuration))
	router.Mount("/api/v1/treatments", treatment.Routes(configuration))

	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Fatal(http.ListenAndServe(":8080", router))
}
