# Requirements

## Scope Decisions

- Beim Start einer Ausleihe werden ausschließlich Ski mit Status `verfügbar`
  angeboten.
- Die verfügbaren Ski werden in einem Pop-up als Liste mit Checkboxen
  dargestellt.
- Der Benutzer darf einen oder mehrere Ski auswählen.
- Jeder ausgewählte Ski wird als separate Ausleihe für dieselbe Person über die
  bestehende `POST /api/loans`-Schnittstelle gestartet.
- Die Auswahl gilt für:
  - den Dashboard-Schnellzugriff `Ausleihe starten`;
  - die Personenaktion `Neue Ausleihe`.
- Bei keiner Auswahl wird keine API-Anfrage ausgeführt.
- Die vorhandene API und die Domainlogik werden nicht erweitert; die Änderung
  bleibt auf Frontend und Dokumentation beschränkt.

## Fachlicher Kontext

- Die Bedienungsanforderung steht in [`README.md`](../../README.md).
- Die bestehende Leihlogik befindet sich in
  [`web/app.js`](../../web/app.js) und der REST-API unter
  [`internal/handler/api.go`](../../internal/handler/api.go).
- Ski-Statuswerte werden durch [`models.go`](../../internal/domain/models.go)
  festgelegt.

## Nicht-Ziele

- Keine Änderung an Ausleihe-, Rückgabe- oder Reservierungsregeln.
- Keine persistente Datenhaltung.
- Keine neue externe UI-Abhängigkeit.
- Keine Auswahl reservierter oder bereits verliehener Ski.
