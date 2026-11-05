package store

import (
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"fmt"
	"math/big"
	"sync"
	"time"

	"skirental/internal/domain"
)

//go:embed max-mustermann.jpg
var maxMustermannPhoto []byte

//go:embed maria-musterfrau.jpg
var mariaMusterfrauPhoto []byte

type Store interface {
	AddSki(brand, model string) (domain.Ski, error)
	AddSkiWithSize(brand, model string, size int) (domain.Ski, error)
	AddSkiWithSizeAndPrice(brand, model string, size, dailyPriceCents int) (domain.Ski, error)
	RegisterPerson(firstName, lastName string, birthDate time.Time, identification, country, city, postalCode, street, houseNumber, mobileNumber string, photo domain.IDPhoto) (domain.Person, error)
	UpdatePerson(id, firstName, lastName string, birthDate time.Time, identification, country, city, postalCode, street, houseNumber, mobileNumber string, photo domain.IDPhoto) (domain.Person, error)
	SeedDefaultPerson() error
	ListSkis() []domain.Ski
	ListPeople() []domain.Person
	ListLoans() []domain.Loan
	ListReservations() []domain.Reservation
	StartLoan(skiID, personID string) (domain.Loan, error)
	StartLoanForPeriod(skiID, personID string, startDate, endDate time.Time) (domain.Loan, error)
	EndLoan(skiID string) error
	CreateReservation(skiID, personID string) (domain.Reservation, error)
	CreateReservationForPeriod(skiID, personID string, startDate, endDate time.Time) (domain.Reservation, error)
	EditReservation(id, newSkiID string) (domain.Reservation, error)
	EditReservationForPeriod(id, newSkiID string, startDate, endDate time.Time) (domain.Reservation, error)
	ConvertReservation(id string) (domain.Loan, error)
	DiscardReservation(id string) error
}

type InMemoryStore struct {
	mu      sync.RWMutex
	service *domain.RentalService
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{service: domain.NewRentalService()}
}

func (store *InMemoryStore) AddSki(brand, model string) (domain.Ski, error) {
	return store.AddSkiWithSize(brand, model, 0)
}

func (store *InMemoryStore) AddSkiWithSize(brand, model string, size int) (domain.Ski, error) {
	return store.AddSkiWithSizeAndPrice(brand, model, size, 1000)
}

func (store *InMemoryStore) AddSkiWithSizeAndPrice(brand, model string, size, dailyPriceCents int) (domain.Ski, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	ski, err := store.service.AddSkiWithSizeAndPrice(brand, model, size, dailyPriceCents)
	if err != nil {
		return domain.Ski{}, fmt.Errorf("add ski: %w", err)
	}
	return ski, nil
}

func (store *InMemoryStore) SeedDefaultSkis() error {
	catalog := []struct {
		brand string
		model string
	}{
		{"Atomic", "Redster S9"},
		{"Rossignol", "Hero Elite ST Ti"},
		{"Head", "Supershape e-Magnum"},
		{"Salomon", "S/Max 10"},
		{"Fischer", "RC4 The Curv"},
		{"Völkl", "Racetiger SL"},
		{"Blizzard", "Firebird SRC"},
		{"Nordica", "Dobermann Spitfire 76 RB"},
		{"Elan", "Amphibio 14 Ti"},
		{"K2", "Disruption 76"},
	}
	sizes := []int{150, 155, 160, 165, 170, 175, 180, 185, 190, 195}
	for modelIndex, item := range catalog {
		for _, size := range sizes {
			for copy := 0; copy < 2; copy++ {
				if _, err := store.AddSkiWithSizeAndPrice(item.brand, item.model, size, 1000+modelIndex*100); err != nil {
					return fmt.Errorf("seed default ski catalog: %w", err)
				}
			}
		}
	}
	return nil
}

func (store *InMemoryStore) RegisterPerson(firstName, lastName string, birthDate time.Time, identification, country, city, postalCode, street, houseNumber, mobileNumber string, photo domain.IDPhoto) (domain.Person, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	person, err := store.service.RegisterPerson(firstName, lastName, birthDate, identification, country, city, postalCode, street, houseNumber, mobileNumber, photo)
	if err != nil {
		return domain.Person{}, fmt.Errorf("register person: %w", err)
	}
	return person, nil
}

func (store *InMemoryStore) UpdatePerson(id, firstName, lastName string, birthDate time.Time, identification, country, city, postalCode, street, houseNumber, mobileNumber string, photo domain.IDPhoto) (domain.Person, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	person, err := store.service.UpdatePerson(id, firstName, lastName, birthDate, identification, country, city, postalCode, street, houseNumber, mobileNumber, photo)
	if err != nil {
		return domain.Person{}, fmt.Errorf("update person: %w", err)
	}
	return person, nil
}

func (store *InMemoryStore) SeedDefaultPerson() error {
	birthDate, err := time.Parse("02.01.2006", "01.01.2000")
	if err != nil {
		return fmt.Errorf("parse default person birth date: %w", err)
	}
	country, postalCode, street, houseNumber, mobileNumber, err := randomDefaultPersonDetails()
	if err != nil {
		return fmt.Errorf("generate default person details: %w", err)
	}
	photo := domain.IDPhoto{
		ContentType: "image/jpeg",
		DataURL:     "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(maxMustermannPhoto),
	}
	if _, err := store.RegisterPerson("Max", "Mustermann", birthDate, "ID-MAX-MUSTERMANN", country, "München", postalCode, street, houseNumber, mobileNumber, photo); err != nil {
		return fmt.Errorf("seed default person: %w", err)
	}
	mariaPhoto := domain.IDPhoto{
		ContentType: "image/jpeg",
		DataURL:     "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(mariaMusterfrauPhoto),
	}
	if _, err := store.RegisterPerson("Maria", "Musterfrau", time.Date(1981, time.January, 1, 0, 0, 0, 0, time.UTC), "ID-MARIA-MUSTERFRAU", "Österreich", "Wien", "1010", "Berggasse", "12", "+436641234567", mariaPhoto); err != nil {
		return fmt.Errorf("seed second default person: %w", err)
	}
	return nil
}

func randomDefaultPersonDetails() (string, string, string, string, string, error) {
	streets := []string{"Alpenstraße", "Bergweg", "Gipfelstraße", "Talweg", "Schneestraße"}
	streetIndex, err := randomInt(len(streets))
	if err != nil {
		return "", "", "", "", "", err
	}
	houseNumber, err := randomInt(99)
	if err != nil {
		return "", "", "", "", "", err
	}
	postalCode, err := randomInt(90000)
	if err != nil {
		return "", "", "", "", "", err
	}
	phone, err := randomInt(10000000)
	if err != nil {
		return "", "", "", "", "", err
	}
	return "Deutschland", fmt.Sprintf("%05d", 10000+postalCode%90000), streets[streetIndex], fmt.Sprintf("%d", houseNumber+1), fmt.Sprintf("+4915%08d", phone), nil
}

func randomInt(max int) (int, error) {
	value, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(value.Int64()), nil
}

func (store *InMemoryStore) ListSkis() []domain.Ski {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.service.Skis()
}

func (store *InMemoryStore) ListPeople() []domain.Person {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.service.People()
}

func (store *InMemoryStore) ListLoans() []domain.Loan {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.service.Loans()
}

func (store *InMemoryStore) ListReservations() []domain.Reservation {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.service.Reservations()
}

func (store *InMemoryStore) StartLoan(skiID, personID string) (domain.Loan, error) {
	today := time.Now().UTC()
	return store.StartLoanForPeriod(skiID, personID, today, today)
}

func (store *InMemoryStore) StartLoanForPeriod(skiID, personID string, startDate, endDate time.Time) (domain.Loan, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	loan, err := store.service.StartLoanForPeriod(skiID, personID, startDate, endDate)
	if err != nil {
		return domain.Loan{}, fmt.Errorf("start loan: %w", err)
	}
	return loan, nil
}

func (store *InMemoryStore) EndLoan(skiID string) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if err := store.service.EndLoan(skiID); err != nil {
		return fmt.Errorf("end loan: %w", err)
	}
	return nil
}

func (store *InMemoryStore) CreateReservation(skiID, personID string) (domain.Reservation, error) {
	today := time.Now().UTC()
	return store.CreateReservationForPeriod(skiID, personID, today, today)
}

func (store *InMemoryStore) CreateReservationForPeriod(skiID, personID string, startDate, endDate time.Time) (domain.Reservation, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	reservation, err := store.service.CreateReservationForPeriod(skiID, personID, startDate, endDate)
	if err != nil {
		return domain.Reservation{}, fmt.Errorf("create reservation: %w", err)
	}
	return reservation, nil
}

func (store *InMemoryStore) EditReservation(id, newSkiID string) (domain.Reservation, error) {
	store.mu.RLock()
	reservations := store.service.Reservations()
	store.mu.RUnlock()
	for _, reservation := range reservations {
		if reservation.ID == id {
			return store.EditReservationForPeriod(id, newSkiID, reservation.StartedAt, reservation.EndDate)
		}
	}
	return store.EditReservationForPeriod(id, newSkiID, time.Now().UTC(), time.Now().UTC())
}

func (store *InMemoryStore) EditReservationForPeriod(id, newSkiID string, startDate, endDate time.Time) (domain.Reservation, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	reservation, err := store.service.EditReservationForPeriod(id, newSkiID, startDate, endDate)
	if err != nil {
		return domain.Reservation{}, fmt.Errorf("edit reservation: %w", err)
	}
	return reservation, nil
}

func (store *InMemoryStore) ConvertReservation(id string) (domain.Loan, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	loan, err := store.service.ConvertReservation(id)
	if err != nil {
		return domain.Loan{}, fmt.Errorf("convert reservation: %w", err)
	}
	return loan, nil
}

func (store *InMemoryStore) DiscardReservation(id string) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if err := store.service.DiscardReservation(id); err != nil {
		return fmt.Errorf("discard reservation: %w", err)
	}
	return nil
}
