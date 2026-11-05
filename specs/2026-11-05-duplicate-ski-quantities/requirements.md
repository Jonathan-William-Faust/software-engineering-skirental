# Requirements

## Scope Decisions

- Der Standardkatalog enthält 20 Ski je Modell und damit 200 Ski insgesamt:
  jedes der zehn Modelle wird in jeder Größe zweimal angelegt.
- Ski derselben Marke, desselben Modells und derselben Größe erhalten weiterhin
  unterschiedliche Ski-IDs.
- Ausleihe und Reservierung wählen die Anzahl je Modell-/Größenkombination.
  Die Anwendung weist dafür verfügbare konkrete Ski zu; eine Menge größer als
  der verfügbare Bestand wird abgelehnt.
- Eine Person darf mehrere konkrete Ski desselben Modells reservieren, sofern
  diese Ski verfügbar sind.
- Im Personen-Tab existiert je Person mit aktiven Reservierungen genau ein
  Button „Miete reservierter Ski beginnen“, genau ein Button
  „Reservierung bearbeiten“ und genau ein Button „Reservierung löschen“.
- „Miete reservierter Ski beginnen“ wandelt alle aktiven Reservierungen dieser
  Person um.
- „Reservierung löschen“ verwirft alle aktiven Reservierungen dieser Person.
- „Reservierung bearbeiten“ zeigt bestehende Reservierungen als Checkboxen,
  sodass einzelne Ski abgewählt werden können, und verfügbare konkrete Ski als
  Checkboxen für neue Auswahl. Beibehaltene Reservierungen behalten ihre
  Zeiträume; neue Reservierungen erhalten einen gemeinsam abgefragten Zeitraum.
- Bestehende Einzelaktionen im Reservierungs-Tab bleiben unverändert.
- Die README-Änderung aus dem Arbeitsbaum wird übernommen und mit diesem
  Feature veröffentlicht.

## Nicht-Ziele

- Keine Änderung der Ski-ID-Erzeugung.
- Keine freie Änderung eines Zeitraums bei beibehaltenen Einzelreservierungen.
- Keine persistente Speicherung oder neue Abhängigkeit.

## Fachlicher Kontext

- Katalogerzeugung liegt in `internal/store/store.go`.
- Ski-Auswahl und Personenaktionen liegen in `web/app.js`.
- Reservierungsstatus und konkrete Ski-Zuordnung bleiben in
  `internal/domain/` und der bestehenden API.
