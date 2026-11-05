package domain

import (
	"errors"
	"testing"
	"time"
)

func TestIDs(t *testing.T) {
	for _, test := range []struct {
		id       string
		validate func(string) error
		valid    bool
	}{
		{"S-0001", ValidateSkiID, true},
		{"S-12", ValidateSkiID, false},
		{"S-12A4", ValidateSkiID, false},
		{"P-9999", ValidatePersonID, true},
		{"P-10000", ValidatePersonID, false},
	} {
		if got := test.validate(test.id) == nil; got != test.valid {
			t.Errorf("validate %q = %v, want %v", test.id, got, test.valid)
		}
	}
}

func TestSkiStatusTransitions(t *testing.T) {
	ski, err := NewSki("S-0001", "Rossignol", "Hero")
	if err != nil {
		t.Fatal(err)
	}
	if err := ski.Return(); !errors.Is(err, ErrInvalidStatusTransition) {
		t.Fatalf("returning available ski: got %v, want invalid transition", err)
	}
	if err := ski.Reserve(); err != nil {
		t.Fatal(err)
	}
	if err := ski.Rent(); err != nil {
		t.Fatal(err)
	}
	if ski.Status != SkiRented {
		t.Fatalf("status = %q, want %q", ski.Status, SkiRented)
	}
}

func TestPersonRequiresPNGPhoto(t *testing.T) {
	_, err := NewPerson("P-0001", "Ada", "Lovelace", "ID-1", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), IDPhoto{ContentType: "image/jpeg"})
	if !errors.Is(err, ErrInvalidIDPhoto) {
		t.Fatalf("invalid photo error = %v, want ErrInvalidIDPhoto", err)
	}
	person, err := NewPerson("P-0001", "Ada", "Lovelace", "ID-1", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), IDPhoto{ContentType: "image/png"})
	if err != nil {
		t.Fatal(err)
	}
	if person.ID != "P-0001" {
		t.Fatalf("person ID = %q", person.ID)
	}
}

func TestRentalLifecycle(t *testing.T) {
	service := NewRentalService()
	ski, err := service.AddSki("Atomic", "Redster")
	if err != nil {
		t.Fatal(err)
	}
	person, err := service.RegisterPerson("Ada", "Lovelace", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "ID-1", IDPhoto{ContentType: "image/png"})
	if err != nil {
		t.Fatal(err)
	}
	loan, err := service.StartLoan(ski.ID, person.ID)
	if err != nil {
		t.Fatal(err)
	}
	if service.skis[ski.ID].Status != SkiRented {
		t.Fatal("ski was not marked as rented")
	}
	if err := service.EndLoan(loan.SkiID); err != nil {
		t.Fatal(err)
	}
	if service.skis[ski.ID].Status != SkiAvailable {
		t.Fatal("ski was not made available after return")
	}
	if err := service.EndLoan(loan.SkiID); !errors.Is(err, ErrLoanNotFound) {
		t.Fatalf("second return error = %v", err)
	}
}

func TestReservationTransitionsAndConflicts(t *testing.T) {
	service := NewRentalService()
	ski, _ := service.AddSki("Atomic", "Redster")
	otherSki, _ := service.AddSki("Head", "Supershape")
	person, _ := service.RegisterPerson("Ada", "Lovelace", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "ID-1", IDPhoto{ContentType: "image/png"})
	reservation, err := service.CreateReservation(ski.ID, person.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := service.CreateReservation(ski.ID, person.ID); !errors.Is(err, ErrSkiUnavailable) && !errors.Is(err, ErrDuplicateReservation) {
		t.Fatalf("duplicate reservation error = %v", err)
	}
	edited, err := service.EditReservation(reservation.ID, otherSki.ID)
	if err != nil {
		t.Fatal(err)
	}
	if edited.SkiID != otherSki.ID || service.skis[ski.ID].Status != SkiAvailable {
		t.Fatal("editing reservation did not release and reserve the correct skis")
	}
	loan, err := service.ConvertReservation(reservation.ID)
	if err != nil {
		t.Fatal(err)
	}
	if loan.SkiID != otherSki.ID || service.skis[otherSki.ID].Status != SkiRented {
		t.Fatal("reservation conversion did not start the loan")
	}
	if err := service.DiscardReservation(reservation.ID); !errors.Is(err, ErrReservationNotActive) {
		t.Fatalf("discarding converted reservation error = %v", err)
	}
}

func TestReservationCanBeDiscarded(t *testing.T) {
	service := NewRentalService()
	ski, _ := service.AddSki("Atomic", "Redster")
	person, _ := service.RegisterPerson("Ada", "Lovelace", time.Now(), "ID-1", IDPhoto{ContentType: "image/png"})
	reservation, err := service.CreateReservation(ski.ID, person.ID)
	if err != nil {
		t.Fatal(err)
	}
	if err := service.DiscardReservation(reservation.ID); err != nil {
		t.Fatal(err)
	}
	if service.skis[ski.ID].Status != SkiAvailable {
		t.Fatal("discarded reservation did not release ski")
	}
}
