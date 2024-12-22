package server

import (
	"log"
	"net/http"
	"calc_service/internal/handler"
)

func Start() {
	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}