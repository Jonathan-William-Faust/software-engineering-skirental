# Requirements: Domain-Modell

## Scope Decisions

Diese Spezifikation beschreibt ausschließlich Phase 2 der
[roadmap.md](../roadmap.md): die fachlichen Structs, zentralen Werte,
Geschäftsregeln und Unit-Tests im Go-Domain-Paket.

Im Scope enthalten sind Ski, Personen, Ausleihen und Reservierungen sowie deren
Validierung und Zustandsübergänge. Nicht enthalten sind der In-Memory-Store,
HTTP-Routing, Handler, Frontend, Persistenz, Preise, Zahlungen und
Berechtigungen. Die späteren Schichten dürfen die Domain verwenden, aber nicht
vorweggenommen werden.

## Fachlicher Kontext

Die Leitplanken aus [mission.md](../mission.md) gelten verbindlich:

- Ski-IDs haben das Format `S-` plus vier Ziffern.
- Personen-IDs haben das Format `P-` plus vier Ziffern.
- Ein Ski ist genau `verfügbar`, `reserviert` oder `verliehen`.
- Ausleihen und Reservierungen beginnen sofort; Preise werden nicht verwaltet.
- Ein Ausweisfoto ist für die Registrierung verpflichtend und muss PNG sein.

Die technische Umsetzung folgt
[techstack.md](../techstack.md): Go-Structs und Methoden im
`internal/domain/`-Paket, keine neue Dependency, keine globale Zustandsablage
und explizites Error-Wrapping mit `fmt.Errorf("%w", err)`, wo Fehler
weitergereicht werden.

## Domänenentscheidungen

- Statusänderungen eines Skis sind nur über fachliche Methoden möglich.
- Eine aktive Ausleihe oder Reservierung verbindet genau eine Person mit genau
  einem Ski.
- Ein reservierter Ski darf nicht gleichzeitig verliehen oder einer zweiten
  aktiven Reservierung zugeordnet sein.
- Eine Reservierung kann nur für einen verfügbaren Ski angelegt werden; ihre
  Umwandlung startet unmittelbar die Ausleihe und gibt die Reservierung frei.
- Eine Rückgabe beendet die Ausleihe und setzt den Ski auf `verfügbar`.
- Die Domain prüft die PNG-Eigenschaft anhand der übergebenen Foto-Metadaten
  beziehungsweise eines klar definierten Domain-Werts; Datei-Upload und
  MIME-Erkennung bleiben Aufgabe späterer Schichten.
- IDs werden serverseitig erzeugt; die konkrete Sequenzverwaltung wird so
  geschnitten, dass sie später vom Store beziehungsweise einer Factory
  bereitgestellt werden kann.

## Skill-Scope

Der Skill `writing-feature-spec` liegt als Workspace-Skill unter
`.github/skills/writing-feature-spec/`. Er automatisiert die wiederkehrende Erstellung eines datierten
`specs/YYYY-MM-DD-taskname/`-Ordners mit `plan.md`, `requirements.md` und
`validation.md`. Vor der Erstellung fragt er die Entscheidungen ausdrücklich
gruppiert nach diesen drei Dateien über das AskQuestionTool ab.
