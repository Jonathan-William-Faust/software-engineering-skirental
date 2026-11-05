# Requirements

## Scope Decisions

- Eine Person erhält zusätzlich:
  - Land
  - PLZ mit vier oder fünf Ziffern
  - Straße
  - Hausnummer
  - Handynummer
- Alle neuen Felder sind Pflichtfelder.
- Die Personenregistrierung verwendet wieder eine Dateiauswahl für das
  Ausweisfoto.
- Nur PNG-Dateien werden akzeptiert. Eine fehlende oder ungültige Datei zeigt
  ein Pop-up; anschließend wird die Fotoauswahl wiederholt.
- Der Foto-Upload sendet die PNG-Datei als Data-URL und den MIME-Typ an die
  bestehende JSON-API.
- Die Personenauswahl für Ausleihe und Reservierung zeigt Name, Geburtsdatum,
  vollständige Adresse und Personen-ID.
- Die Auswahl bleibt auf genau eine Person begrenzt.

## Fachlicher Kontext

- Personenmodell und Validierung liegen in
  [`internal/domain/models.go`](../../internal/domain/models.go).
- API- und JSON-Übersetzung liegen in
  [`internal/handler/api.go`](../../internal/handler/api.go).
- Registrierung und Personenauswahl liegen in
  [`web/app.js`](../../web/app.js).

## Nicht-Ziele

- Keine Änderung der Ausleih- oder Reservierungsregeln.
- Keine persistente Speicherung.
- Keine Unterstützung anderer Bildformate.
- Keine externe Upload-Abhängigkeit.
