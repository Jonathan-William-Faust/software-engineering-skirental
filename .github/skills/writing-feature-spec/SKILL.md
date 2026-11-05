---
name: writing-feature-spec
description: "Erstellt eine datierte Feature-Spezifikation aus Roadmap, Mission, Techstack und README. Verwende den Skill für plan.md, requirements.md und validation.md vor einer Implementierung."
argument-hint: "[optional: Taskname oder Roadmap-Phase]"
user-invocable: true
---

# Writing Feature Spec

## Zweck

Dieser Skill standardisiert die Vorbereitung des nächsten Roadmap-Tasks. Er
erstellt `specs/YYYY-MM-DD-taskname/` mit genau drei Kernartefakten:
`plan.md`, `requirements.md` und `validation.md`.

## Vorgehen

1. Lies `README.md`, `specs/mission.md`, `specs/techstack.md` und
   `specs/roadmap.md`. Prüfe außerdem die zuletzt abgeschlossene
   Spezifikation unter `specs/`.
2. Bestimme den nächsten offenen Roadmap-Task. Wenn mehrere Tasks gleichrangig
   sind oder der Scope unklar ist, frage nach, statt eine Phase zu überspringen.
3. Verwende das AskQuestionTool zwingend **vor jeder Änderung**. Stelle eine
   einzige gruppierte Fragerunde mit drei Abschnitten und den Feldern
   `plan.md`, `requirements.md` und `validation.md`:
   - `plan.md`: Task, Aufgabengruppen und gewünschte Reihenfolge.
   - `requirements.md`: Scope Decisions, Kontext, Nicht-Ziele und gewünschter
     Speicherort des Skills.
   - `validation.md`: messbare Erfolgs-, Test- und Merge-Kriterien.
4. Leite aus den Antworten einen kurzen, kebab-case-fähigen Tasknamen ab und
   erstelle den Ordner `specs/YYYY-MM-DD-taskname/`.
5. Schreibe `plan.md` als nummerierte Serie von Aufgabengruppen. Jede Gruppe
   enthält konkrete, überprüfbare Arbeitsschritte.
6. Schreibe `requirements.md` mit Scope Decisions, fachlichem Kontext,
   Nicht-Zielen und Verweisen auf die maßgeblichen Projektdateien.
7. Schreibe `validation.md` mit nummerierten Kriterien, konkreten Befehlen
   beziehungsweise Prüfnachweisen und einer eindeutigen Merge-Entscheidung.
8. Prüfe, dass die drei Dateien den bestehenden Leitplanken nicht
   widersprechen und keine Implementierung vorgreifen.
9. Wenn der Nutzer einen Branch verlangt, erstelle vor den Änderungen einen
   neuen Branch mit dem Muster `feature/YYYY-MM-DD-taskname`. Führe keine
   Commits oder Merges ohne ausdrückliche Aufforderung aus.

## Qualitätsregeln

- Verwende relative Markdown-Links innerhalb der Spezifikationsordner.
- Trenne Scope Decisions klar von Annahmen und Nicht-Zielen.
- Formuliere Validation-Kriterien so, dass sie durch Tests, Formatierung,
  Diff-Prüfung oder einen reproduzierbaren manuellen Check belegbar sind.
- Behalte die Sprache der bestehenden Projektdokumentation bei.
- Erzeuge keine zusätzlichen Planungsdateien außerhalb des angeforderten
  Spezifikationsordners.
