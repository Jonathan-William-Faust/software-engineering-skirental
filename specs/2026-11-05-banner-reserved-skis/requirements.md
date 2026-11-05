# Requirements

## Scope Decisions

- `skirental-banner.png` wird oberhalb der Reiter und am oberen Ende der App
  angezeigt. Das Bild streckt sich über die gesamte verfügbare Fensterbreite,
  bleibt responsiv und verzerrt sich nicht.
- Die Personen-Tabelle erhält eine Spalte mit der exakten Überschrift
  `Reservierte Ski`.
- Für jede Person werden dort ausschließlich aktive Reservierungen angezeigt.
  Jeder Eintrag enthält die konkrete Ski-ID sowie Modell und Größe, sofern
  verfügbar.
- Der Zeitraum-Fehlertext lautet exakt:
  `Ungültiger Zeitraum. Bitte verwenden Sie gültige Daten; das Ende muss am oder nach dem Beginn liegen. Ausleihen können nur sofort beginnen.`
- Backend-Verhalten und bestehende Reservierungs-/Ausleihlogik bleiben
  unverändert.

## Nicht-Ziele

- Keine Änderung der API oder Datenmodelle.
- Keine neue externe Abhängigkeit.
