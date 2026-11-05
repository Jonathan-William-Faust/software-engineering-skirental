# Requirements: HTTP-API

## Scope Decisions

Diese Spezifikation beschreibt Phase 4 der [roadmap.md](../roadmap.md).
Die HTTP-Schicht stellt die fachlichen Abläufe des Domain-Modells über den
thread-sicheren Store bereit.

Standardisierte Ressourcen:

- `GET /api/dashboard`
- `GET|POST /api/skis`
- `GET|POST /api/people`
- `GET|POST /api/loans` sowie `DELETE /api/loans/{skiID}`
- `GET|POST /api/reservations`
- `PUT|DELETE /api/reservations/{id}`
- `POST /api/reservations/{id}/convert`

JSON ist das Austauschformat. Personen senden für das Ausweisfoto zunächst
Foto-Metadaten mit `content_type: "image/png"`; Binär-Upload und MIME-Erkennung
bleiben späteren Schichten vorbehalten.

## Nicht-Ziele

Nicht enthalten sind Authentifizierung, Berechtigungen, Persistenz, Preise,
Zahlungen, Frontend, Binärdateiablage und externe Dependencies.
