# Validation

1. `go fmt ./...` läuft erfolgreich.
2. `go test ./...` läuft erfolgreich.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. Domain-/API-Tests bestätigen die neuen Pflichtfelder und die PLZ-Regel.
6. API-/Frontend-Prüfungen bestätigen, dass ungültige PNG-Fotos abgelehnt
   werden und die UI die Fotoeingabe nach einem Pop-up wiederholt.
7. Die Personenauswahl zeigt Name, Geburtsdatum, Adresse und Personen-ID und
   erlaubt genau eine Auswahl.
8. Spezifikation, Implementierung und Validierung werden separat committed.
