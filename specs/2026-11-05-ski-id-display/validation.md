# Validation: Ski-ID-Anzeige

Die Korrektur gilt als erfolgreich, wenn:

1. Das Frontend `ID`-Felder aus den bestehenden Go-JSON-Antworten auflöst und
   keine `undefined`-Ski-IDs rendert.
2. Alle 100 Startkatalog-Ski eine nicht-leere, eindeutige ID besitzen.
3. Ein neu angelegter Ski automatisch eine eindeutige ID `S-0000` erhält.
4. `go test ./...`, fokussierte Race-Tests, `go vet ./...` und
   `git diff --check` erfolgreich durchlaufen.
5. Keine clientseitige ID-Erzeugung oder neue Dependency eingeführt wird.

Nachweise:

- Katalogtest prüft 100 nicht-leere und eindeutige Ski-IDs.
- Store-Test prüft `S-0001` für den ersten benutzerseitig angelegten Ski.
- Frontend verwendet die korrigierte Auflösung für `id`, `Id` und `ID`.
