# Plan: Standard-Skikatalog

## 1. Größenmodell erweitern

- Das Domain-Modell um eine optionale Ski-Größe ergänzen.
- Bestehende Ski-Erzeugung ohne Größenangabe kompatibel halten.
- Größenwerte validieren und über Domain sowie Store kapseln.

## 2. Startkatalog seed-en

- Beim produktiven App-Start einen deterministischen In-Memory-Katalog
  erzeugen.
- Zehn real verfügbare Skimodelle mit jeweils zehn Größen von 150 bis
  195 cm anlegen.
- Testserver weiterhin ohne vorbefüllte Daten erzeugen können.

## 3. Darstellung und Nachweise

- Die Größe in API-Antworten und Ski-Verwaltung anzeigen.
- Tests für exakt 100 Ski, zehn Modelle und zehn Größen je Modell ergänzen.
- README um das Verhalten des Startkatalogs erweitern.
