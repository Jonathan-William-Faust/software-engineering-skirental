# Plan

1. **Spezifikation**
   - Mehrfachreservierungen pro Person, Checkbox-Auswahl und Altersprüfung
     dokumentieren.

2. **Implementierung**
   - Dashboard um `Reservierung anlegen` ergänzen.
   - Eine gemeinsame Checkbox-Auswahl verfügbarer Ski für Reservierungen und
     Ausleihen verwenden.
   - Für jede ausgewählte Kombination eine Reservierung anlegen.
   - Mehrere aktive Reservierungen pro Person im Domain-Service erlauben und
     die Personenansicht für mehrere Reservierungen darstellen.
   - Geburtsdaten in der Zukunft sowie Personen unter 18 Jahren im Backend und
     im Frontend ablehnen.
   - README und API-Tests aktualisieren.

3. **Validierung**
   - Formatierung, Tests, Race-Tests, Vet und Diff-Prüfung ausführen.
   - Checkbox-, Dashboard- und Alterslogik statisch prüfen.
   - Validierungsnachweise dokumentieren und separat committen.
