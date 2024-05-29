package routes

import (
	"todo-api/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/api/v1/resource", handlers.ResourceHandler).Methods("GET")
}
