package handler

import (
	"github.com/fatma0306/feedback/database"
	"github.com/gorilla/mux"
)

// Router is function to exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/review", database.GetReview()).Methods("GET", "OPTIONS")

	return router
}
