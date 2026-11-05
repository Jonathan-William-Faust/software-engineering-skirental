# Plan

1. **Startdaten erweitern**
   - Den seeded In-Memory-Store um Max Mustermann mit Geburtsdatum
     `01.01.2000` erweitern.
   - Eine gültige PNG-Ausweisfoto-Metadatenrepräsentation verwenden und die
     erwartete `P-0001`-ID durch einen Test absichern.

2. **Ski-Aktionen vervollständigen**
   - Die Spalte `Aktion` in der Ski-Verwaltung für verliehene Ski mit dem
     README-Button `Leihe beenden` verbinden.
   - Die vorhandene Rückgabe-Aktion über die bestehende DELETE-API ausführen
     und die Tabellen danach aktualisieren.

3. **Personen-Aktionen ergänzen**
   - Für jede Person ohne aktive Reservierung den Button `Neue Ausleihe`
     anzeigen.
   - Für Personen mit aktiver Reservierung die Buttons `Miete reservierter
     Ski beginnen`, `Reservierung bearbeiten` und `Reservierung löschen`
     anzeigen.
   - Jede Aktion über ein Bestätigungs- oder Eingabe-Pop-up ausführen und
     anschließend den Zustand aktualisieren.

4. **Dokumentation und Nachweise**
   - README um die Standardperson und die konkreten Personen-Aktionen ergänzen.
   - Backend-, Store- und Frontend-nahe Tests beziehungsweise reproduzierbare
     Prüfungen ergänzen.
