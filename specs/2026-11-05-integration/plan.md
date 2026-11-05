# Plan: Integration und fachliche Durchläufe

## 1. Server- und UI-Erreichbarkeit

- Startseite, Health-Endpunkt und API gemeinsam über den Server prüfen.
- Einen reproduzierbaren HTTP-Testaufbau mit `httptest` verwenden.

## 2. Fachliche Durchläufe

- Person registrieren und Ski anlegen.
- Ausleihe starten und zurückgeben.
- Reservierung anlegen, bearbeiten, umwandeln und verwerfen.

## 3. Konsistenznachweise

- Dashboard-Zähler nach jeder Zustandsänderung prüfen.
- Fehlerfälle und Statusänderungen über öffentliche HTTP-Schnittstellen
  absichern.
