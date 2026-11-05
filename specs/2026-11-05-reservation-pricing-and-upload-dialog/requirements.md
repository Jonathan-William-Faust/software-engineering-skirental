# Requirements

## Scope Decisions

- Die Ausweisregistrierung öffnet ein sichtbares Dialog-Pop-up mit dem Text
  „Datei auswählen“. Darin kann der native lokale Dateiauswahldialog geöffnet
  werden. Der bisherige PNG-/JPEG-Upload, die Validierung und die Wiederholung
  bei ungültigen Dateien bleiben erhalten.
- Eine Reservierung benötigt ein Start- und Enddatum. Beide Daten werden
  inklusive gezählt, das Enddatum darf nicht vor dem Startdatum liegen, und der
  Start darf nicht in der Vergangenheit liegen.
- Reservierungen können für zukünftige Zeiträume angelegt werden und markieren
  den Ski unmittelbar als `reserviert`. Ein reservierter Ski darf nicht
  überlappend erneut reserviert oder ausgeliehen werden.
- Eine Ausleihe startet weiterhin sofort, benötigt aber ein Enddatum. Start-
  und Enddatum werden inklusive gezählt; die Mindestdauer beträgt einen Tag.
- Jeder Ski besitzt einen Tagespreis von 10,00 € bis 20,00 €. Der Preis wird
  beim Anlegen erfasst und serverseitig validiert. Der Gesamtpreis ist
  Tagespreis mal Anzahl der Miettager.
- Das Anlegen eines Skis, die Ausleihe, Rückgabe, Reservierung, Umwandlung und
  Bearbeitung müssen mit den bisherigen Status- und Personenzuordnungen
  weiterhin funktionieren.
- Die API liefert Tagespreis, Startdatum, Enddatum, Miettage und Gesamtpreis
  für Ausleihen sowie Tagespreis und Zeitraum für Reservierungen.

## Nicht-Ziele

- Keine zukünftigen Ausleihstarts.
- Keine Teilzeiträume oder Uhrzeitgenauigkeit.
- Keine Zahlungsabwicklung, Rechnungsstellung oder Preisrabatte.
- Keine externe Datenbank oder neue Frontend-Abhängigkeit.

## Fachlicher Kontext

- Domain-Regeln liegen in `internal/domain/`.
- Store-Abbildung und HTTP-Übersetzung liegen in `internal/store/` und
  `internal/handler/api.go`.
- Dialoge und Anzeige liegen in `web/index.html` und `web/app.js`.
