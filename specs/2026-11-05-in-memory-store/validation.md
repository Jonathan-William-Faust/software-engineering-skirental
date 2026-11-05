# Validation: Thread-sicherer In-Memory-Store

Die Phase gilt als erfolgreich, wenn alle folgenden Kriterien erfüllt sind:

1. `internal/store/` enthält eine Store-Schnittstelle und eine konkrete
   In-Memory-Implementierung für die vier Domain-Bereiche.
2. Der Store verwendet `sync.RWMutex`; konkurrierende Lese- und
   Schreibzugriffe verändern keine inkonsistenten oder fremden Zustände.
3. Leere Stores liefern leere, nicht-nil-fähige Sammlungen ohne Panic.
4. Gültige Erzeugungs-, Lese-, Ausleih-, Rückgabe- und
   Reservierungsabläufe funktionieren über den Store.
5. Unbekannte Objekte, doppelte Reservierungen, nicht verfügbare Ski und
   Rückgaben ohne aktive Ausleihe liefern explizite Fehler.
6. Store-Unit-Tests decken Erfolgs- und Fehlerfälle sowie parallele Zugriffe
   ab.
7. `go test ./...` und `go test -race ./...` laufen erfolgreich durch.
8. `go fmt ./...`, `go vet ./...` und `git diff --check` hinterlassen keine
   Fehler; es werden keine neuen Dependencies und keine HTTP-/Frontend-Code
   eingeführt.
9. Die Implementierung bleibt mit README, Mission und Techstack konsistent.

Der Merge nach `main` ist ausschließlich nach erfolgreicher Prüfung aller neun
Kriterien als Fast-Forward-Merge freigegeben.
