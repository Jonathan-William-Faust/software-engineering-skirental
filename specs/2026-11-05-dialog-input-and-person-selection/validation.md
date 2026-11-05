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
