package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const serverAddr = ":8080"

func main() {
	server := newServer()

	log.Printf("ski rental server listening on %s", serverAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func newServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)

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
