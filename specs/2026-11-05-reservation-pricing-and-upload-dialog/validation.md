# Validation

## Automatisierte Prüfungen

```text
go fmt ./...
go test ./...
go test -race ./...
go vet ./...
git diff --check
```

Zusätzlich müssen Domain- und API-Tests nachweisen:

- gültige und ungültige Datumsbereiche sowie vergangene Reservierungsstarts,
- Preisgrenzen 10,00 € und 20,00 € sowie Ablehnung außerhalb der Grenzen,
- inklusive Tageszählung und Gesamtpreis,
- zukünftige Reservierung, Doppel-/Überschneidungsablehnung und bestehende
  Statusübergänge,
- Regression von Registrierung, Datei-Upload, sofortiger Ausleihe, Rückgabe,
  Bearbeitung, Verwerfen und Umwandlung einer Reservierung.

## Manuelle Frontend-Prüfung

1. Personenregistrierung öffnen und im sichtbaren Pop-up `Datei auswählen`
   anklicken; eine lokale PNG- oder JPEG-Datei auswählen und Registrierung
   abschließen.
2. Einen Ski mit gültigem Preis anlegen und eine Ausleihe für mehrere Tage
   starten; Zeitraum und berechneten Gesamtpreis prüfen.
3. Eine zukünftige Reservierung anlegen; Start, Ende, Preis und reservierten
   Status prüfen. Eine überschneidende zweite Reservierung muss fehlschlagen.
4. Bestehende Rückgabe-, Bearbeitungs-, Lösch- und
   „Miete reservierter Ski beginnen“-Aktionen ausführen.

## Merge-Regel

Der Merge ist nur zulässig, wenn alle automatisierten Prüfungen erfolgreich
sind, die manuellen Kriterien nachvollziehbar erfüllt sind und
`git diff --check` sowie `git status --short` keine offenen Änderungen außer
den vorgesehenen Validierungsartefakten melden.
