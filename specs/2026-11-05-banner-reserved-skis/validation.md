# Validation

## Automatisierte Prüfungen

```text
go fmt ./...
go test ./...
go vet ./...
git diff --check
```

Zusätzlich wird geprüft:

- `skirental-banner.png` ist Bestandteil des Feature-Commits.
- `web/index.html` bindet das Banner oberhalb der Navigation ein.
- `web/app.js` enthält die Spalte `Reservierte Ski`, aktive Reservierungen und
  den exakten neuen Zeitraum-Text.

## Regression

Die bestehenden Dashboard-, Personen-, Ski- und Reservierungsfunktionen müssen
unverändert über die vorhandene API funktionieren.
