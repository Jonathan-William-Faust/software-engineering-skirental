# Plan: Zeitraumbezogene Ski-Verfügbarkeit

## 1. Verfügbarkeit fachlich bündeln

- Eine zentrale Domain-Prüfung für Ski-Verfügbarkeit über einen inklusiven
  Zeitraum ergänzen.
- Aktive Ausleihen und Reservierungen mit überschneidendem Zeitraum
  ausschließen.
- Die Konvertierung einer Reservierung auf den tatsächlichen Reservierungsbeginn
  beschränken.

## 2. API und Bedienablauf anpassen

- Eine API-Abfrage für verfügbare Ski innerhalb eines Zeitraums bereitstellen.
- Bei neuen Ausleihen und Reservierungen den Zeitraum vor der Ski-Auswahl
  abfragen.
- Nur die vom Backend bestätigten Ski im Auswahl-Pop-up anzeigen.

## 3. Nachweis und Rückwärtskompatibilität

- Domain-, Handler- und Integrationsregressionen ergänzen.
- README und bestehende Abläufe konsistent halten.
- Alle vorgesehenen Go- und Diff-Prüfungen ausführen.
