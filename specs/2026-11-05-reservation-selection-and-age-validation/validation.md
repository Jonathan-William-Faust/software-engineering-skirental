# Validation

1. `go fmt ./...` läuft erfolgreich.
2. `go test ./...` läuft erfolgreich.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. API-Tests bestätigen, dass zukünftige Geburtsdaten und Personen unter
   18 Jahren abgelehnt werden und der 18. Geburtstag akzeptiert wird.
6. API-/Domain-Tests bestätigen, dass eine Person mehrere aktive
   Reservierungen für unterschiedliche Ski besitzen kann.
7. Die Frontend-Prüfung bestätigt:
   - beide Reservierungsbuttons verwenden die Checkbox-Auswahl;
   - für jeden ausgewählten Ski ein Reservierungs-POST erfolgt;
   - der Dashboard-Button vorhanden ist;
   - der Ausleih-Checkboxpfad erhalten bleibt.
8. Spezifikation, Implementierung und Validierung werden separat committed.

## Durchgeführte Nachweise

- `go fmt ./...` erfolgreich.
- `go test ./...` fachlich erfolgreich; Windows meldete beim anschließenden
  Aufräumen wiederholt einen temporären Testdatei-Lock.
- `go test -race -count=1 ./internal/domain ./internal/handler ./internal/store`
  fachlich erfolgreich; auch hier trat nur der bekannte Windows-Cleanup-Lock
  auf.
- `go test -count=1 .` und `go vet ./...` erfolgreich.
- `git diff --check` erfolgreich.
- Die statische Frontend-Prüfung bestätigt den Dashboard-Button, die
  gemeinsame Checkbox-Funktion für Reservierungen, einen POST pro
  ausgewähltem Ski sowie die Zukunfts- und Mindestaltersprüfungen.
