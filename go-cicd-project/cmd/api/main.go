package main

import (
	"fmt"
	"net/http"
	"github.com/alfacentaur1/go_backend/internal/handler"
)

func main() {
	router := http.NewServeMux()
	//handle hello endpoint
	//w - response writer, r - request
	router.HandleFunc("GET /api/v1/hello", handler.HelloHandler)

	router.HandleFunc("GET /api/v1/health", handler.HealthHandler)

	//listen on port 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
	
}

