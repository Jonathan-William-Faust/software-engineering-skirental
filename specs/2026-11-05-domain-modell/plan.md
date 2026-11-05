# Plan: Domain-Modell

## 1. Fachliche Domänenobjekte

- `Ski` mit serverseitiger ID, Marke, Modell und einem der drei erlaubten
  Statuswerte definieren.
- `Person` mit ID, Identifikationsdaten und dem verpflichtenden Ausweisfoto
  modellieren.
- `Ausleihe` und `Reservierung` mit eindeutiger Zuordnung von Ski und Person
  sowie ihren fachlichen Zuständen definieren.

## 2. Zentrale Regeln und Werte

- ID-Formate `S-` und `P-` mit jeweils vier Ziffern zentral festlegen.
- Erlaubte Ski-Statuswerte und zulässige Übergänge als typsichere Konstanten
  oder vergleichbare Domain-Werte abbilden.
- Erzeugung und Prüfung der IDs so vorbereiten, dass sie innerhalb des
  laufenden Systems eindeutig bleiben.
- Die verpflichtende PNG-Eigenschaft des Ausweisfotos als fachliche
  Validierungsregel modellieren, ohne Upload- oder HTTP-Logik vorzuziehen.

## 3. Geschäftsregeln und Fehlerfälle

- Registrierung einer Person validieren und fehlerhafte Eingaben explizit
  zurückweisen.
- Ausleihe starten und beenden; dabei nur konsistente Zuordnungen von Ski und
  Person zulassen.
- Reservierung anlegen, bearbeiten, in eine Ausleihe umwandeln oder verwerfen.
- Ungültige Statusübergänge, unbekannte Objekte und doppelte Zuordnungen mit
  benannten Fehlern behandeln und nicht verschlucken.

## 4. Unit-Tests

- Gültige und ungültige ID- sowie Statuswerte testen.
- Erfolgreiche Registrierung, Ausleihe, Rückgabe und Reservierungsübergänge
  testen.
- Fehlerfälle und fachliche Invarianten, insbesondere inkonsistente Ski- und
  Personen-Zuordnungen, absichern.
- Tests unabhängig vom späteren Store und von HTTP-Handlern halten.

