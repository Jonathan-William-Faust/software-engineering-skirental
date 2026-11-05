# Validation

1. `go fmt ./...` läuft erfolgreich.
2. `go test ./...` läuft erfolgreich.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. Die Frontend-Prüfung bestätigt:
   - Ski-Hinzufügen fragt Marke, Modell und positive Länge ab;
   - leere Eingaben werden nicht an die API gesendet;
   - Abbrechen beendet den Gesamtvorgang;
   - Ausleihe und Reservierung fragen Nachname und Vorname ab;
   - Personentreffer werden als genau-eine-Radio-Auswahl dargestellt;
   - Abbrechen oder keine Treffer öffnen keine Ski-Auswahl.
6. Spezifikation, Implementierung und Validierung werden separat committed.

## Durchgeführte Nachweise

- `go fmt ./...` und `go test ./...` erfolgreich.
- `go test -race -count=1 ./internal/domain ./internal/handler ./internal/store`
  sowie `go test -count=1 .` erfolgreich.
- `go vet ./...` und `git diff --check` erfolgreich.
- Die statische Frontend-Prüfung bestätigt die verpflichtende Längenabfrage,
  `promptRequired`, die Suche über `selectPersonByName`, die Radio-Auswahl
  sowie die gemeinsame Ski-Checkbox-Funktion.
- Node.js ist in der Umgebung nicht installiert; eine JavaScript-Parserprüfung
  konnte daher nicht ausgeführt werden.
