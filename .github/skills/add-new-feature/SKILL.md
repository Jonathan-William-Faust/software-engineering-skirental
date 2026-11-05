---
name: add-new-feature
description: "Automatisiert einen vollständigen Feature-Workflow: Feature-Spezifikation aus Roadmap und Projektkontext erstellen, den nächsten Feature-Branch implementieren, anhand von validation.md prüfen und bei vollständig bestandenen Tests nach main mergen."
argument-hint: "[optional: Featurename oder Roadmap-Phase]"
user-invocable: true
---

# Add New Feature

## Zweck

Dieser Skill führt den vollständigen, nachvollziehbaren Workflow für ein neues
Feature aus:

1. Spezifikation mit dem Workflow von `writing-feature-spec` erstellen.
2. Das spezifizierte Feature im Feature-Branch implementieren.
3. Die Kriterien aus `validation.md` reproduzierbar prüfen.
4. Nach jedem Hauptschritt einen eigenen Commit erzeugen und erst nach
   vollständig bestandener Validierung nach `main` mergen.

Der Skill ist für dieses Repository gedacht und hält die Sprache und Struktur
der vorhandenen Spezifikationen ein.

## Voraussetzungen und Sicherheitsregeln

- Lies vor jeder Änderung `README.md`, `specs/mission.md`,
  `specs/techstack.md`, `specs/roadmap.md` und die zuletzt abgeschlossene
  Spezifikation unter `specs/`.
- Prüfe zuerst den Git-Status. Bei uncommitteten Änderungen, die nicht
  eindeutig zu diesem Feature gehören, stoppe und frage nach; fremde Änderungen
  dürfen nicht überschrieben, zurückgesetzt oder mitcommittet werden.
- Arbeite nie direkt auf `main`. Erzeuge vor der ersten Änderung den Branch
  `feature/YYYY-MM-DD-taskname` aus dem aktuellen `main`-Stand.
- Verwende keine destruktiven Git-Befehle wie `reset --hard` oder
  `checkout --`.
- Committe und merge nur, weil dieser Composite-Skill das ausdrücklich
  verlangt. Ein Merge darf ausschließlich als `--ff-only` erfolgen.
- Bei einem fehlenden Artefakt, einem fehlgeschlagenen Test oder einer
  unerfüllten Validierungsbedingung darf kein Merge stattfinden. Behebe die
  Ursache und validiere erneut; wenn sie nicht sicher lösbar ist, stoppe mit
  einem klaren Fehlerbericht.

## Vorgehen

### 1. Kontext und Branch vorbereiten

1. Lies die oben genannten Projektdateien und prüfe, welche Roadmap-Phase als
   Nächstes offen ist.
2. Prüfe, dass `main` existiert, und stelle fest, ob der aktuelle Branch
   bereits ein passender Feature-Branch ist. Falls nicht, wechsle auf `main`,
   aktualisiere nur mit einer nicht-destruktiven Operation, und erzeuge einen
   neuen Branch mit dem Muster `feature/YYYY-MM-DD-taskname`.
3. Halte den Ausgangszustand mit `git status --short --branch` fest.

### 2. Feature-Spezifikation erstellen

Führe die Arbeitsschritte des Skills
[`writing-feature-spec`](../writing-feature-spec/SKILL.md) vollständig aus:

1. Bestimme den nächsten offenen Roadmap-Task. Bei mehreren gleichrangigen
   Tasks oder unklarem Scope frage nach.
2. Verwende **vor jeder Dateiänderung** eine einzige gruppierte
   `ask_user`-Runde mit den Abschnitten `plan.md`, `requirements.md` und
   `validation.md`.
3. Erstelle `specs/YYYY-MM-DD-taskname/` mit genau den Kernartefakten
   `plan.md`, `requirements.md` und `validation.md`.
4. Verifiziere, dass Scope, Nicht-Ziele, messbare Kriterien und Merge-Regeln
   mit Mission, Techstack und Roadmap übereinstimmen.
5. Prüfe `git diff --check` und committe ausschließlich die
   Spezifikationsartefakte mit einer imperativen Nachricht, zum Beispiel
   `Write <feature> specification`. Der Commit muss den vorgeschriebenen
   Copilot-Trailer enthalten. Ohne erfolgreichen Spezifikations-Commit darf
   die Implementierung nicht beginnen.

### 3. Implementierung

1. Lies `plan.md`, `requirements.md` und `validation.md` erneut und leite
   daraus konkrete Arbeitspakete ab.
2. Implementiere ausschließlich den spezifizierten Scope. Bestehende
   Konventionen, Schichten, Fehlerbehandlung und Tests sind zu verwenden.
3. Schreibe zuerst oder zusammen mit der Implementierung aussagekräftige Tests
   für Erfolgs- und Fehlerfälle. Tests dürfen keine späteren Store-, HTTP- oder
   Frontend-Schichten vorwegnehmen, wenn diese nicht spezifiziert sind.
4. Aktualisiere direkt betroffene Dokumentation und die Spezifikationsartefakte,
   insbesondere einen vorhandenen Batch-Report oder Checkpoint.
5. Führe die kleinsten passenden Tests während der Implementierung aus und
   behebe Fehler vollständig. Keine stillen Fallbacks oder verschluckten
   Fehler.
6. Prüfe den Implementierungsstand mit den passenden Tests, `go fmt ./...`
   beziehungsweise dem projektspezifischen Formatter und `git diff --check`.
   Committe anschließend ausschließlich die Implementierung mit einer
   imperativen Nachricht, zum Beispiel `Implement <feature>`. Dieser Commit
   muss den vorgeschriebenen Copilot-Trailer enthalten. Bei einem
   fehlgeschlagenen Test wird nicht committen; zuerst Ursache beheben und
   erneut prüfen.

### 4. Validierung gegen `validation.md`

Prüfe jedes Kriterium einzeln und dokumentiere den Nachweis, bevor du
entscheidest. Mindestens ausführen:

```text
go test ./...
go fmt ./...
go vet ./...
git diff --check
git diff main...HEAD --stat
git status --short
```

Nach `go fmt ./...` muss erneut geprüft werden, dass keine unbeabsichtigten
Formatänderungen offen sind. Verwende die konkreten Befehle und manuellen
Nachweise aus `validation.md`; ein allgemeiner Testlauf ersetzt kein
fachliches Akzeptanzkriterium. Prüfe außerdem:

- alle geforderten Modelle, Methoden und Fehlerfälle,
- keine neuen Dependencies außerhalb des spezifizierten Scopes,
- keine unerlaubten HTTP-, Store- oder Frontend-Änderungen,
- Konsistenz mit README, Mission und Techstack,
- einen sauberen Arbeitsbaum vor dem Commit.

Wenn ein Kriterium fehlschlägt, führe keinen Validierungs-Commit und keinen
Merge durch. Dokumentiere die Abweichung, behebe die Ursache in einem
separaten Korrektur-Commit und starte danach die komplette Validierung erneut.

### 5. Commit und Merge

Nur wenn jeder Hauptschritt seinen eigenen Commit besitzt und jedes Kriterium
aus `validation.md` nachweislich erfüllt ist:

1. Prüfe `git diff` und `git status` auf Vollständigkeit und fremde Änderungen.
2. Dokumentiere die Validierungsergebnisse in `batch-report.yaml` oder dem
   vorgesehenen Validierungsartefakt und committe ausschließlich diese
   Nachweise mit einer imperativen Nachricht, zum Beispiel
   `Validate <feature>`. Verwende den vorgeschriebenen Trailer:

   ```text
   Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>
   ```

3. Verifiziere, dass die drei Commits in der Reihenfolge Spezifikation,
   Implementierung und Validierung vorhanden sind.
4. Wechsle auf `main` und führe ausschließlich
   `git merge --ff-only feature/YYYY-MM-DD-taskname` aus.
5. Führe auf dem resultierenden `main` erneut `go test ./...` und die für
   `validation.md` relevanten Prüfungen aus.
6. Melde Branch, die drei Zwischen-Commits, Merge-Status, Testausgaben und verbleibende
   Änderungen. Bei einem nicht möglichen Fast-Forward-Merge stoppe statt einen
   automatischen Merge-Commit zu erzeugen.

## Abschlussbericht

Der Abschlussbericht enthält:

- den Pfad der drei Spezifikationsdateien,
- den Feature-Branch und den Commit,
- eine Tabelle aller Validierungskriterien mit Nachweis und Status,
- die ausgeführten Test-, Format- und Diffbefehle,
- den finalen `main`-Stand und Merge-Ergebnis,
- etwaige Warnungen oder bewusst nicht erfüllte Kriterien.
