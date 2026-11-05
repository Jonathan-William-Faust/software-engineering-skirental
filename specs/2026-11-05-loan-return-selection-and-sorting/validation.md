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
