package main

import (
	"log"
	"net/http"
	"api-gateway/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	log.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", mux)
}