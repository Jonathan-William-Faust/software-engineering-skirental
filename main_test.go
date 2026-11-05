package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	newServer().Handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if contentType := response.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected JSON content type, got %q", contentType)
	}

	if body := response.Body.String(); body != "{\"status\":\"ok\"}\n" {
		t.Fatalf("unexpected response body: %q", body)
	}
}

func TestSeededServerStartsWithDefaultSkiCatalog(t *testing.T) {
	server, err := newSeededServer()
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/api/skis", nil)
	response := httptest.NewRecorder()
	server.Handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
	var skis []struct {
		Model string
		Size  int
	}
	if err := json.NewDecoder(response.Body).Decode(&skis); err != nil {
		t.Fatal(err)
	}
	if len(skis) != 100 {
		t.Fatalf("seeded ski count = %d, want 100", len(skis))
	}
}

func TestSeededServerStartsWithDefaultPerson(t *testing.T) {
	server, err := newSeededServer()
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/api/people", nil)
	response := httptest.NewRecorder()
	server.Handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
	var people []struct {
		ID        string
		FirstName string
		LastName  string
		BirthDate string
	}
	if err := json.NewDecoder(response.Body).Decode(&people); err != nil {
		t.Fatal(err)
	}
	foundMax := false
	for _, person := range people {
		if person.ID == "P-0001" && person.FirstName == "Max" && person.LastName == "Mustermann" {
			foundMax = true
		}
	}
	if len(people) != 2 || !foundMax {
		t.Fatalf("unexpected seeded people: %#v", people)
	}
	var maxBirthDate string
	for _, person := range people {
		if person.ID == "P-0001" {
			maxBirthDate = person.BirthDate
		}
	}
	if !strings.HasPrefix(maxBirthDate, "2000-01-01") {
		t.Fatalf("unexpected default birth date: %q", maxBirthDate)
	}
}
