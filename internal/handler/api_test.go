package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"skirental/internal/store"
)

func TestAPIRegistrationLoanAndDashboard(t *testing.T) {
	api := NewAPI(store.NewInMemoryStore()).Handler()
	ski := request(t, api, http.MethodPost, "/api/skis", `{"brand":"Atomic","model":"Redster"}`, http.StatusCreated)
	var skiBody struct {
		ID string `json:"ID"`
	}
	decode(t, ski, &skiBody)
	person := request(t, api, http.MethodPost, "/api/people", `{"first_name":"Ada","last_name":"Lovelace","birth_date":"1990-01-01T00:00:00Z","identification":"ID-1","photo":{"content_type":"image/png"}}`, http.StatusCreated)
	var personBody struct {
		ID string `json:"ID"`
	}
	decode(t, person, &personBody)
	request(t, api, http.MethodPost, "/api/loans", `{"ski_id":"`+skiBody.ID+`","person_id":"`+personBody.ID+`"}`, http.StatusCreated)
	dashboard := request(t, api, http.MethodGet, "/api/dashboard", "", http.StatusOK)
	var counts map[string]int
	decode(t, dashboard, &counts)
	if counts["rented_skis"] != 1 || counts["available_skis"] != 0 || counts["people"] != 1 {
		t.Fatalf("unexpected dashboard: %#v", counts)
	}
}

func TestAPIRejectsInvalidPhotoAndUnknownSki(t *testing.T) {
	api := NewAPI(store.NewInMemoryStore()).Handler()
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Ada","last_name":"Lovelace","birth_date":"1990-01-01T00:00:00Z","identification":"ID-1","photo":{"content_type":"image/jpeg"}}`, http.StatusBadRequest)
	request(t, api, http.MethodPost, "/api/loans", `{"ski_id":"S-9999","person_id":"P-0001"}`, http.StatusNotFound)
	request(t, api, http.MethodPost, "/api/skis", `{"brand":"Atomic"}`, http.StatusBadRequest)
}

func TestAPIAcceptsGermanBirthDateFormat(t *testing.T) {
	api := NewAPI(store.NewInMemoryStore()).Handler()
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Ada","last_name":"Lovelace","birth_date":"05.11.1990","identification":"ID-1","photo":{"content_type":"image/png"}}`, http.StatusCreated)
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Ada","last_name":"Lovelace","birth_date":"31.02.1990","identification":"ID-2","photo":{"content_type":"image/png"}}`, http.StatusBadRequest)
}

func request(t *testing.T, handler http.Handler, method, path, body string, status int) *httptest.ResponseRecorder {
	t.Helper()
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != status {
		t.Fatalf("%s %s: status = %d, want %d; body=%s", method, path, response.Code, status, response.Body.String())
	}
	return response
}

func decode(t *testing.T, response *httptest.ResponseRecorder, target any) {
	t.Helper()
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		t.Fatal(err)
	}
}
