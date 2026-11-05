package store

import (
	"fmt"
	"sync"
	"time"

	"skirental/internal/domain"
)

type Store interface {
	AddSki(brand, model string) (domain.Ski, error)
	RegisterPerson(firstName, lastName string, birthDate time.Time, identification string, photo domain.IDPhoto) (domain.Person, error)
	ListSkis() []domain.Ski
	ListPeople() []domain.Person
	ListLoans() []domain.Loan
	ListReservations() []domain.Reservation
	StartLoan(skiID, personID string) (domain.Loan, error)
	EndLoan(skiID string) error
	CreateReservation(skiID, personID string) (domain.Reservation, error)
	EditReservation(id, newSkiID string) (domain.Reservation, error)
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
	store.mu.Lock()
	defer store.mu.Unlock()
	ski, err := store.service.AddSki(brand, model)
	if err != nil {
		return domain.Ski{}, fmt.Errorf("add ski: %w", err)
	}
	return ski, nil
}

func (store *InMemoryStore) RegisterPerson(firstName, lastName string, birthDate time.Time, identification string, photo domain.IDPhoto) (domain.Person, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	person, err := store.service.RegisterPerson(firstName, lastName, birthDate, identification, photo)
	if err != nil {
		return domain.Person{}, fmt.Errorf("register person: %w", err)
	}
	return person, nil
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
	store.mu.Lock()
	defer store.mu.Unlock()
	loan, err := store.service.StartLoan(skiID, personID)
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
	store.mu.Lock()
	defer store.mu.Unlock()
	reservation, err := store.service.CreateReservation(skiID, personID)
	if err != nil {
		return domain.Reservation{}, fmt.Errorf("create reservation: %w", err)
	}
	return reservation, nil
}

func (store *InMemoryStore) EditReservation(id, newSkiID string) (domain.Reservation, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	reservation, err := store.service.EditReservation(id, newSkiID)
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
