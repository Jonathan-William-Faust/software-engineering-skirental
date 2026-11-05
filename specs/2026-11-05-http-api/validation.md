# Validation: HTTP-API

Die Phase gilt als erfolgreich, wenn:

1. Die beschriebenen API-Ressourcen geroutet werden und JSON-Antworten liefern.
2. Gültige Ski-, Personen-, Ausleih- und Reservierungsabläufe über HTTP
   ausführbar sind.
3. Dashboard-Zähler die Store-Zustände korrekt widerspiegeln.
4. Fehlende Pflichtfelder, ungültige JSON-Daten, unbekannte Objekte und
   ungültige Domain-Übergänge explizite 4xx-Antworten liefern.
5. Personen nur mit `image/png`-Ausweisfoto registriert werden können.
6. `go test ./...`, `go test -race ./...` und `go vet ./...` erfolgreich
   durchlaufen und aussagekräftige Handler-Tests enthalten.
7. `go fmt ./...` und `git diff --check` ohne Fehler durchlaufen und keine
   neuen Dependencies oder Frontend-Dateien eingeführt werden.
8. README, Mission und Techstack weiterhin konsistent sind.

Der Merge nach `main` ist nur nach vollständigem Nachweis aller Kriterien als
Fast-Forward-Merge freigegeben.
