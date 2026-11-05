# Validation: Standard-Skikatalog

Die Dokumentation und Implementierung gelten als nachvollziehbar, wenn:

1. Das Domain-Modell und der Store Ski-Größen abbilden und bestehende
   Erzeugungsaufrufe ohne Größe weiterhin funktionieren.
2. Der produktive Serverstart exakt 100 Ski bereitstellt.
3. Der Katalog exakt zehn unterschiedliche Modelle und je Modell zehn Größen
   zwischen 150 und 195 cm enthält.
4. `GET /api/skis` die Größen liefert und die Ski-Verwaltung sie sichtbar
   darstellt.
5. Der leere Testserver weiterhin unabhängig vom Startkatalog erzeugt werden
   kann.
6. `go fmt ./...`, `go test ./...`, `go test -race ./...`, `go vet ./...` und
   `git diff --check` erfolgreich durchlaufen.
7. Keine neuen Dependencies oder persistente Speicherung eingeführt wurden.

Nachweise aus diesem Arbeitsschritt:

- Katalogtest für 100 Ski, zehn Modelle und zehn Größen je Modell: bestanden.
- Seeded-Server-Test mit 100 Ski: bestanden.
- Fokussierte Race-Tests für Domain, Handler, Store und Seeded Server:
  bestanden.
- Formatierung, Vet und Diff-Prüfung: bestanden.

Die Änderungen sind für die weitere Verwendung bereit, sobald die üblichen
Projekt-Commitregeln angewendet wurden.
