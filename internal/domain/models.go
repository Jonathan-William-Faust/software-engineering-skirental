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
	ID              string    `json:"ID"`
	Brand           string    `json:"Brand"`
	Model           string    `json:"Model"`
	Size            int       `json:"Size"`
	DailyPriceCents int       `json:"DailyPriceCents"`
	Status          SkiStatus `json:"Status"`
}

func NewSki(id, brand, model string) (Ski, error) {
	return NewSkiWithSizeAndPrice(id, brand, model, 0, 1000)
}

func NewSkiWithSize(id, brand, model string, size int) (Ski, error) {
	return NewSkiWithSizeAndPrice(id, brand, model, size, 1000)
}

func NewSkiWithSizeAndPrice(id, brand, model string, size, dailyPriceCents int) (Ski, error) {
	if err := ValidateSkiID(id); err != nil {
		return Ski{}, err
	}
	if strings.TrimSpace(brand) == "" || strings.TrimSpace(model) == "" {
		return Ski{}, fmt.Errorf("%w: ski brand and model are required", ErrInvalidSki)
	}
	if size < 0 {
		return Ski{}, fmt.Errorf("%w: ski size cannot be negative", ErrInvalidSki)
	}
	if dailyPriceCents < 1000 || dailyPriceCents > 2000 {
		return Ski{}, fmt.Errorf("%w: daily price must be between 10.00 and 20.00 EUR", ErrInvalidPrice)
	}
	return Ski{ID: id, Brand: brand, Model: model, Size: size, DailyPriceCents: dailyPriceCents, Status: SkiAvailable}, nil
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
	DataURL     string
}

func (photo IDPhoto) IsSupported() bool {
	contentType := strings.ToLower(strings.TrimSpace(photo.ContentType))
	return contentType == "image/png" || contentType == "image/jpeg"
}

type Person struct {
	ID             string
	FirstName      string
	LastName       string
	BirthDate      time.Time
	Identification string
	Country        string
	City           string
	PostalCode     string
	Street         string
	HouseNumber    string
	MobileNumber   string
	IDPhoto        IDPhoto
}

func NewPerson(id, firstName, lastName, identification string, birthDate time.Time, country, city, postalCode, street, houseNumber, mobileNumber string, photo IDPhoto) (Person, error) {
	if err := ValidatePersonID(id); err != nil {
		return Person{}, err
	}
	if strings.TrimSpace(firstName) == "" || strings.TrimSpace(lastName) == "" ||
		strings.TrimSpace(identification) == "" || birthDate.IsZero() ||
		strings.TrimSpace(country) == "" || strings.TrimSpace(city) == "" || strings.TrimSpace(street) == "" ||
		strings.TrimSpace(houseNumber) == "" || strings.TrimSpace(mobileNumber) == "" {
		return Person{}, fmt.Errorf("%w: all identification fields are required", ErrInvalidPerson)
	}
	if !validPostalCode(postalCode) {
		return Person{}, fmt.Errorf("%w: postal code must contain four or five digits", ErrInvalidPerson)
	}
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if birthDate.After(today) {
		return Person{}, fmt.Errorf("%w: birth date cannot be in the future", ErrInvalidPerson)
	}
	if birthDate.AddDate(18, 0, 0).After(today) {
		return Person{}, fmt.Errorf("%w: person must be at least 18 years old", ErrInvalidPerson)
	}
	if birthDate.AddDate(100, 0, 0).Before(today) {
		return Person{}, fmt.Errorf("%w: person must not be older than 100 years", ErrInvalidPerson)
	}
	if !photo.IsSupported() {
		return Person{}, fmt.Errorf("%w: ID photo must be PNG or JPEG", ErrInvalidIDPhoto)
	}
	return Person{
		ID: id, FirstName: firstName, LastName: lastName,
		BirthDate: birthDate, Identification: identification,
		Country: country, City: city, PostalCode: postalCode, Street: street,
		HouseNumber: houseNumber, MobileNumber: mobileNumber, IDPhoto: photo,
	}, nil
}

func validPostalCode(value string) bool {
	if len(value) != 4 && len(value) != 5 {
		return false
	}
	for _, character := range value {
		if character < '0' || character > '9' {
			return false
		}
	}
	return true
}

type Loan struct {
	SkiID           string     `json:"SkiID"`
	PersonID        string     `json:"PersonID"`
	StartedAt       time.Time  `json:"StartedAt"`
	EndDate         time.Time  `json:"EndDate"`
	EndedAt         *time.Time `json:"EndedAt"`
	Days            int        `json:"Days"`
	TotalPriceCents int        `json:"TotalPriceCents"`
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
	ID              string            `json:"ID"`
	SkiID           string            `json:"SkiID"`
	PersonID        string            `json:"PersonID"`
	StartedAt       time.Time         `json:"StartedAt"`
	EndDate         time.Time         `json:"EndDate"`
	DailyPriceCents int               `json:"DailyPriceCents"`
	EndedAt         *time.Time        `json:"EndedAt"`
	Status          ReservationStatus `json:"Status"`
}

func (reservation Reservation) Active() bool {
	return reservation.Status == ReservationActive
}
