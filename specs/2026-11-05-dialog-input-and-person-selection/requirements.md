# Requirements

## Scope Decisions

- Beim Hinzufügen eines Skis werden Marke, Modell und Länge in Zentimetern
  abgefragt.
- Marke, Modell und Länge dürfen nicht leer sein; die Länge muss eine positive
  ganze Zahl sein.
- Abbrechen bei einem beliebigen Feld beendet das vollständige Hinzufügen,
  ohne weitere Dialoge zu öffnen.
- Für Ausleihen und Reservierungen werden Nachname und anschließend Vorname
  abgefragt.
- Die Treffer werden in einem Pop-up als Radio-Auswahl angezeigt; genau eine
  Person muss ausgewählt werden.
- Keine Treffer, leere Eingaben oder Abbrechen beenden den Vorgang ohne
  nachgelagerte Ski-Auswahl.
- Die Auswahl- und Abbruchlogik gilt für Dashboard- und Reservierungs-Buttons
  gleichermaßen.
- Die bestehende API bleibt kompatibel; die Ski-Länge wird als vorhandenes
  `size`-Feld gesendet.

## Fachlicher Kontext

- Dialoge und Aktionen liegen in [`web/app.js`](../../web/app.js).
- Dashboard- und Tab-Buttons liegen in
  [`web/index.html`](../../web/index.html).
- Die API erwartet Ski-Größe in [`internal/handler/api.go`](../../internal/handler/api.go).

## Nicht-Ziele

- Keine Änderung der Domain-ID- oder Statusregeln.
- Keine persistente Speicherung.
- Keine neuen externen Frontend-Abhängigkeiten.
