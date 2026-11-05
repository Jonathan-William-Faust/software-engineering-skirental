# Requirements: Ski-ID-Anzeige

## Scope Decisions

Die Domain erzeugt Ski-IDs bereits serverseitig. Die Ursache der sichtbaren
`undefined`-Werte lag in der JavaScript-Feldauflösung: Go serialisiert die
exportierte Struktur ohne JSON-Tags als `ID`, während das Frontend zuerst
`id` und danach nur `Id` suchte.

Die Feldauflösung wird deshalb um Akronyme wie `ID` ergänzt. Dadurch werden
Startkatalog-Ski und per API angelegte Ski mit ihrer eindeutigen ID angezeigt,
ohne die bestehende API oder Domain-Kompatibilität unnötig zu verändern.

## Nicht-Ziele

Nicht enthalten sind neue ID-Formate, clientseitige ID-Erzeugung,
Persistenzänderungen oder Änderungen an Ausleih- und Reservierungsregeln.

Die serverseitigen Regeln aus [mission.md](../mission.md) bleiben verbindlich:
Ski-IDs verwenden das Format `S-` plus vier Ziffern.
