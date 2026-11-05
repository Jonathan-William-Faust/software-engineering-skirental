# Requirements: Thread-sicherer In-Memory-Store

## Scope Decisions

Diese Spezifikation beschreibt Phase 3 der
[roadmap.md](../roadmap.md). Der Store kapselt den laufenden Domain-Zustand
für Ski, Personen, Ausleihen und Reservierungen und schützt ihn vor
konkurrierenden Zugriffen.

Enthalten sind die Store-Schnittstelle, eine In-Memory-Implementierung,
Leseoperationen, fachliche Schreiboperationen und Unit-Tests. Nicht enthalten
sind HTTP-Routing, Handler, Frontend, Persistenz, Preise, Zahlungen und
Berechtigungen.

## Fachlicher Kontext

Die Domain-Regeln aus [mission.md](../mission.md) und dem
[Domain-Modell](../2026-11-05-domain-modell/requirements.md) bleiben
verbindlich. Der Store darf Zustandsänderungen nicht an der Domain vorbei
durchführen. IDs werden weiterhin serverseitig durch die Domain erzeugt.

## Technische Entscheidungen

- Der Store wird mit Go-Maps und `sync.RWMutex` implementiert.
- Leseoperationen geben Werte beziehungsweise Kopien zurück, damit Aufrufer
  keine internen Maps oder Pointer verändern können.
- Schreiboperationen validieren unbekannte Objekte, doppelte Zuordnungen und
  ungültige Übergänge durch die Domain.
- Es werden keine neuen Dependencies eingeführt.
