package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"skirental/internal/domain"
	"skirental/internal/store"
)

type API struct {
	store store.Store
}

func NewAPI(rentalStore store.Store) *API {
	return &API{store: rentalStore}
}

func (api *API) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", api.handleAPI)
	return mux
}

func (api *API) handleAPI(response http.ResponseWriter, request *http.Request) {
	path := strings.TrimPrefix(request.URL.Path, "/api/")
	switch {
	case path == "dashboard":
		api.dashboard(response, request)
	case path == "skis":
		api.skis(response, request)
	case path == "people":
		api.people(response, request)
	case strings.HasPrefix(path, "people/"):
		api.personByID(response, request, strings.TrimPrefix(path, "people/"))
	case path == "loans":
		api.loans(response, request)
	case strings.HasPrefix(path, "loans/"):
		api.loanBySki(response, request, strings.TrimPrefix(path, "loans/"))
	case path == "reservations":
		api.reservations(response, request)
	case strings.HasPrefix(path, "reservations/"):
		api.reservationByID(response, request, strings.TrimPrefix(path, "reservations/"))
	default:
		writeError(response, http.StatusNotFound, "resource not found")
	}
}

type skiRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Size  int    `json:"size"`
}

type personRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	Identification string `json:"identification"`
	Country        string `json:"country"`
	City           string `json:"city"`
	PostalCode     string `json:"postal_code"`
	Street         string `json:"street"`
	HouseNumber    string `json:"house_number"`
	MobileNumber   string `json:"mobile_number"`
	Photo          struct {
		ContentType string `json:"content_type"`
		DataURL     string `json:"data_url"`
	} `json:"photo"`
}

type loanRequest struct {
	SkiID    string `json:"ski_id"`
	PersonID string `json:"person_id"`
}

type reservationRequest struct {
	SkiID    string `json:"ski_id"`
	PersonID string `json:"person_id"`
}

type reservationEditRequest struct {
	SkiID string `json:"ski_id"`
}

func (api *API) dashboard(response http.ResponseWriter, request *http.Request) {
	if !requireMethod(response, request, http.MethodGet) {
		return
	}
	var available, rented int
	for _, ski := range api.store.ListSkis() {
		switch ski.Status {
		case domain.SkiAvailable:
			available++
		case domain.SkiRented:
			rented++
		}
	}
	writeJSON(response, http.StatusOK, map[string]int{
		"available_skis": available,
		"rented_skis":    rented,
		"total_skis":     len(api.store.ListSkis()),
		"people":         len(api.store.ListPeople()),
	})
}

func (api *API) skis(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		writeJSON(response, http.StatusOK, api.store.ListSkis())
	case http.MethodPost:
		var input skiRequest
		if !decodeJSON(response, request, &input) {
			return
		}
		ski, err := api.store.AddSkiWithSize(input.Brand, input.Model, input.Size)
		if err != nil {
			writeDomainError(response, err)
			return
		}
		writeJSON(response, http.StatusCreated, ski)
	default:
		writeMethodNotAllowed(response, http.MethodGet, http.MethodPost)
	}
}

func (api *API) people(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		writeJSON(response, http.StatusOK, api.store.ListPeople())
	case http.MethodPost:
		var input personRequest
		if !decodeJSON(response, request, &input) {
			return
		}

		birthDate, err := parseBirthDate(input.BirthDate)
		if err != nil {
			writeError(response, http.StatusBadRequest, err.Error())
			return
		}
		person, err := api.store.RegisterPerson(input.FirstName, input.LastName, birthDate, input.Identification, input.Country, input.City, input.PostalCode, input.Street, input.HouseNumber, input.MobileNumber, domain.IDPhoto{ContentType: input.Photo.ContentType, DataURL: input.Photo.DataURL})
		if err != nil {
			writeDomainError(response, err)
			return
		}

		writeJSON(response, http.StatusCreated, person)
	default:
		writeMethodNotAllowed(response, http.MethodGet, http.MethodPost)
	}
}

func (api *API) personByID(response http.ResponseWriter, request *http.Request, id string) {
	if request.Method != http.MethodPut {
		writeMethodNotAllowed(response, http.MethodPut)
		return
	}
	var input personRequest
	if !decodeJSON(response, request, &input) {
		return
	}
	birthDate, err := parseBirthDate(input.BirthDate)
	if err != nil {
		writeError(response, http.StatusBadRequest, err.Error())
		return
	}
	person, err := api.store.UpdatePerson(id, input.FirstName, input.LastName, birthDate, input.Identification, input.Country, input.City, input.PostalCode, input.Street, input.HouseNumber, input.MobileNumber, domain.IDPhoto{ContentType: input.Photo.ContentType, DataURL: input.Photo.DataURL})
	if err != nil {
		writeDomainError(response, err)
		return
	}
	writeJSON(response, http.StatusOK, person)
}

func parseBirthDate(value string) (time.Time, error) {
	if birthDate, err := time.Parse("02.01.2006", value); err == nil {
		return validateBirthDate(birthDate)
	}
	if birthDate, err := time.Parse(time.RFC3339, value); err == nil {
		return validateBirthDate(birthDate)
	}
	return time.Time{}, fmt.Errorf("birth_date must use format DD.MM.YYYY")
}

func validateBirthDate(birthDate time.Time) (time.Time, error) {
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if birthDate.After(today) {
		return time.Time{}, fmt.Errorf("birth_date cannot be in the future")
	}
	if birthDate.AddDate(18, 0, 0).After(today) {
		return time.Time{}, fmt.Errorf("person must be at least 18 years old")
	}
	if birthDate.AddDate(100, 0, 0).Before(today) {
		return time.Time{}, fmt.Errorf("person must not be older than 100 years")
	}
	return birthDate, nil
}

func (api *API) loans(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		writeJSON(response, http.StatusOK, api.store.ListLoans())
	case http.MethodPost:
		var input loanRequest
		if !decodeJSON(response, request, &input) {
			return
		}
		loan, err := api.store.StartLoan(input.SkiID, input.PersonID)
		if err != nil {
			writeDomainError(response, err)
			return
		}
		writeJSON(response, http.StatusCreated, loan)
	default:
		writeMethodNotAllowed(response, http.MethodGet, http.MethodPost)
	}
}

func (api *API) loanBySki(response http.ResponseWriter, request *http.Request, skiID string) {
	if request.Method != http.MethodDelete {
		writeMethodNotAllowed(response, http.MethodDelete)
		return
	}
	if err := api.store.EndLoan(skiID); err != nil {
		writeDomainError(response, err)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}

func (api *API) reservations(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		writeJSON(response, http.StatusOK, api.store.ListReservations())
	case http.MethodPost:
		var input reservationRequest
		if !decodeJSON(response, request, &input) {
			return
		}
		reservation, err := api.store.CreateReservation(input.SkiID, input.PersonID)
		if err != nil {
			writeDomainError(response, err)
			return
		}
		writeJSON(response, http.StatusCreated, reservation)
	default:
		writeMethodNotAllowed(response, http.MethodGet, http.MethodPost)
	}
}

func (api *API) reservationByID(response http.ResponseWriter, request *http.Request, path string) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 2 && parts[1] == "convert" {
		if request.Method != http.MethodPost {
			writeMethodNotAllowed(response, http.MethodPost)
			return
		}
		loan, err := api.store.ConvertReservation(parts[0])
		if err != nil {
			writeDomainError(response, err)
			return
		}
		writeJSON(response, http.StatusCreated, loan)
		return
	}
	if len(parts) != 1 || parts[0] == "" {
		writeError(response, http.StatusNotFound, "reservation not found")
		return
	}
	switch request.Method {
	case http.MethodPut:
		var input reservationEditRequest
		if !decodeJSON(response, request, &input) {
			return
		}
		reservation, err := api.store.EditReservation(parts[0], input.SkiID)
		if err != nil {
			writeDomainError(response, err)
			return
		}
		writeJSON(response, http.StatusOK, reservation)
	case http.MethodDelete:
		if err := api.store.DiscardReservation(parts[0]); err != nil {
			writeDomainError(response, err)
			return
		}
		response.WriteHeader(http.StatusNoContent)
	default:
		writeMethodNotAllowed(response, http.MethodPut, http.MethodDelete)
	}
}

func requireMethod(response http.ResponseWriter, request *http.Request, method string) bool {
	if request.Method == method {
		return true
	}
	writeMethodNotAllowed(response, method)
	return false
}

func decodeJSON(response http.ResponseWriter, request *http.Request, target any) bool {
	if request.Body == nil {
		writeError(response, http.StatusBadRequest, "request body is required")
		return false
	}
	decoder := json.NewDecoder(http.MaxBytesReader(response, request.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		writeError(response, http.StatusBadRequest, fmt.Sprintf("invalid JSON: %v", err))
		return false
	}
	return true
}

func writeDomainError(response http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	switch {
	case errors.Is(err, domain.ErrSkiNotFound), errors.Is(err, domain.ErrPersonNotFound),
		errors.Is(err, domain.ErrReservationNotFound), errors.Is(err, domain.ErrLoanNotFound):
		status = http.StatusNotFound
	}
	writeError(response, status, err.Error())
}

func writeMethodNotAllowed(response http.ResponseWriter, methods ...string) {
	response.Header().Set("Allow", strings.Join(methods, ", "))
	writeError(response, http.StatusMethodNotAllowed, "method not allowed")
}

func writeError(response http.ResponseWriter, status int, message string) {
	writeJSON(response, status, map[string]string{"error": message})
}

func writeJSON(response http.ResponseWriter, status int, value any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_ = json.NewEncoder(response).Encode(value)
}
