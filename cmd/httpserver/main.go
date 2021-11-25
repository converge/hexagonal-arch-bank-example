package main

import (
	"github.com/gorilla/mux"
	"hexagonal-example/internal/core/services"
	"hexagonal-example/internal/handlers/bank-handler"
	"hexagonal-example/internal/repositories"
	"log"
	"net/http"
)

func main() {

	dbRepository := repositories.NewMemoryDb()

	srv := services.New(dbRepository)

	apiHandler := bank_handler.NewHTTPHandler(srv)

	r := mux.NewRouter()
	r.HandleFunc("/create", apiHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/balance/{id}", apiHandler.Balance).Methods(http.MethodGet)
	log.Println("listening at port 7000...")
	panic(http.ListenAndServe(":7000", r))

}