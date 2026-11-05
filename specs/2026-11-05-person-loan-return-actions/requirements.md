# Requirements

## Scope Decisions

- Die bestehende Backend-Rückgabe bleibt unverändert, sofern die Tests
  bestätigen, dass sie den aktiven Loan beendet und den Ski auf `verfügbar`
  setzt.
- Die Personenansicht zeigt für jede aktive Ausleihe der Person einen Button
  `Ausleihe beenden`.
- Der Button verwendet `DELETE /api/loans/{skiID}`.
- Nach einer Rückgabe wird `refresh()` ausgeführt, damit Ski-Verwaltung,
  Personen-Tab und Dashboard konsistent aktualisiert werden.
- Die Rückgabe aus der Ski-Verwaltung bleibt erhalten und verwendet denselben
  Aktualisierungspfad.

## Fachlicher Kontext

- Die Anforderungen für Rückgabeaktionen stehen in
  [`README.md`](../../README.md).
- Die Personenansicht und der bestehende Refresh befinden sich in
  [`web/app.js`](../../web/app.js).
- Die Rückgabelogik liegt in
  [`internal/domain/service.go`](../../internal/domain/service.go),
  [`internal/store/store.go`](../../internal/store/store.go) und
  [`internal/handler/api.go`](../../internal/handler/api.go).

## Nicht-Ziele

- Keine Änderung an Loan-, Ski- oder Personenmodellen.
- Keine neue API-Schnittstelle.
- Keine Änderung der Reservierungslogik.
- Keine persistente Datenhaltung.
