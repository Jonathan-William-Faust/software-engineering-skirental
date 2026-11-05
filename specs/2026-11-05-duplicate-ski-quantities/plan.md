# Plan: Doppelte Ski, Mengen und Sammelreservierungen

## Ziel

Der Startkatalog enthält jedes Modell in jeder Größe zweimal. Ausleihen und
Reservierungen können die gewünschte Anzahl je Modell-/Größenkombination über
Checkbox-/Mengenauswahl erfassen. Im Personen-Tab werden mehrere aktive
Reservierungen einer Person als gemeinsame Aktionen dargestellt.

## Umsetzung

1. Startkatalog und Dokumentation auf zwei konkrete Ski je Modell und Größe
   erweitern, ohne die eindeutigen Ski-IDs zu verändern.
2. Ski-Auswahldialoge um Mengenfelder je Modell und Größe erweitern und nur
   verfügbare konkrete Ski in der gewünschten Anzahl auswählen.
3. Reservierungsaktionen im Personen-Tab gruppieren: eine Sammelumwandlung,
   Sammelbearbeitung und Sammellöschung je Person.
4. Die Sammelbearbeitung mit Checkboxen für bestehende Reservierungen und neue
   verfügbare Ski implementieren; bestehende Zeiträume bleiben erhalten.
5. Domain-/Store-/Frontend-Regressionstests und Validierungsnachweise ergänzen.

## Abgrenzung

- Ski gleicher Modell-/Größenkombination bleiben eigene konkrete Ski mit eigener
  ID.
- Die Mengenangabe wird nicht als neues Feld an Ausleihe oder Reservierung
  gespeichert; die konkrete Auswahl bleibt je Ski-ID nachvollziehbar.
- Bei Sammelbearbeitung behalten bestehende Reservierungen ihre Zeiträume.
