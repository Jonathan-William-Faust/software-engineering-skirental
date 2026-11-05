# Plan: Ski-ID-Anzeige

## 1. Ursache beheben

- Die JSON-Schreibweise der Go-API und die Feldauflösung des Frontends
  abgleichen.
- Die ID-Anzeige für hinterlegte und neu angelegte Ski korrigieren.

## 2. Eindeutigkeit absichern

- Prüfen, dass der Startkatalog ausschließlich nicht-leere und eindeutige
  Ski-IDs enthält.
- Prüfen, dass ein vom Benutzer angelegter Ski eine serverseitige ID im
  Format `S-0000` erhält.

## 3. Dokumentation und Nachweis

- Ursache und Korrektur dokumentieren.
- Bestehende Go-, Race-, Format- und Diff-Prüfungen erneut ausführen.
