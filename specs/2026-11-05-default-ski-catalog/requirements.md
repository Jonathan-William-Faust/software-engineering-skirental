# Requirements: Standard-Skikatalog

## Scope Decisions

Diese Spezifikation dokumentiert die Änderungen seit dem letzten Commit
`e8f3045` im laufenden Skiverleih-Prototyp.

Beim normalen App-Start werden automatisch 100 Ski im In-Memory-Store
angelegt. Der Katalog besteht aus zehn real verfügbaren Modellen:

- Atomic Redster S9
- Rossignol Hero Elite ST Ti
- Head Supershape e-Magnum
- Salomon S/Max 10
- Fischer RC4 The Curv
- Völkl Racetiger SL
- Blizzard Firebird SRC
- Nordica Dobermann Spitfire 76 RB
- Elan Amphibio 14 Ti
- K2 Disruption 76

Jedes Modell wird in den Größen 150, 155, 160, 165, 170, 175, 180, 185,
190 und 195 cm angelegt. Die Daten sind beim Neustart des nicht-persistenten
In-Memory-Stores erneut vorhanden.

Das Größenfeld wird durch Domain, Store, HTTP-API und Frontend weitergereicht
und in der Ski-Verwaltung angezeigt. Bestehende Aufrufer ohne Größenangabe
bleiben kompatibel.

## Nicht-Ziele

Nicht enthalten sind persistente Speicherung, Bestandsimport, externe
Produktkataloge, Preisverwaltung, neue Dependencies und Änderungen an
Personen-, Ausleih- oder Reservierungsregeln.

Die Umsetzung bleibt mit [mission.md](../mission.md),
[techstack.md](../techstack.md) und [roadmap.md](../roadmap.md) konsistent.
