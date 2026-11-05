package handler

import (
	"net/http"
	"testing"

	"skirental/internal/store"
)

func TestAPIRejectsFutureAndUnderageBirthDates(t *testing.T) {
	api := NewAPI(store.NewInMemoryStore()).Handler()
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Future","last_name":"Person","birth_date":"31.12.2099","identification":"ID-2","country":"Deutschland","city":"Berlin","postal_code":"10115","street":"Musterstraße","house_number":"1","mobile_number":"+491234567890","photo":{"content_type":"image/png"}}`, http.StatusBadRequest)
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Young","last_name":"Person","birth_date":"06.11.2010","identification":"ID-3","country":"Deutschland","city":"Berlin","postal_code":"10115","street":"Musterstraße","house_number":"1","mobile_number":"+491234567890","photo":{"content_type":"image/png"}}`, http.StatusBadRequest)
	request(t, api, http.MethodPost, "/api/people", `{"first_name":"Old","last_name":"Person","birth_date":"01.01.1900","identification":"ID-4","country":"Deutschland","city":"Berlin","postal_code":"10115","street":"Musterstraße","house_number":"1","mobile_number":"+491234567890","photo":{"content_type":"image/png"}}`, http.StatusBadRequest)
}
