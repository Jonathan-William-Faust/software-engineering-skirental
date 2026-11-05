# Plan: Thread-sicherer In-Memory-Store

## 1. Store-Vertrag und Zustandszugriff

- Eine Store-Schnittstelle für Ski, Personen, Ausleihen und Reservierungen
  definieren.
- Leseoperationen als konsistente Kopien veröffentlichen.
- Erzeugung und Statusänderungen über die bestehende Domain kapseln.

## 2. In-Memory-Implementierung

- Eine konkrete Implementierung mit Maps und `sync.RWMutex` anlegen.
- Schreiboperationen unter exklusivem Lock und Leseoperationen unter
  Read-Lock ausführen.
- Domain-Fehler unverändert beziehungsweise explizit gewrappt zurückgeben.

## 3. Tests und Nachweise

- Leere Stores, gültige Abläufe und ungültige Übergänge testen.
- Parallelisierte Lese- und Schreibzugriffe mit dem Race Detector ausführen.
- Formatierung, Tests und Diff-Prüfung vor dem Merge validieren.
