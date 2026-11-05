# Mission

## Zweck

Das Skiverleihsystem ist eine leicht verständliche Webanwendung zur Verwaltung von
Skiern, Personen, Ausleihen und Reservierungen. Es soll den Verleihprozess von der
Registrierung einer Person bis zur Rückgabe eines Skis nachvollziehbar und mit
wenigen Interaktionen abbilden.

## Produktziele

- Verfügbare, verliehene und reservierte Ski jederzeit übersichtlich anzeigen.
- Personen mit den für den Verleih erforderlichen Identifikationsdaten registrieren.
- Ausleihen sofort beginnen und über eine eindeutige Aktion beenden können.
- Reservierungen erstellen, bearbeiten, in eine Ausleihe umwandeln oder verwerfen.
- Die wichtigsten Kennzahlen im Dashboard auf einen Blick bereitstellen.
- Einen klaren Prototyp liefern, der schnell verständlich und leicht änderbar ist.

## Fachliche Leitplanken

- Ski-IDs verwenden das Format `S-` gefolgt von vier Ziffern.
- Personen-IDs verwenden das Format `P-` gefolgt von vier Ziffern.
- Ein Ski hat genau einen der Status `verfügbar`, `reserviert` oder `verliehen`.
- Ausleihen beginnen sofort und haben ein frei wählbares inklusives Enddatum.
  Reservierungen dürfen in der Zukunft beginnen; ihr Zeitraum wird sofort
  blockiert. Ski verwalten einen Tagespreis.
- Aktionen werden über Pop-ups bestätigt oder mit den erforderlichen Eingaben
  durchgeführt.
- Ein Ausweisfoto ist bei der Registrierung verpflichtend und muss im PNG-Format
  vorliegen.

## Nicht-Ziele

- Persistente Datenhaltung mit PostgreSQL, SQLite oder einem ORM.
- Zahlungs- oder Abrechnungsfunktionen.
- Ein ausgearbeitetes Berechtigungs- oder Benutzerverwaltungssystem.
- Eine produktionsreife Cloud-Bereitstellung.

## Erfolgskriterium

Der Prototyp gilt als erfolgreich, wenn ein Nutzer die beschriebenen Kernabläufe
über Dashboard, Ski-Verwaltung, Personen und Reservierungen ohne manuelle
Änderung interner Daten durchführen kann und die Anzeigen danach konsistent sind.
