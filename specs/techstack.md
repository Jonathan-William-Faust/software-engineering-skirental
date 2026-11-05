# Techstack und technische Konstitution

## Grundsatz

Der Prototyp priorisiert Verständlichkeit, geringe Änderungsaufwände und wenige
Abhängigkeiten. Neue Dependencies werden nur nach vorheriger Entscheidung
hinzugefügt.

## Festgelegter Stack

| Bereich | Festlegung |
|---|---|
| Backend | Go mit Structs, Methoden und Interfaces nur bei echtem Bedarf |
| Speicherung | Thread-sicherer In-Memory-Store mit Maps und Mutex oder `sync.Map` |
| Frontend | Reines HTML und JavaScript |
| Datenbank | Keine externe Datenbank und kein ORM |
| Versionierung | Lokales Git |

## Backend-Struktur

- `main.go` oder `cmd/server/` enthält Einstiegspunkt, HTTP-Server und Routing.
- `internal/domain/` enthält Kern-Structs und Geschäftslogik.
- `internal/store/` kapselt den In-Memory-Store.
- `internal/handler/` enthält HTTP-Handler und die Übersetzung zwischen HTTP und
  Domain-Objekten.
- Globale Variablen werden vermieden; ein notwendiger Store-Singleton bleibt die
  einzige zulässige Ausnahme.
- Fehler werden mit `fmt.Errorf("%w", err)` gewrappt und nicht verschluckt.
- Zeitraum- und Statusregeln werden zentral in `internal/domain/` geprüft;
  Handler und Store reichen validierte Vorgaben an diese Geschäftslogik weiter.

## Frontend-Regeln

- HTML, CSS und JavaScript bleiben möglichst klein und klar getrennt.
- Das Frontend kommuniziert ausschließlich über die Backend-HTTP-Schnittstelle.
- Es greift niemals direkt auf In-Memory-Strukturen zu.
- Die Tabs Dashboard, Ski-Verwaltung, Personen und Reservierungen bilden die
  primäre Navigation.
- Jede Aktion öffnet ein Pop-up für Eingaben oder Bestätigung.
- Bei neuen Ausleihen und Reservierungen fragt das Frontend zuerst den Zeitraum
  ab und lädt anschließend die zeitraumbezogene Ski-Auswahl vom Backend.

## Fachliche Datenregeln

- IDs werden serverseitig erzeugt und sind innerhalb des laufenden Systems
  eindeutig.
- Statusänderungen eines Skis erfolgen nur über die Geschäftslogik.
- Eine Ausleihe oder Reservierung darf keine inkonsistente Zuordnung von Ski und
  Person erzeugen.
- Mietzeiträume verwenden Datum ohne Uhrzeit und zählen Start- und Enddatum
  inklusiv. Tagespreise werden als ganze Centbeträge gespeichert.
- Das Ausweisfoto wird als PNG oder JPEG akzeptiert; die Validierung erfolgt vor
  dem Abschluss der Registrierung.
- Zeiträume verwenden inklusive Start- und Enddaten. Überschneidungen mit
  aktiven Ausleihen oder Reservierungen machen einen konkreten Ski für den
  gesamten angefragten Zeitraum nicht verfügbar.
- Die API stellt mit `GET /api/skis/available?start_date=...&end_date=...`
  eine zeitraumbezogene Auswahl bereit.

## Lokale Befehle

```text
go run .
go test ./...
gofmt -w .
go fmt ./...
go test -race ./...
go vet ./...
git diff --check
```

`golangci-lint run` wird verwendet, falls ein Linter eingerichtet ist. Das
Frontend kann direkt geöffnet oder über einen kleinen Static-Server bereitgestellt
werden.
