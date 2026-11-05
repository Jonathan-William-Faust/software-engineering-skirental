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
