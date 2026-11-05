# Plan: HTTP-API

## 1. API-Verträge und Routing

- REST-Ressourcen für Dashboard, Ski, Personen, Ausleihen und Reservierungen
  definieren.
- JSON-Anfragen und JSON-Antworten mit passenden HTTP-Statuscodes abbilden.
- Ressourcenpfade und Aktionspfade zentral im HTTP-Server registrieren.

## 2. Handler und Validierung

- Handler als dünne Übersetzung zwischen HTTP, Store und Domain implementieren.
- Pflichtfelder, IDs und PNG-Foto-Metadaten validieren.
- Domain-Fehler verständlich als 4xx-Antworten ausgeben.

## 3. API-Tests

- Erfolgreiche Kernabläufe und ungültige Eingaben mit `httptest` absichern.
- Dashboard-Zähler nach Statusänderungen prüfen.
- Bestehenden Health-Endpunkt unverändert halten.
