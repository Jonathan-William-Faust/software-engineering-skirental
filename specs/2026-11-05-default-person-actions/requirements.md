# Requirements

## Scope Decisions

- Beim Start wird genau eine Standardperson angelegt:
  - Vorname: `Max`
  - Nachname: `Mustermann`
  - Geburtsdatum: `01.01.2000`
  - Eine gültige PNG-Ausweisfoto-Metadatenrepräsentation wird verwendet.
- Die Ski-Verwaltung enthält eine Aktionsspalte. Für verliehene Ski wird
  `Leihe beenden` angezeigt und beendet die aktive Ausleihe.
- Im Personen-Tab werden die im README beschriebenen kontextabhängigen
  Aktionsbuttons angezeigt:
  - `Neue Ausleihe` für Personen ohne aktive Reservierung.
  - `Miete reservierter Ski beginnen` für Personen mit aktiver Reservierung.
  - `Reservierung bearbeiten` für Personen mit aktiver Reservierung.
  - `Reservierung löschen` für Personen mit aktiver Reservierung.
- Aktionen verwenden die vorhandenen REST-Endpunkte und aktualisieren die
  betroffenen Ansichten nach erfolgreicher Ausführung.
- Die bestehenden uncommitted Änderungen zur deutschen Datumseingabe sind
  unabhängig und werden nicht erneut verändert.

## Fachlicher Kontext

- Die Anforderungen für die Aktionsbuttons stehen in
  [`README.md`](../../README.md).
- Der produktive Server verwendet den seeded Store in
  [`main.go`](../../main.go).
- Domain- und Store-Geschäftslogik liegt in
  [`internal/domain`](../../internal/domain) und
  [`internal/store/store.go`](../../internal/store/store.go).
- Das Frontend kommuniziert ausschließlich über die API in
  [`web/app.js`](../../web/app.js).

## Nicht-Ziele

- Keine persistente Datenbank.
- Keine neuen Preis-, Zahlungs- oder Berechtigungsfunktionen.
- Keine Änderung der bestehenden Reservierungs- oder Ausleihregeln.
- Keine Änderung des Datumsformats oder der bereits vorhandenen Ski-Katalogdaten.

