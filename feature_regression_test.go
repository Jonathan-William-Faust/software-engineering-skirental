package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFeatureRegressionMariaLoanAndPersonCity(t *testing.T) {
	server, err := newSeededServer()
	if err != nil {
		t.Fatal(err)
	}

	peopleResponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(peopleResponse, httptest.NewRequest(http.MethodGet, "/api/people", nil))
	if peopleResponse.Code != http.StatusOK {
		t.Fatalf("GET /api/people status = %d", peopleResponse.Code)
	}
	var people []struct {
		ID        string
		FirstName string
		LastName  string
		City      string
	}
	if err := json.NewDecoder(peopleResponse.Body).Decode(&people); err != nil {
		t.Fatal(err)
	}
	var mariaID string
	for _, person := range people {
		if person.FirstName == "Maria" && person.LastName == "Musterfrau" {
			mariaID = person.ID
			if person.City == "" {
				t.Fatal("Maria is missing city")
			}
		}
	}
	if mariaID == "" {
		t.Fatal("Maria Musterfrau was not seeded")
	}

	skisResponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(skisResponse, httptest.NewRequest(http.MethodGet, "/api/skis", nil))
	var skis []struct {
		ID     string
		Status string
	}
	if err := json.NewDecoder(skisResponse.Body).Decode(&skis); err != nil {
		t.Fatal(err)
	}
	var availableSkiID string
	for _, ski := range skis {
		if ski.Status == "verfügbar" {
			availableSkiID = ski.ID
			break
		}
	}
	if availableSkiID == "" {
		t.Fatal("no available ski for Maria loan")
	}

	loanBody := `{"ski_id":"` + availableSkiID + `","person_id":"` + mariaID + `"}`
	loanResponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(loanResponse, httptest.NewRequest(http.MethodPost, "/api/loans", strings.NewReader(loanBody)))
	if loanResponse.Code != http.StatusCreated {
		t.Fatalf("Maria loan status = %d, body = %s", loanResponse.Code, loanResponse.Body.String())
	}
}

func TestFeatureRegressionPersonValidationAndUpdate(t *testing.T) {
	server, err := newSeededServer()
	if err != nil {
		t.Fatal(err)
	}
	createBody := `{"first_name":"Test","last_name":"Person","birth_date":"05.11.1990","identification":"ID-TEST","country":"Deutschland","city":"Berlin","postal_code":"10115","street":"Teststraße","house_number":"1","mobile_number":"+491234567890","photo":{"content_type":"image/jpeg","data_url":"data:image/jpeg;base64,dGVzdA=="}}`
	createResponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(createResponse, httptest.NewRequest(http.MethodPost, "/api/people", strings.NewReader(createBody)))
	if createResponse.Code != http.StatusCreated {
		t.Fatalf("person creation status = %d, body = %s", createResponse.Code, createResponse.Body.String())
	}
	var person struct {
		ID   string
		City string
	}
	if err := json.NewDecoder(createResponse.Body).Decode(&person); err != nil {
		t.Fatal(err)
	}
	if person.City != "Berlin" {
		t.Fatalf("created city = %q", person.City)
	}
	updateBody := `{"first_name":"Test","last_name":"Person","birth_date":"05.11.1990","identification":"ID-TEST","country":"Österreich","city":"Wien","postal_code":"1010","street":"Teststraße","house_number":"2","mobile_number":"+436641234567","photo":{"content_type":"image/png","data_url":"data:image/png;base64,dGVzdA=="}}`
	updateResponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(updateResponse, httptest.NewRequest(http.MethodPut, "/api/people/"+person.ID, strings.NewReader(updateBody)))
	if updateResponse.Code != http.StatusOK {
		t.Fatalf("person update status = %d, body = %s", updateResponse.Code, updateResponse.Body.String())
	}
}
