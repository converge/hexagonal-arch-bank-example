package main

import (
	"github.com/gorilla/mux"
	"hexagonal-example/internal/core/domain/health"
	"hexagonal-example/internal/core/services"
	"hexagonal-example/internal/handlers"
	"hexagonal-example/internal/repositories"
	"log"
	"net/http"
)

func main() {

	dbRepository := repositories.NewMemoryDb()

	srv := services.New(dbRepository)

	apiHandler := handlers.NewHTTPHandler(srv)

	x := health.Health{
		Id: "1",
	}
	healthService := services.HealthService{
		Health: x,
	}
	healthHandler := handlers.HTTPHealthHandler{
		HealthService: healthService,
	}

	r := mux.NewRouter()
	r.HandleFunc("/create", apiHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/balance/{id}", apiHandler.Balance).Methods(http.MethodGet)
	r.HandleFunc("/health", healthHandler.HealthCheck).Methods(http.MethodGet)
	log.Println("listening at port 7000...")
	panic(http.ListenAndServe(":7000", r))

}
