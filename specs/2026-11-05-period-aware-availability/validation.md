# Validation: Zeitraumbezogene Ski-Verfügbarkeit

Die Umsetzung gilt als erfolgreich, wenn:

1. Überlappende aktive Ausleihen und Reservierungen einen Ski für den
   betroffenen Zeitraum ausschließen.
2. Nicht überlappende Zeiträume und andere verfügbare Ski korrekt angeboten
   werden.
3. Die API eine zeitraumbezogene Ski-Liste liefert und ungültige Zeiträume
   ablehnt.
4. Neue Ausleihen und Reservierungen den Zeitraum vor der Ski-Auswahl abfragen.
5. Eine Reservierung nur am Reservierungsbeginn konvertiert werden kann.
6. `go test ./...`, `go test -race ./...`, `go vet ./...`, `go fmt ./...` und
   `git diff --check` erfolgreich sind.
