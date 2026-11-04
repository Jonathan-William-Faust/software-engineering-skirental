# Requirements: Projektgerüst und lokales Git

## Scope

Diese Phase schafft den kleinsten lauffähigen Rahmen für die spätere
Skiverleih-Webanwendung. Sie umfasst Modulkonfiguration, Verzeichnisstruktur,
einen minimalen HTTP-Einstiegspunkt sowie die lokale Git-Versionierung.

Nicht Bestandteil dieser Phase sind Domain-Objekte, In-Memory-Geschäftslogik,
fachliche API-Endpunkte, Frontend-Seiten, Datei-Uploads und persistente
Datenhaltung. Diese Themen folgen in späteren Roadmap-Phasen.

## Entscheidungen

- Das Backend wird in Go umgesetzt.
- Das Frontend bleibt in dieser Phase unimplementiert; die spätere Variante ist
  reines HTML und JavaScript gemäß [techstack.md](../techstack.md).
- Die Struktur orientiert sich an `main.go` oder `cmd/server/` sowie
  `internal/domain/`, `internal/store/` und `internal/handler/`.
- Es werden keine zusätzlichen Dependencies eingeführt.
- Die Speicherung bleibt gemäß [mission.md](../mission.md) und
  [techstack.md](../techstack.md) zunächst außerhalb dieser Phase.
- Go-Version und Modulname werden bei der Implementierung festgelegt und müssen
  anschließend nachvollziehbar dokumentiert sein.
- Der lokale Server-Port wird bei der Implementierung festgelegt; diese
  Spezifikation schreibt keinen konkreten Port vor.
- Versionierung erfolgt ausschließlich lokal mit Git. Ein GitHub-Repository ist
  für diese Phase nicht erforderlich.

## Kontext und Leitplanken

Das Projekt ist ein verständlicher, leicht änderbarer Prototyp. Der Einstieg
und die Struktur müssen deshalb klein bleiben und dürfen keine Framework- oder
Dependency-Komplexität einführen. Fehler dürfen nicht still verschluckt werden;
der Server muss einen überprüfbaren Startzustand liefern.

Die Phase muss spätere Erweiterungen für Domain, Store, Handler und Frontend
ermöglichen, ohne die fachlichen Regeln aus Mission oder Techstack zu umgehen.
