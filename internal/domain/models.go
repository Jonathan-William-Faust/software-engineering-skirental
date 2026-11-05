package domain

import (
	"fmt"
	"strings"
	"time"
)

type SkiStatus string

const (
	SkiAvailable SkiStatus = "verfügbar"
	SkiReserved  SkiStatus = "reserviert"
	SkiRented    SkiStatus = "verliehen"
)

func (status SkiStatus) valid() bool {
	return status == SkiAvailable || status == SkiReserved || status == SkiRented
}

type Ski struct {
	ID     string
	Brand  string
	Model  string
	Status SkiStatus
}

func NewSki(id, brand, model string) (Ski, error) {
	if err := ValidateSkiID(id); err != nil {
		return Ski{}, err
	}
	if strings.TrimSpace(brand) == "" || strings.TrimSpace(model) == "" {
		return Ski{}, fmt.Errorf("%w: ski brand and model are required", ErrInvalidSki)
	}
	return Ski{ID: id, Brand: brand, Model: model, Status: SkiAvailable}, nil
}

func (ski *Ski) Reserve() error {
	return ski.changeStatus(SkiReserved)
}

func (ski *Ski) Rent() error {
	return ski.changeStatus(SkiRented)
}

func (ski *Ski) Return() error {
	return ski.changeStatus(SkiAvailable)
}

func (ski *Ski) changeStatus(next SkiStatus) error {
	if !next.valid() {
		return fmt.Errorf("%w: %q", ErrInvalidSkiStatus, next)
	}
	allowed := (ski.Status == SkiAvailable && (next == SkiReserved || next == SkiRented)) ||
		(ski.Status == SkiReserved && (next == SkiAvailable || next == SkiRented)) ||
		(ski.Status == SkiRented && next == SkiAvailable)
	if !allowed {
		return fmt.Errorf("%w: %s -> %s", ErrInvalidStatusTransition, ski.Status, next)
	}
	ski.Status = next
	return nil
}

type IDPhoto struct {
	ContentType string
}

func (photo IDPhoto) IsPNG() bool {
	return strings.EqualFold(strings.TrimSpace(photo.ContentType), "image/png")
}

type Person struct {
	ID             string
	FirstName      string
	LastName       string
	BirthDate      time.Time
	Identification string
	IDPhoto        IDPhoto
}

func NewPerson(id, firstName, lastName, identification string, birthDate time.Time, photo IDPhoto) (Person, error) {
	if err := ValidatePersonID(id); err != nil {
		return Person{}, err
	}
	if strings.TrimSpace(firstName) == "" || strings.TrimSpace(lastName) == "" ||
		strings.TrimSpace(identification) == "" || birthDate.IsZero() {
		return Person{}, fmt.Errorf("%w: all identification fields are required", ErrInvalidPerson)
	}
	if !photo.IsPNG() {
		return Person{}, fmt.Errorf("%w: ID photo must be PNG", ErrInvalidIDPhoto)
	}
	return Person{
		ID: id, FirstName: firstName, LastName: lastName,
		BirthDate: birthDate, Identification: identification, IDPhoto: photo,
	}, nil
}

type Loan struct {
	SkiID     string
	PersonID  string
	StartedAt time.Time
	EndedAt   *time.Time
}

func (loan Loan) Active() bool {
	return loan.EndedAt == nil
}

type ReservationStatus string

const (
	ReservationActive    ReservationStatus = "aktiv"
	ReservationConverted ReservationStatus = "umgewandelt"
	ReservationDiscarded ReservationStatus = "verworfen"
)

type Reservation struct {
	ID        string
	SkiID     string
	PersonID  string
	StartedAt time.Time
	EndedAt   *time.Time
	Status    ReservationStatus
}

func (reservation Reservation) Active() bool {
	return reservation.Status == ReservationActive
}
