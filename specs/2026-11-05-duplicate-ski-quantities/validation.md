# Validation

## Automatisierte Prüfungen

```text
go fmt ./...
go test ./...
go test -race ./...
go vet ./...
git diff --check
```

Zusätzlich müssen Tests nachweisen:

- 200 Start-Ski mit zwei identischen Modell-/Größenkombinationen und
  unterschiedlichen IDs,
- Mengenbegrenzung auf verfügbare konkrete Ski,
- bestehende Ausleih-, Reservierungs-, Rückgabe- und Konvertierungsregeln,
- unveränderte API-/Domain-Regression.

## Manuelle Frontend-Prüfung

1. Im Ausleihdialog für ein Modell/eine Größe die Menge zwei wählen und prüfen,
   dass zwei unterschiedliche Ski-IDs verarbeitet werden.
2. Im Reservierungsdialog dieselbe Mengenwahl und eine zu große Menge prüfen.
3. Für eine Person mehrere Reservierungen anlegen und im Personen-Tab genau
   drei Sammelbuttons prüfen.
4. Sammelumwandlung und Sammellöschung ausführen.
5. Sammelbearbeitung öffnen, bestehende Ski per Checkbox abwählen und neue
   verfügbare Ski per Checkbox auswählen; Zeiträume prüfen.

## Merge-Regel

Merge nur bei bestandenem Test-, Format-, Vet- und Diff-Check sowie sauberem
Arbeitsbaum nach dokumentierter Validierung.
