package main

import (
	"fmt"
	"net/http"
	"errors"
	"encoding/json"
)

func main() {
	router := http.NewServeMux()
	//handle hello endpoint
	//w - response writer, r - request
	router.HandleFunc("GET /api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
	})

	router.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})

	//listen on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		err = errors.New("failed to start server: " + err.Error())
		fmt.Println(err)
	}
}

