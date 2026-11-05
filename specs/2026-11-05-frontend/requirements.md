# Requirements: HTML/JavaScript-Frontend

## Scope Decisions

Diese Spezifikation beschreibt Phase 5 der [roadmap.md](../roadmap.md). Das
Frontend besteht aus statischem HTML, CSS und JavaScript ohne zusätzliche
Dependency oder Framework.

Die Startseite enthält die Tabs Dashboard, Ski-Verwaltung, Personen und
Reservierungen. Sie kommuniziert ausschließlich über die API unter `/api/`,
zeigt die im README geforderten Felder und bietet die beschriebenen Aktionen
über Dialoge an.

## Nicht-Ziele

Nicht enthalten sind Authentifizierung, Preise, Zahlungen, Persistenz,
Frameworks und direkte Zugriffe auf Store- oder Domain-Strukturen.
