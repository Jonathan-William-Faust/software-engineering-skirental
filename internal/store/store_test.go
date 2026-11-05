package store

import (
	"errors"
	"sync"
	"testing"
	"time"

	"skirental/internal/domain"
)

func TestEmptyStoreReturnsEmptyCollections(t *testing.T) {
	store := NewInMemoryStore()
	if store.ListSkis() == nil || store.ListPeople() == nil ||
		store.ListLoans() == nil || store.ListReservations() == nil {
		t.Fatal("empty store returned nil collection")
	}
	if len(store.ListSkis()) != 0 || len(store.ListPeople()) != 0 ||
		len(store.ListLoans()) != 0 || len(store.ListReservations()) != 0 {
		t.Fatal("empty store returned unexpected data")
	}
}

func TestDefaultSkiCatalogContains100Skis(t *testing.T) {
	store := NewInMemoryStore()
	if err := store.SeedDefaultSkis(); err != nil {
		t.Fatal(err)
	}
	skis := store.ListSkis()
	if len(skis) != 100 {
		t.Fatalf("seeded ski count = %d, want 100", len(skis))
	}
	models := make(map[string]map[int]bool)
	for _, ski := range skis {
		if ski.Size < 150 || ski.Size > 195 {
			t.Fatalf("ski %s has unexpected size %d", ski.ID, ski.Size)
		}
		if models[ski.Model] == nil {
			models[ski.Model] = make(map[int]bool)
		}
		models[ski.Model][ski.Size] = true
	}
	if len(models) != 10 {
		t.Fatalf("seeded model count = %d, want 10", len(models))
	}
	for model, sizes := range models {
		if len(sizes) != 10 {
			t.Fatalf("model %q has %d sizes, want 10", model, len(sizes))
		}
	}
}

func TestStoreLifecycleAndErrors(t *testing.T) {
	store := NewInMemoryStore()
	ski, err := store.AddSki("Atomic", "Redster")
	if err != nil {
		t.Fatal(err)
	}
	person, err := store.RegisterPerson("Ada", "Lovelace", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "ID-1", domain.IDPhoto{ContentType: "image/png"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.StartLoan("S-9999", person.ID); !errors.Is(err, domain.ErrSkiNotFound) {
		t.Fatalf("unknown ski error = %v", err)
	}
	if _, err := store.StartLoan(ski.ID, person.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := store.StartLoan(ski.ID, person.ID); !errors.Is(err, domain.ErrSkiUnavailable) {
		t.Fatalf("unavailable ski error = %v", err)
	}
	if err := store.EndLoan(ski.ID); err != nil {
		t.Fatal(err)
	}
	if err := store.EndLoan(ski.ID); !errors.Is(err, domain.ErrLoanNotFound) {
		t.Fatalf("second return error = %v", err)
	}
	reservation, err := store.CreateReservation(ski.ID, person.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.CreateReservation(ski.ID, person.ID); !errors.Is(err, domain.ErrSkiUnavailable) {
		t.Fatalf("duplicate reservation error = %v", err)
	}
	if err := store.DiscardReservation(reservation.ID); err != nil {
		t.Fatal(err)
	}
}

func TestStoreSupportsConcurrentAccess(t *testing.T) {
	store := NewInMemoryStore()
	var writer sync.WaitGroup
	for range 20 {
		writer.Add(1)
		go func() {
			defer writer.Done()
			_, _ = store.AddSki("Atomic", "Redster")
		}()
	}
	var reader sync.WaitGroup
	for range 20 {
		reader.Add(1)
		go func() {
			defer reader.Done()
			_ = store.ListSkis()
			_ = store.ListPeople()
		}()
	}
	writer.Wait()
	reader.Wait()
	if got := len(store.ListSkis()); got != 20 {
		t.Fatalf("ski count = %d, want 20", got)
	}
}
