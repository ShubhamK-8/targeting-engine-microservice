package main

import (
	"log"
	"net/http"

	"targeting-engine/internal/delivery"
)

func main() {
	http.HandleFunc("/v1/delivery", delivery.HandleDelivery)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
