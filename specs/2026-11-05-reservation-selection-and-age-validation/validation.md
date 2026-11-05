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
