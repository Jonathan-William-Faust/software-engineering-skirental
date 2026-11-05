# Requirements: Integration und fachliche Durchläufe

## Scope Decisions

Diese Spezifikation beschreibt Phase 6 der [roadmap.md](../roadmap.md).
Frontend, HTTP-API, Store und Domain werden gemeinsam über den bestehenden
Server geprüft. Die Tests verwenden `httptest` und bleiben lokal sowie
dependency-frei.

## Nicht-Ziele

Nicht enthalten sind Browser-Automatisierung, externe Systeme, Persistenz,
Lasttests und Produktionsdeployment.
