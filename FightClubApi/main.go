package main

import (
	"encoding/json"
	"fight_club/entities"
	"fight_club/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	model := entities.User{Id: 1, Name: "Kamaru", LastName: "Usman", NickName: "The Nigerian Nightmare"}
	if err := json.NewEncoder(w).Encode(model); err != nil {
		fmt.Println(err)
	}
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", handler).Methods(http.MethodGet)

	// Fighters handler
	fightersHandler := handlers.NewFightersHandler()
	r.HandleFunc("/api/fighters/list", fightersHandler.GetFighters).Methods(http.MethodGet)
	r.HandleFunc("/api/fighter/{id}", fightersHandler.GetFighter).Methods(http.MethodGet)

	r.Use(JSONMiddleware)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("server failed: %v", err)
	}
}
