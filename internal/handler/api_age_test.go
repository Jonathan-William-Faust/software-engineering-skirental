package handler

import (
	"net/http"
	"testing"

	"skirental/internal/store"
)

func TestAPIRejectsFutureAndUnderageBirthDates(t *testing.T) {
	api := NewAPI(store.NewInMemoryStore()).Handler()
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Future","last_name":"Person","birth_date":"31.12.2099","identification":"ID-2","photo":{"content_type":"image/png"}}`, http.StatusBadRequest)
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Young","last_name":"Person","birth_date":"06.11.2010","identification":"ID-3","photo":{"content_type":"image/png"}}`, http.StatusBadRequest)
}
