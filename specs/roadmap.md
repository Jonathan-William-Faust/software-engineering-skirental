# Roadmap

Die Umsetzung erfolgt in kleinen, überprüfbaren Phasen. Jede Phase endet mit
einem lauffähigen Zwischenstand; unnötige Erweiterungen werden zugunsten eines
klaren Prototyps zurückgestellt.

## Phase 1 – Projektgerüst und lokales Git

- Go-Modul und minimale Verzeichnisstruktur anlegen.
- Lokales Git-Repository initialisieren und einen ersten funktionierenden Stand
  versionieren.
- Start- und Testbefehle aus `template.txt` ausführbar machen.

## Phase 2 – Domain-Modell

- Structs für Ski, Person, Ausleihe und Reservierung definieren.
- ID-Formate und Statuswerte zentral festlegen.
- Geschäftsregeln für Registrierung, Ausleihe, Rückgabe und Reservierung
  implementieren.
- Fehlerfälle explizit modellieren und mit Unit-Tests absichern.

## Phase 3 – Thread-sicherer In-Memory-Store

- Store-Schnittstelle und konkrete In-Memory-Implementierung anlegen.
- Lesen, Anlegen und Statusänderungen kapseln.
- Mutex-Schutz für konkurrierende Zugriffe ergänzen.
- Store-Tests für leere Daten, gültige Abläufe und ungültige Übergänge schreiben.

## Phase 4 – HTTP-API

- HTTP-Server und Routing aufbauen.
- Endpoints für Dashboard-Daten, Ski, Personen, Ausleihen und Reservierungen
  ergänzen.
- Eingaben validieren, PNG-Uploads prüfen und Fehler verständlich an Clients
  zurückgeben.
- Handler- und API-Tests für die Kernabläufe hinzufügen.

## Phase 5 – HTML/JavaScript-Frontend

- Navigation und vier geforderte Ansichten umsetzen.
- Tabellen, Kennzahlen und Aktionsbuttons an die API anbinden.
- Pop-ups für Eingaben und Bestätigungen implementieren.
- Sichtbare Zustände und Aktionen für verfügbare, reservierte und verliehene Ski
  konsistent darstellen.

## Phase 6 – Integration und fachliche Durchläufe

- Frontend, HTTP-API und Store gemeinsam starten.
- Registrierung mit PNG, Ausleihe, Rückgabe sowie alle Reservierungsübergänge
  manuell und automatisiert durchspielen.
- Konsistenz der Dashboard-Kennzahlen nach jeder Zustandsänderung prüfen.

## Phase 7 – Qualität und Abschluss

- `gofmt` und `go test ./...` ausführen; vorhandenen Linter verwenden.
- Fehler- und Randfälle nachschärfen, ohne den Prototyp unnötig zu vergrößern.
- README um Start- und Bedienhinweise ergänzen.
- Einen nachvollziehbaren lokalen Git-Verlauf mit kurzen, imperativen
  Commit-Nachrichten hinterlassen.
