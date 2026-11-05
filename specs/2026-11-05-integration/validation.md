# Validation: Integration und fachliche Durchläufe

Die Phase gilt als erfolgreich, wenn:

1. `/`, `/healthz` und die API über den gemeinsamen Server erreichbar sind.
2. Registrierung, Ski-Anlage, Ausleihe und Rückgabe als HTTP-Ablauf bestehen.
3. Reservierung anlegen, bearbeiten, umwandeln und verwerfen als HTTP-Abläufe
   bestehen.
4. Dashboard-Zähler nach jeder Zustandsänderung konsistent sind.
5. `go test ./...`, `go test -race ./...`, `go vet ./...` und `git diff --check`
   erfolgreich laufen.

Der Merge nach `main` ist nur nach vollständigem Nachweis als
Fast-Forward-Merge freigegeben.
