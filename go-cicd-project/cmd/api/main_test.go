package main

import (
	"testing"
	"github.com/alfacentaur1/go_backend/internal/handler"
	"net/http/httptest"
	"net/http"
)

func TestHelloHandler(t *testing.T) {
	// mock a new HTTP request to the hello endpoint
	req, err := http.NewRequest("GET", "/api/v1/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler.HelloHandler(response, req)
	// check if the status code is 200 OK
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	
}

func TestHealthHandler(t *testing.T) {
	// mock a new HTTP request to the health endpoint
	req, err := http.NewRequest("GET", "/api/v1/health", nil)	
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler.HealthHandler(response, req)
	// check if the status code is 200 OK
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}