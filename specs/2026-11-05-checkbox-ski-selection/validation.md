# Validation

1. `go fmt ./...` läuft erfolgreich.
2. `go test ./...` läuft erfolgreich.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. Die Frontend-Prüfung bestätigt, dass das Ausleihe-Pop-up:
   - nur verfügbare Ski als Checkboxen rendert;
   - mehrere Ski akzeptiert;
   - bei leerer Auswahl keine Anfrage sendet;
   - für jeden ausgewählten Ski `POST /api/loans` mit derselben Person-ID
     ausführt.
6. Die drei Workflow-Schritte werden separat committed: Spezifikation,
   Implementierung und Validierung.

## Durchgeführte Nachweise

- `go fmt ./...` und `go test ./...` erfolgreich.
- `go test -race -count=1 ./...` erfolgreich.
- `go vet ./...` erfolgreich.
- `git diff --check` erfolgreich.
- Die statische Prüfung von [`app.js`](../../web/app.js) bestätigt die
  Filterung auf verfügbare Ski, Checkbox-Elemente, die gemeinsame Personen-ID
  und einen `POST /api/loans`-Aufruf pro ausgewähltem Ski.
- Eine JavaScript-Syntaxprüfung mit Node.js war nicht möglich, da Node.js in
  der Umgebung nicht installiert ist.
