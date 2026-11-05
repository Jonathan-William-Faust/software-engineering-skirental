# Plan: Reservierung, Mietzeitraum, Preise und Upload-Dialog

## Ziel

Die bestehende Skiverleih-App soll Ausweisdateien sichtbar über ein Dialog-
Pop-up auswählen lassen und Reservierungen mit einem frei wählbaren zukünftigen
Mietzeitraum unterstützen. Ausleihen bleiben Sofortvorgänge, erhalten aber eine
flexible Dauer. Für jeden Ski wird ein Tagespreis zwischen 10,00 € und 20,00 €
verwaltet und der Zeitraumspreis angezeigt.

## Umsetzung

1. Domain-Modell um Miet-/Reservierungszeitraum, Tagespreis und Preisberechnung
   erweitern; Datums- und Überschneidungsregeln explizit validieren.
2. Store und HTTP-API so erweitern, dass Preise und Zeiträume vollständig
   gespeichert, übertragen und bei Reservierung/Ausleihe geprüft werden.
3. Frontend-Dialoge für Ausweisdatei, Reservierung und Ausleihe um sichtbare
   Zeitraum-/Preis-Eingaben und Ergebnisanzeigen ergänzen.
4. Bestehende Abläufe und Tests auf die neuen Pflichtfelder anpassen und
   zusätzliche Erfolgs- und Fehlerfalltests ergänzen.
5. README und Fachleitplanken an den tatsächlichen Funktionsumfang anpassen.

## Abgrenzung

- Ausleihen können nicht in der Zukunft gestartet werden.
- Reservierungen werden sofort als `reserviert` markiert, auch wenn ihr Beginn
  in der Zukunft liegt.
- Es gibt keine Zahlungen, Rechnungen, Rabatte oder persistente Speicherung.
- Der Preis wird beim Ski-Anlegen erfasst; der Startkatalog erhält eine feste
  Staffelung im Bereich 10,00 € bis 19,00 €.
