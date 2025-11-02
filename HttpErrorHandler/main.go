package main

import (
	"http_error_handler/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// handlers
	errorHandler := handlers.NewErrorHandler()
	homeApi := handlers.NewHomeApi(&errorHandler)
	r := mux.NewRouter()
	r.HandleFunc("/hello/{id}", homeApi.Hello).Methods("GET")

	// Custom 404 handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorHandler.NotFound(w, "Resource not found")
	})

	// Custom 405 handler (method not allowed), optional
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorHandler.MethodNotAllowed(w)
	})

	http.ListenAndServe(":8080", r)
}
