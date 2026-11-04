# Plan: Projektgerüst und lokales Git

## 1. Go-Modul und Grundkonfiguration

- Go-Modul mit einem noch festzulegenden Modulnamen anlegen.
- Eine unterstützte Go-Version festlegen und in der Modulkonfiguration
  dokumentieren.
- Keine zusätzlichen Dependencies einführen.

## 2. Verzeichnisstruktur

- Den Einstiegspunkt als `main.go` oder unter `cmd/server/` anlegen.
- Die vorgesehenen Verzeichnisse `internal/domain/`, `internal/store/` und
  `internal/handler/` vorbereiten.
- Nur notwendige Platzhalter oder minimale Go-Dateien anlegen; keine
  generierten Dateien manuell pflegen.

## 3. Minimalen Server bereitstellen

- Einen minimalen HTTP-Server als lauffähigen Einstiegspunkt implementieren.
- Einen einfachen Health- oder Start-Endpunkt bereitstellen, damit der
  erfolgreiche Start überprüfbar ist.
- Den Port nicht in dieser Spezifikation vorwegnehmen, sondern in der
  Implementierung mit einer klaren lokalen Konfiguration festlegen.

## 4. Entwicklungsbefehle verifizieren

- `go run .` ausführen und den Serverstart überprüfen.
- `go test ./...` ausführen, auch wenn zunächst nur wenige Tests vorhanden sind.
- `gofmt -w .` beziehungsweise `go fmt ./...` ausführen und einen sauberen
  Formatierungsstand sicherstellen.

## 5. Lokalen Git-Stand herstellen

- Die neuen Projektdateien mit einem kurzen, imperativen Commit versionieren.
- Vor dem Merge prüfen, dass nur zum Projektgerüst gehörende Änderungen auf dem
  Branch enthalten sind.
- Den Branch für den Merge in `main` vorbereiten; ein GitHub-Repository ist in
  dieser Phase kein Ziel.
