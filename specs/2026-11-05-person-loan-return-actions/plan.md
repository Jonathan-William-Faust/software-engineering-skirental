# Plan

1. **Ursache und bestehendes Verhalten absichern**
   - Prüfen, dass `DELETE /api/loans/{skiID}` den Store und den Ski-Status
     korrekt aktualisiert.
   - Prüfen, dass die Ski-Verwaltung nach erfolgreicher Rückgabe bereits
     `refresh()` ausführt.

2. **Personenansicht erweitern**
   - Aktive Leihen je Person ermitteln.
   - Für jede aktive Leihe den Button `Ausleihe beenden` in der Personenzeile
     anzeigen.
   - Den Button an dieselbe Rückgabeaktion wie in der Ski-Verwaltung anbinden.
   - Nach erfolgreicher Rückgabe alle geladenen Tabellen aktualisieren.

3. **Dokumentation und Validierung**
   - README und diese Spezifikation um die Personenaktion ergänzen.
   - Backend-, Race-, Format- und Diff-Prüfungen ausführen.
   - Die Ursache und das bestätigte Refresh-Verhalten im Validierungsnachweis
     festhalten.
