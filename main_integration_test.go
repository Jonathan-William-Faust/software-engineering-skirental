package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerServesFrontendHealthAndRentalFlow(t *testing.T) {
	server := newServer()
	request(t, server.Handler, http.MethodGet, "/", "", http.StatusOK)
	request(t, server.Handler, http.MethodGet, "/healthz", "", http.StatusOK)

	ski := postJSON(t, server.Handler, "/api/skis", `{"brand":"Atomic","model":"Redster"}`, http.StatusCreated)
	person := postJSON(t, server.Handler, "/api/people", `{"first_name":"Ada","last_name":"Lovelace","birth_date":"1990-01-01T00:00:00Z","identification":"ID-1","photo":{"content_type":"image/png"}}`, http.StatusCreated)
	skiID := stringField(t, ski, "ID")
	personID := stringField(t, person, "ID")

	postJSON(t, server.Handler, "/api/loans", `{"ski_id":"`+skiID+`","person_id":"`+personID+`"}`, http.StatusCreated)
	assertDashboard(t, server.Handler, 0, 1, 1, 1)
	request(t, server.Handler, http.MethodDelete, "/api/loans/"+skiID, "", http.StatusNoContent)
	assertDashboard(t, server.Handler, 1, 0, 1, 1)
}

func TestServerReservationLifecycle(t *testing.T) {
	server := newServer()
	firstSki := postJSON(t, server.Handler, "/api/skis", `{"brand":"Atomic","model":"Redster"}`, http.StatusCreated)
	secondSki := postJSON(t, server.Handler, "/api/skis", `{"brand":"Head","model":"Supershape"}`, http.StatusCreated)
	person := postJSON(t, server.Handler, "/api/people", `{"first_name":"Ada","last_name":"Lovelace","birth_date":"1990-01-01T00:00:00Z","identification":"ID-1","photo":{"content_type":"image/png"}}`, http.StatusCreated)
	reservation := postJSON(t, server.Handler, "/api/reservations", `{"ski_id":"`+stringField(t, firstSki, "ID")+`","person_id":"`+stringField(t, person, "ID")+`"}`, http.StatusCreated)
	reservationID := stringField(t, reservation, "ID")
	requestBody := `{"ski_id":"` + stringField(t, secondSki, "ID") + `"}`
	putJSON(t, server.Handler, "/api/reservations/"+reservationID, requestBody, http.StatusOK)
	request(t, server.Handler, http.MethodPost, "/api/reservations/"+reservationID+"/convert", "", http.StatusCreated)
	assertDashboard(t, server.Handler, 1, 1, 2, 1)

	thirdSki := postJSON(t, server.Handler, "/api/skis", `{"brand":"K2","model":"Mindbender"}`, http.StatusCreated)
	secondReservation := postJSON(t, server.Handler, "/api/reservations", `{"ski_id":"`+stringField(t, thirdSki, "ID")+`","person_id":"`+stringField(t, person, "ID")+`"}`, http.StatusCreated)
	request(t, server.Handler, http.MethodDelete, "/api/reservations/"+stringField(t, secondReservation, "ID"), "", http.StatusNoContent)
	assertDashboard(t, server.Handler, 2, 1, 3, 1)
}

func assertDashboard(t *testing.T, handler http.Handler, available, rented, total, people int) {
	t.Helper()
	response := request(t, handler, http.MethodGet, "/api/dashboard", "", http.StatusOK)
	var dashboard map[string]int
	decodeResponse(t, response, &dashboard)
	if dashboard["available_skis"] != available || dashboard["rented_skis"] != rented ||
		dashboard["total_skis"] != total || dashboard["people"] != people {
		t.Fatalf("dashboard = %#v, want available=%d rented=%d total=%d people=%d", dashboard, available, rented, total, people)
	}
}

func postJSON(t *testing.T, handler http.Handler, path, body string, status int) map[string]any {
	t.Helper()
	response := request(t, handler, http.MethodPost, path, body, status)
	var result map[string]any
	decodeResponse(t, response, &result)
	return result
}

func putJSON(t *testing.T, handler http.Handler, path, body string, status int) map[string]any {
	t.Helper()
	response := request(t, handler, http.MethodPut, path, body, status)
	var result map[string]any
	decodeResponse(t, response, &result)
	return result
}

func request(t *testing.T, handler http.Handler, method, path, body string, status int) *httptest.ResponseRecorder {
	t.Helper()
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != status {
		t.Fatalf("%s %s returned %d, want %d: %s", method, path, response.Code, status, response.Body.String())
	}
	return response
}

func decodeResponse(t *testing.T, response *httptest.ResponseRecorder, target any) {
	t.Helper()
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		t.Fatal(err)
	}
}

func stringField(t *testing.T, object map[string]any, name string) string {
	t.Helper()
	value, ok := object[name].(string)
	if !ok {
		t.Fatalf("response field %s missing in %#v", name, object)
	}
	return value
}
