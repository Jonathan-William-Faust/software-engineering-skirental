# Roadmap

Die Umsetzung erfolgt in kleinen, überprüfbaren Phasen. Jede Phase endet mit
einem lauffähigen Zwischenstand; unnötige Erweiterungen werden zugunsten eines
klaren Prototyps zurückgestellt.

Die Phasen 1 bis 6 sind umgesetzt. Die folgenden Punkte beschreiben den
aktuellen Stand und die verbleibenden Qualitätsschritte.

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
- Flexible inklusive Miet- und Reservierungszeiträume sowie Preisberechnung
  unterstützen.
- Zeitraumbezogene Verfügbarkeit prüfen und Überschneidungen aktiver
  Ausleihen oder Reservierungen ausschließen.
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
- Den Endpunkt für Ski bereitstellen, die über einen vollständigen Zeitraum
  verfügbar sind.
- Eingaben validieren, PNG- und JPEG-Uploads prüfen und Fehler verständlich an Clients
  zurückgeben.
- Handler- und API-Tests für die Kernabläufe hinzufügen.

## Phase 5 – HTML/JavaScript-Frontend

- Navigation und vier geforderte Ansichten umsetzen.
- Tabellen, Kennzahlen und Aktionsbuttons an die API anbinden.
- Pop-ups für Eingaben und Bestätigungen implementieren.
- Sichtbare Zustände und Aktionen für verfügbare, reservierte und verliehene Ski
  konsistent darstellen.
- Zeitraum zuerst abfragen und danach ausschließlich zeitraumbezogen verfügbare
  Ski zur Auswahl anbieten.

## Phase 6 – Integration und fachliche Durchläufe

- Frontend, HTTP-API und Store gemeinsam starten.
- Registrierung mit PNG, Ausleihe, Rückgabe sowie alle Reservierungsübergänge
  manuell und automatisiert durchspielen.
- Konsistenz der Dashboard-Kennzahlen nach jeder Zustandsänderung prüfen.
- Prüfen, dass Reservierungen erst am Reservierungsbeginn in Ausleihen
  umgewandelt werden können.

## Phase 7 – Qualität und Abschluss

- `go fmt ./...`, `go test ./...`, `go test -race ./...`, `go vet ./...` und
  `git diff --check` ausführen; vorhandenen Linter verwenden.
- Fehler- und Randfälle nachschärfen, ohne den Prototyp unnötig zu vergrößern.
- README um Start- und Bedienhinweise ergänzen.
- Einen nachvollziehbaren lokalen Git-Verlauf mit kurzen, imperativen
  Commit-Nachrichten hinterlassen.
