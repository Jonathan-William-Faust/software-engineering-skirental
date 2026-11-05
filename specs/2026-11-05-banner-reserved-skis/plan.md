# Plan: Banner und reservierte Ski

## Ziel

Die Weboberfläche erhält ein vollbreites Banner am oberen Rand. Zusätzlich
zeigt die Personenansicht die aktuell reservierten konkreten Ski je Person.
Der Zeitraum-Fehlertext wird an die fachliche Regel angepasst.

## Umsetzung

1. Das vorhandene `skirental-banner.png` als responsives Banner oberhalb der
   Navigation einbinden und visuell über die gesamte Fensterbreite strecken.
2. Die Personen-Tabelle um die Spalte „Reservierte Ski“ erweitern und aktive
   Reservierungen mit Ski-ID, Modell und Größe anzeigen.
3. Den exakten Zeitraum-Hinweis in der bestehenden Eingabevalidierung ersetzen.
4. Bestehende Tests und Qualitätsprüfungen ausführen.

## Abgrenzung

- Keine Änderung an Domain, Store oder API.
- Keine Änderung an Reservierungsregeln oder bestehenden Aktionen.
