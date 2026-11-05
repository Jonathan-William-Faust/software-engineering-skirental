# Plan

1. **Spezifikation**
   - Rückgabe-Pop-up und Sortierreihenfolgen dokumentieren.
   - Bestehende API- und Frontend-Leitplanken referenzieren.

2. **Implementierung**
   - Pro Person genau einen Rückgabe-Button anzeigen, sobald mindestens eine
     aktive Ausleihe besteht.
   - Im Rückgabe-Pop-up die aktiven Ski dieser Person mit Checkboxen auflisten.
   - Für jede ausgewählte Ausleihe die bestehende Rückgabe-API aufrufen und
     danach die Ansichten aktualisieren.
   - Ski- und Personenlisten mit den geforderten Vergleichsschlüsseln sortieren.
   - README um das neue Verhalten ergänzen.

3. **Validierung**
   - Go-Formatierung, Tests, Race-Test, Vet und Diff-Prüfung ausführen.
   - Die Sortiervergleiche, Checkbox-Auswahl und Rückgabeaufrufe statisch
     prüfen.
   - Validierungsnachweise dokumentieren und separat committen.
