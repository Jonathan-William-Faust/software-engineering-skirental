package main

import (
	"encoding/json"
	"log"
	"net/http"

	"skirental/internal/handler"
	"skirental/internal/store"
)

const serverAddr = ":8080"

func main() {
	server, err := newSeededServer()
	if err != nil {
		log.Fatalf("create seeded server: %v", err)
	}

	log.Printf("ski rental server listening on %s", serverAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func newServer() *http.Server {
	return newServerWithStore(store.NewInMemoryStore())
}

func newSeededServer() (*http.Server, error) {
	rentalStore := store.NewInMemoryStore()
	if err := rentalStore.SeedDefaultSkis(); err != nil {
		return nil, err
	}
	return newServerWithStore(rentalStore), nil
}

func newServerWithStore(rentalStore store.Store) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)
	api := handler.NewAPI(rentalStore)
	mux.Handle("/api/", api.Handler())
	mux.Handle("/", http.FileServer(http.Dir("web")))

	return &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
