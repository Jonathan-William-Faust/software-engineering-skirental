# Validation

1. `go fmt ./...` läuft erfolgreich.
2. `go test ./...` läuft erfolgreich.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. Die Frontend-Prüfung bestätigt:
   - höchstens ein Rückgabe-Button pro Person gerendert wird;
   - das Rückgabe-Pop-up Checkboxen für alle aktiven Leihen der Person
     enthält;
   - pro ausgewähltem Ski die DELETE-API aufgerufen wird;
   - nach der Aktion `refresh()` ausgeführt wird;
   - Ski nach Marke, Modell, Länge und Ski-ID sortiert werden;
   - Personen nach Nachname, Vorname und Personen-ID sortiert werden.
6. Spezifikation, Implementierung und Validierung werden separat committed.

## Durchgeführte Nachweise

- `go fmt ./...` und `go test ./...` erfolgreich.
- `go test -race -count=1 ./...` erfolgreich.
- `go vet ./...` und `git diff --check` erfolgreich.
- Die statische Frontend-Prüfung bestätigt genau einen
  `data-person-return`-Button pro Person mit aktiven Leihen, das
  `selectPersonLoans`-Pop-up mit `loan-ski`-Checkboxen sowie einen
  `DELETE /api/loans/{skiID}`-Aufruf je Auswahl.
- Die Sortierfunktionen verwenden für Ski Marke, Modell, numerische Länge und
  Ski-ID sowie für Personen Nachname, Vorname und Personen-ID.
