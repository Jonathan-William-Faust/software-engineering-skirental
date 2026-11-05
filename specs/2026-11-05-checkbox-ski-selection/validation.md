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
