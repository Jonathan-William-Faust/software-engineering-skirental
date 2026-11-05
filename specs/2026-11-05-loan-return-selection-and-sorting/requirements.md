# Requirements

## Scope Decisions

- Im Personen-Tab wird bei mindestens einer aktiven Ausleihe genau ein Button
  `Ausleihe beenden` neben `Neue Ausleihe` angezeigt.
- Der Button öffnet ein Pop-up mit den aktiven Skiern der jeweiligen Person.
- Jeder Ski wird als Checkbox angezeigt; mehrere Ski dürfen ausgewählt werden.
- Für jeden ausgewählten Ski wird `DELETE /api/loans/{skiID}` aufgerufen.
- Nach den Rückgaben werden die Daten per `refresh()` neu geladen.
- Ski werden in jeder Auflistung aufsteigend sortiert nach:
  1. Marke
  2. Modell
  3. Länge
  4. Ski-ID
- Personen werden im Personen-Tab aufsteigend sortiert nach:
  1. Nachname
  2. Vorname
  3. Personen-ID
- Die bestehende Ski-Verwaltungs-Rückgabe bleibt unverändert.

## Fachlicher Kontext

- Die Bedienungsanforderungen stehen in [`README.md`](../../README.md).
- Die Tabellen und Pop-ups werden in [`web/app.js`](../../web/app.js)
  gerendert.
- Die bestehende Rückgabe-API ist in
  [`internal/handler/api.go`](../../internal/handler/api.go) und
  [`internal/domain/service.go`](../../internal/domain/service.go)
  implementiert.

## Nicht-Ziele

- Keine Änderung der Domain- oder REST-API.
- Keine Änderung der Ausleih- oder Reservierungsregeln.
- Keine persistente Speicherung.
- Keine neuen externen Frontend-Abhängigkeiten.
