# Requirements

## Scope Decisions

- Der Dashboard-Schnellzugriff erhält den Button `Reservierung anlegen`.
- Der Button im Dashboard und der Button im Reservierungs-Tab verwenden dieselbe
  Funktion.
- Die Funktion fragt eine Personen-ID ab und zeigt anschließend alle
  verfügbaren Ski als sortierte Checkbox-Liste an.
- Für jeden ausgewählten Ski wird eine eigene Reservierung für dieselbe Person
  angelegt.
- Eine Person darf beliebig viele aktive Reservierungen haben, sofern jeweils
  ein anderer verfügbarer Ski reserviert wird.
- Die Ausleihe über `Ausleihe starten` im Dashboard verwendet weiterhin die
  vorhandene Checkbox-Auswahl für mehrere Ski.
- Geburtsdaten in der Zukunft werden abgelehnt.
- Personen, die am aktuellen Tag noch nicht 18 Jahre alt sind, werden
  abgelehnt. Der 18. Geburtstag ist zulässig.
- Die Alters- und Zukunftsprüfung erfolgt sowohl im Frontend als auch im
  Backend.

## Fachlicher Kontext

- UI und Checkbox-Auswahl liegen in [`web/app.js`](../../web/app.js).
- Der Dashboard-Schnellzugriff liegt in
  [`web/index.html`](../../web/index.html).
- Reservierungsregeln liegen in
  [`internal/domain/service.go`](../../internal/domain/service.go).
- API-Datumsvalidierung liegt in
  [`internal/handler/api.go`](../../internal/handler/api.go).

## Nicht-Ziele

- Keine persistente Datenhaltung.
- Keine neuen externen Frontend-Abhängigkeiten.
- Keine Änderung an Preis- oder Berechtigungsfunktionen.
- Keine Reservierung nicht verfügbarer Ski.
