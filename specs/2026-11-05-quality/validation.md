# Validation: Qualität und Abschluss

Die Phase gilt als erfolgreich, wenn:

1. README Start, Frontend, API und Entwicklungsbefehle nachvollziehbar
   dokumentiert.
2. `go fmt ./...`, `go test ./...`, `go test -race ./...`, `go vet ./...`
   und `git diff --check` erfolgreich sind.
3. Ein vorhandener Linter verwendet wird; falls keiner eingerichtet ist, wird
   dies als Hinweis dokumentiert.
4. Keine neuen Dependencies oder ungewollten Dateien eingeführt wurden.
5. Der finale Branch nach erfolgreicher Prüfung per Fast-Forward nach `main`
   gemergt werden kann.
