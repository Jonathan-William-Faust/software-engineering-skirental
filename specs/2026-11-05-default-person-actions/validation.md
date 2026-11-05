# Validation

1. `gofmt -w` beziehungsweise `go fmt ./...` läuft erfolgreich und hinterlässt
   keine nicht formatierten Go-Dateien.
2. `go test ./...` läuft erfolgreich. Dabei wird mindestens abgesichert, dass
   der seeded Store Max Mustermann mit Geburtsdatum `01.01.2000` enthält und
   die bestehende Ausleih-/Reservierungs-API weiterhin funktioniert.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. Ein reproduzierbarer Frontend-Check bestätigt:
   - `Leihe beenden` erscheint nur bei verliehenen Ski und aktualisiert den
     Ski-Status nach Bestätigung.
   - `Neue Ausleihe` erscheint bei Personen ohne aktive Reservierung und fragt
     Ski-ID und Bestätigung ab.
   - Die drei Reservierungsaktionen erscheinen nur bei Personen mit aktiver
     Reservierung und lösen die passenden API-Aufrufe aus.
6. Nach erfolgreicher Validierung werden die drei Workflow-Schritte jeweils
   separat committed: Spezifikation, Implementierung und Validierung. Danach
   wird der Feature-Stand per Fast-Forward nach `main` integriert.
