# Validation

1. `go fmt ./...` läuft erfolgreich.
2. `go test ./...` läuft erfolgreich.
3. `go test -race ./...` läuft erfolgreich.
4. `go vet ./...` und `git diff --check` laufen erfolgreich.
5. Ein API-/Integrationstest bestätigt, dass die Rückgabe den aktiven Loan
   beendet und den Ski wieder als verfügbar ausliefert.
6. Eine Frontend-Prüfung bestätigt:
   - aktive Leihen werden im Personen-Tab als Rückgabeaktionen angezeigt;
   - der Button ruft `DELETE /api/loans/{skiID}` auf;
   - nach einer Rückgabe wird `refresh()` ausgeführt;
   - die Aktion verschwindet nach erfolgreicher Rückgabe aus der
     Personenansicht.
7. Spezifikation, Implementierung und Validierung werden separat committed.

## Durchgeführte Nachweise

- `go fmt ./...` und `go test ./...` erfolgreich.
- `go test -race -count=1 ./...` erfolgreich.
- `go vet ./...` und `git diff --check` erfolgreich.
- Der Integrationstest `TestReturningLoanMakesSkiAvailableForAllViews`
  bestätigt, dass die Backend-Rückgabe den Ski auf `verfügbar` setzt und den
  Loan mit gesetztem `EndedAt` zurückliefert.
- Die statische Frontend-Prüfung bestätigt, dass aktive Loans im Personen-Tab
  mit `Ausleihe beenden` gerendert werden, dieselbe
  `DELETE /api/loans/{skiID}`-Aktion nutzen und danach `refresh()` ausführen.
- Damit liegt die Ursache der ursprünglichen fehlenden Anzeige im Frontend,
  nicht in der internen Backend-Verwaltung.
