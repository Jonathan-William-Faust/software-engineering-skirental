package domain

import (
	"fmt"
	"time"
)

type RentalService struct {
	ids             IDGenerator
	skis            map[string]*Ski
	people          map[string]*Person
	loans           map[string]Loan
	reservations    map[string]Reservation
	nextReservation uint64
}

func NewRentalService() *RentalService {
	return &RentalService{
		skis: make(map[string]*Ski), people: make(map[string]*Person),
		loans: make(map[string]Loan), reservations: make(map[string]Reservation),
	}
}

func (service *RentalService) AddSki(brand, model string) (Ski, error) {
	return service.AddSkiWithSizeAndPrice(brand, model, 0, 1000)
}

func (service *RentalService) AddSkiWithSize(brand, model string, size int) (Ski, error) {
	return service.AddSkiWithSizeAndPrice(brand, model, size, 1000)
}

func (service *RentalService) AddSkiWithSizeAndPrice(brand, model string, size, dailyPriceCents int) (Ski, error) {
	id, err := service.ids.NewSkiID()
	if err != nil {
		return Ski{}, fmt.Errorf("generate ski ID: %w", err)
	}
	ski, err := NewSkiWithSizeAndPrice(id, brand, model, size, dailyPriceCents)
	if err != nil {
		return Ski{}, fmt.Errorf("create ski: %w", err)
	}
	service.skis[id] = &ski
	return ski, nil
}

func (service *RentalService) RegisterPerson(firstName, lastName string, birthDate time.Time, identification, country, city, postalCode, street, houseNumber, mobileNumber string, photo IDPhoto) (Person, error) {
	id, err := service.ids.NewPersonID()
	if err != nil {
		return Person{}, fmt.Errorf("generate person ID: %w", err)
	}
	person, err := NewPerson(id, firstName, lastName, identification, birthDate, country, city, postalCode, street, houseNumber, mobileNumber, photo)
	if err != nil {
		return Person{}, fmt.Errorf("register person: %w", err)
	}
	service.people[id] = &person
	return person, nil
}

func (service *RentalService) UpdatePerson(id, firstName, lastName string, birthDate time.Time, identification, country, city, postalCode, street, houseNumber, mobileNumber string, photo IDPhoto) (Person, error) {
	if _, ok := service.people[id]; !ok {
		return Person{}, fmt.Errorf("%w: %s", ErrPersonNotFound, id)
	}
	person, err := NewPerson(id, firstName, lastName, identification, birthDate, country, city, postalCode, street, houseNumber, mobileNumber, photo)
	if err != nil {
		return Person{}, fmt.Errorf("update person: %w", err)
	}
	service.people[id] = &person
	return person, nil
}

func (service *RentalService) StartLoan(skiID, personID string) (Loan, error) {
	today := dateOnly(time.Now().UTC())
	return service.StartLoanForPeriod(skiID, personID, today, today)
}

func (service *RentalService) StartLoanForPeriod(skiID, personID string, startDate, endDate time.Time) (Loan, error) {
	ski, person, err := service.objects(skiID, personID)
	if err != nil {
		return Loan{}, err
	}
	if ski.Status != SkiAvailable {
		return Loan{}, fmt.Errorf("%w: %s", ErrSkiUnavailable, skiID)
	}
	startDate = dateOnly(startDate)
	endDate = dateOnly(endDate)
	if !startDate.Equal(dateOnly(time.Now().UTC())) {
		return Loan{}, fmt.Errorf("%w: loans must start today", ErrInvalidRentalPeriod)
	}
	days, err := rentalDays(startDate, endDate)
	if err != nil {
		return Loan{}, err
	}
	if existing, ok := service.loans[skiID]; ok && existing.Active() {
		return Loan{}, fmt.Errorf("%w: %s", ErrActiveLoanExists, skiID)
	}
	if err := ski.Rent(); err != nil {
		return Loan{}, fmt.Errorf("rent ski: %w", err)
	}
	loan := Loan{SkiID: ski.ID, PersonID: person.ID, StartedAt: startDate, EndDate: endDate, Days: days, TotalPriceCents: days * ski.DailyPriceCents}
	service.loans[ski.ID] = loan
	return loan, nil
}

func (service *RentalService) AvailableSkisForPeriod(startDate, endDate time.Time) ([]Ski, error) {
	startDate = dateOnly(startDate)
	endDate = dateOnly(endDate)
	if _, err := rentalDays(startDate, endDate); err != nil {
		return nil, err
	}
	available := make([]Ski, 0, len(service.skis))
	for _, ski := range service.skis {
		if ski.Status != SkiAvailable || service.skiUnavailableForPeriod(ski.ID, startDate, endDate) {
			continue
		}
		available = append(available, *ski)
	}
	return available, nil
}

func (service *RentalService) skiUnavailableForPeriod(skiID string, startDate, endDate time.Time) bool {
	if loan, ok := service.loans[skiID]; ok && loan.Active() {
		return true
	}
	for _, reservation := range service.reservations {
		if reservation.Active() && reservation.SkiID == skiID &&
			periodsOverlap(startDate, endDate, dateOnly(reservation.StartedAt), dateOnly(reservation.EndDate)) {
			return true
		}
	}
	return false
}

func periodsOverlap(firstStart, firstEnd, secondStart, secondEnd time.Time) bool {
	return !firstEnd.Before(secondStart) && !secondEnd.Before(firstStart)
}

func (service *RentalService) EndLoan(skiID string) error {
	ski, ok := service.skis[skiID]
	if !ok {
		return fmt.Errorf("%w: %s", ErrSkiNotFound, skiID)
	}
	loan, ok := service.loans[skiID]
	if !ok || !loan.Active() {
		return fmt.Errorf("%w: %s", ErrLoanNotFound, skiID)
	}
	if err := ski.Return(); err != nil {
		return fmt.Errorf("return ski: %w", err)
	}
	now := time.Now()
	loan.EndedAt = &now
	service.loans[skiID] = loan
	return nil
}

func (service *RentalService) CreateReservation(skiID, personID string) (Reservation, error) {
	today := dateOnly(time.Now().UTC())
	return service.CreateReservationForPeriod(skiID, personID, today, today)
}

func (service *RentalService) CreateReservationForPeriod(skiID, personID string, startDate, endDate time.Time) (Reservation, error) {
	ski, person, err := service.objects(skiID, personID)
	if err != nil {
		return Reservation{}, err
	}
	if ski.Status != SkiAvailable {
		return Reservation{}, fmt.Errorf("%w: %s", ErrSkiUnavailable, skiID)
	}
	startDate = dateOnly(startDate)
	endDate = dateOnly(endDate)
	if startDate.Before(dateOnly(time.Now().UTC())) {
		return Reservation{}, fmt.Errorf("%w: reservation cannot start in the past", ErrInvalidRentalPeriod)
	}
	if _, err := rentalDays(startDate, endDate); err != nil {
		return Reservation{}, err
	}
	for _, reservation := range service.reservations {
		if reservation.Active() && reservation.SkiID == ski.ID {
			return Reservation{}, fmt.Errorf("%w: ski %s", ErrDuplicateReservation, skiID)
		}
	}
	if err := ski.Reserve(); err != nil {
		return Reservation{}, fmt.Errorf("reserve ski: %w", err)
	}
	service.nextReservation++
	reservation := Reservation{
		ID: fmt.Sprintf("R-%04d", service.nextReservation), SkiID: ski.ID,
		PersonID: person.ID, StartedAt: startDate, EndDate: endDate,
		DailyPriceCents: ski.DailyPriceCents, Status: ReservationActive,
	}
	service.reservations[reservation.ID] = reservation
	return reservation, nil
}

func (service *RentalService) EditReservation(id, newSkiID string) (Reservation, error) {
	reservation, ok := service.reservations[id]
	if !ok {
		return Reservation{}, fmt.Errorf("%w: %s", ErrReservationNotFound, id)
	}
	return service.EditReservationForPeriod(id, newSkiID, reservation.StartedAt, reservation.EndDate)
}

func (service *RentalService) EditReservationForPeriod(id, newSkiID string, startDate, endDate time.Time) (Reservation, error) {
	reservation, ok := service.reservations[id]
	if !ok {
		return Reservation{}, fmt.Errorf("%w: %s", ErrReservationNotFound, id)
	}
	if !reservation.Active() {
		return Reservation{}, fmt.Errorf("%w: %s", ErrReservationNotActive, id)
	}
	startDate = dateOnly(startDate)
	endDate = dateOnly(endDate)
	if startDate.Before(dateOnly(time.Now().UTC())) {
		return Reservation{}, fmt.Errorf("%w: reservation cannot start in the past", ErrInvalidRentalPeriod)
	}
	if _, err := rentalDays(startDate, endDate); err != nil {
		return Reservation{}, err
	}
	oldSki := service.skis[reservation.SkiID]
	newSki, ok := service.skis[newSkiID]
	if !ok {
		return Reservation{}, fmt.Errorf("%w: %s", ErrSkiNotFound, newSkiID)
	}
	if newSkiID != reservation.SkiID && newSki.Status != SkiAvailable {
		return Reservation{}, fmt.Errorf("%w: %s", ErrSkiUnavailable, newSkiID)
	}
	if newSkiID == reservation.SkiID {
		reservation.StartedAt = startDate
		reservation.EndDate = endDate
		service.reservations[id] = reservation
		return reservation, nil
	}
	if err := oldSki.Return(); err != nil {
		return Reservation{}, fmt.Errorf("release old reservation ski: %w", err)
	}
	if err := newSki.Reserve(); err != nil {
		return Reservation{}, fmt.Errorf("reserve replacement ski: %w", err)
	}
	reservation.SkiID = newSkiID
	reservation.StartedAt = startDate
	reservation.EndDate = endDate
	service.reservations[id] = reservation
	return reservation, nil
}

func (service *RentalService) ConvertReservation(id string) (Loan, error) {
	reservation, ok := service.reservations[id]
	if !ok {
		return Loan{}, fmt.Errorf("%w: %s", ErrReservationNotFound, id)
	}
	if !reservation.Active() {
		return Loan{}, fmt.Errorf("%w: %s", ErrReservationNotActive, id)
	}
	today := dateOnly(time.Now().UTC())
	if !dateOnly(reservation.StartedAt).Equal(today) {
		return Loan{}, fmt.Errorf("%w: conversion is only possible on the reservation start date", ErrReservationNotStarted)
	}
	ski := service.skis[reservation.SkiID]
	startDate := today
	endDate := dateOnly(reservation.EndDate)
	days, err := rentalDays(startDate, endDate)
	if err != nil {
		return Loan{}, fmt.Errorf("convert reservation: %w", err)
	}
	if err := ski.Rent(); err != nil {
		return Loan{}, fmt.Errorf("convert reservation: %w", err)
	}
	loan := Loan{SkiID: reservation.SkiID, PersonID: reservation.PersonID, StartedAt: startDate, EndDate: endDate, Days: days, TotalPriceCents: days * ski.DailyPriceCents}
	service.loans[reservation.SkiID] = loan
	now := time.Now()
	reservation.EndedAt = &now
	reservation.Status = ReservationConverted
	service.reservations[id] = reservation
	return loan, nil
}

func dateOnly(value time.Time) time.Time {
	value = value.UTC()
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)
}

func rentalDays(startDate, endDate time.Time) (int, error) {
	if endDate.Before(startDate) {
		return 0, fmt.Errorf("%w: end date cannot be before start date", ErrInvalidRentalPeriod)
	}
	return int(endDate.Sub(startDate).Hours()/24) + 1, nil
}

func (service *RentalService) DiscardReservation(id string) error {
	reservation, ok := service.reservations[id]
	if !ok {
		return fmt.Errorf("%w: %s", ErrReservationNotFound, id)
	}
	if !reservation.Active() {
		return fmt.Errorf("%w: %s", ErrReservationNotActive, id)
	}
	if err := service.skis[reservation.SkiID].Return(); err != nil {
		return fmt.Errorf("release reservation: %w", err)
	}
	now := time.Now()
	reservation.EndedAt = &now
	reservation.Status = ReservationDiscarded
	service.reservations[id] = reservation
	return nil
}

func (service *RentalService) Skis() []Ski {
	result := make([]Ski, 0, len(service.skis))
	for _, ski := range service.skis {
		result = append(result, *ski)
	}
	return result
}

func (service *RentalService) People() []Person {
	result := make([]Person, 0, len(service.people))
	for _, person := range service.people {
		result = append(result, *person)
	}
	return result
}

func (service *RentalService) Loans() []Loan {
	result := make([]Loan, 0, len(service.loans))
	for _, loan := range service.loans {
		result = append(result, loan)
	}
	return result
}

func (service *RentalService) Reservations() []Reservation {
	result := make([]Reservation, 0, len(service.reservations))
	for _, reservation := range service.reservations {
		result = append(result, reservation)
	}
	return result
}

func (service *RentalService) objects(skiID, personID string) (*Ski, *Person, error) {
	ski, ok := service.skis[skiID]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %s", ErrSkiNotFound, skiID)
	}
	person, ok := service.people[personID]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %s", ErrPersonNotFound, personID)
	}
	return ski, person, nil
}
