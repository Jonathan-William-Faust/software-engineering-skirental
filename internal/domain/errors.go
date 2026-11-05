package domain

import "errors"

var (
	ErrInvalidID               = errors.New("invalid domain ID")
	ErrInvalidSki              = errors.New("invalid ski")
	ErrInvalidSkiStatus        = errors.New("invalid ski status")
	ErrInvalidStatusTransition = errors.New("invalid ski status transition")
	ErrInvalidPerson           = errors.New("invalid person")
	ErrInvalidIDPhoto          = errors.New("invalid ID photo")
	ErrSkiNotFound             = errors.New("ski not found")
	ErrPersonNotFound          = errors.New("person not found")
	ErrReservationNotFound     = errors.New("reservation not found")
	ErrSkiUnavailable          = errors.New("ski is not available")
	ErrDuplicateReservation    = errors.New("active reservation already exists")
	ErrActiveLoanExists        = errors.New("active loan already exists")
	ErrLoanNotFound            = errors.New("active loan not found")
	ErrReservationNotActive    = errors.New("reservation is not active")
	ErrInvalidRentalPeriod     = errors.New("invalid rental period")
	ErrInvalidPrice            = errors.New("invalid ski price")
)
