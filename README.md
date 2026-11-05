# Skiverleih-Prototyp

Das Projekt ist ein leicht verständlicher Prototyp für eine
Skiverleih-Webanwendung. Es enthält ein Go-Backend, einen thread-sicheren
In-Memory-Store und ein statisches HTML/JavaScript-Frontend.

## Technischer Stand

- Go-Modul: `skirental`
- Unterstützte Go-Version: `1.22`
- Lokaler Server-Port: `8080`
- Health-Endpunkt: `GET http://localhost:8080/healthz`
- Weboberfläche: `http://localhost:8080/`
- API: `/api/dashboard`, `/api/skis`, `/api/skis/available`,
  `/api/people`, `/api/loans` und `/api/reservations`

## Entwicklung

```text
go run .
go test ./...
go fmt ./...
```

Der Server meldet sich unter `http://localhost:8080/healthz` mit
`{"status":"ok"}`. Die Weboberfläche bietet Dashboard, Ski-Verwaltung,
Personen und Reservierungen. Aktionen werden über Dialoge eingegeben und
anschließend über die JSON-API ausgeführt.

Beim Hinzufügen eines Skis werden Marke, Modell, Länge in Zentimetern und ein
Tagespreis zwischen 10,00 € und 20,00 € abgefragt. Leere oder ungültige
Eingaben sind nicht zulässig; ein Abbruch beendet den gesamten Vorgang.

Bei der Personenregistrierung werden Land und Stadt sowie die übrigen
Adressdaten erfasst. Die Adresse wird im Personen-Tab einschließlich Stadt
angezeigt. PNG- und JPEG-Ausweisfotos werden akzeptiert.
Ungültige oder fehlende Dateien werden als Pop-up gemeldet und die Fotoauswahl
wird wiederholt. Zusätzlich werden Land, PLZ, Straße, Hausnummer und
Handynummer erfasst. Personen dürfen höchstens 100 Jahre alt sein.

Bei der Personenregistrierung wird das Geburtsdatum im Format `DD.MM.YYYY`
abgefragt. Ungültige, zukünftige oder zu junge Datumsangaben werden als Pop-up
gemeldet und können direkt erneut eingegeben werden. Das Mindestalter beträgt
18 Jahre.

Beim Start werden automatisch 200 Ski angelegt: zehn real verfügbare Modelle
mit jeweils zehn Größen von 150 bis 195 cm und zwei konkreten Ski je
Modell/Größenkombination. Diese Startdaten stehen direkt in der Ski-Verwaltung
zur Verfügung und werden bei jedem Neustart des In-Memory-Stores neu erzeugt.

Zusätzlich wird die Standardperson Max Mustermann mit dem Geburtsdatum
`01.01.2000` angelegt. In der Ski-Verwaltung erscheint für verliehene Ski in
der Spalte „Aktion“ der Button „Leihe beenden“. Im Personen-Tab stehen je nach
Zustand die Buttons „Neue Ausleihe“ oder „Miete reservierter Ski beginnen“,
„Reservierung bearbeiten“ und „Reservierung löschen“ zur Verfügung. Jede Aktion
öffnet ein Bestätigungs- oder Eingabe-Pop-up.

Beim Start einer Ausleihe werden alle verfügbaren Ski gruppiert nach Modell und
Größe in einem Auswahl-Pop-up angezeigt. Je Kombination kann die gewünschte
Menge angegeben werden; es können ein oder mehrere konkrete Ski ausgewählt werden;
für jeden ausgewählten Ski wird eine eigene Ausleihe für die angegebene Person
gestartet. Der Ausleihbeginn ist heute; das inklusive Mietende kann gewählt
werden. Tages- und Gesamtpreis werden angezeigt. Die Auswahl steht im
Dashboard und im Personen-Tab zur Verfügung. Reservierungen können im
Dashboard und im Reservierungs-Tab über dieselbe Checkbox-Auswahl mit
zukünftigem Beginn und inklusive Mietende angelegt werden. Pro Person können mehrere konkrete Ski, auch desselben Modells,
gleichzeitig reserviert werden. Mehrere aktive Reservierungen einer Person
werden im Personen-Tab über jeweils einen gemeinsamen Button zum Starten,
Bearbeiten oder Löschen verwaltet. Beim Bearbeiten können bestehende
Reservierungen abgewählt und neue verfügbare Ski per Checkbox ausgewählt werden.
Bei einer neuen Ausleihe oder Reservierung wird der Zeitraum zuerst abgefragt.
Danach zeigt die Auswahl ausschließlich Ski, die im gesamten inklusiven
Zeitraum weder verliehen noch reserviert sind. Reservierungen können erst am
gespeicherten Beginn in eine Ausleihe umgewandelt werden.
Eine Umwandlung vor dem Reservierungsbeginn wird auch serverseitig abgelehnt.
Für Ausleihen und Reservierungen werden Nachname und Vorname abgefragt.
Passende Personen werden anschließend in einer Radio-Auswahl angezeigt, aus
der genau eine Person ausgewählt werden kann. Abbrechen oder keine Treffer
beenden den Vorgang ohne weitere Eingaben.
Für jede aktive Ausleihe wird zusätzlich im Personen-Tab der Button
„Ausleihe beenden“ angezeigt. Nach der Bestätigung werden Personen-, Ski- und
Dashboard-Ansicht automatisch aktualisiert.
Der Button öffnet ein Pop-up mit den ausgeliehenen Skiern; die zu beendenden
Ausleihen können per Checkbox ausgewählt werden.

Ski werden in der Ski-Verwaltung und in Ski-Auswahl-Pop-ups aufsteigend nach
Marke, Modell, Länge und Ski-ID sortiert. Personen werden im Personen-Tab nach
Nachname, Vorname und Personen-ID sortiert.

## API-Kurzreferenz

| Methode | Pfad | Zweck |
|---|---|---|
| `GET` | `/api/dashboard` | Kennzahlen |
| `GET`, `POST` | `/api/skis` | Ski auflisten oder anlegen |
| `GET` | `/api/skis/available?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD` | Ski für den vollständigen Zeitraum auflisten |
| `GET`, `POST` | `/api/people` | Personen auflisten oder registrieren |
| `GET`, `POST` | `/api/loans` | Ausleihen auflisten oder starten |
| `DELETE` | `/api/loans/{skiID}` | Ausleihe beenden |
| `GET`, `POST` | `/api/reservations` | Reservierungen auflisten oder anlegen |
| `PUT`, `DELETE` | `/api/reservations/{id}` | Reservierung bearbeiten oder verwerfen |
| `POST` | `/api/reservations/{id}/convert` | Reservierung in Ausleihe umwandeln |

Ski liefern ihren Tagespreis in Cent. Ausleihen und Reservierungen enthalten
Beginn und Ende; Ausleihen enthalten zusätzlich Miettage und Gesamtpreis.

## Qualität

```text
go fmt ./...
go test ./...
go test -race ./...
go vet ./...
git diff --check
```

Ein separater Linter ist im Projekt nicht eingerichtet.

# Fachlicher Überblick

Ziel ist es ein Verleihsystem für Ski zu entwickeln. Es soll als Webanwendung implementiert werden. 
Über Reiter "Dashboard", "Ski-Verwaltung", "Personen" und "Reservierungen" soll es möglich sein, innerhalb der Webanwendung zwischen verschiedenen Seiten zu wechseln. Die Funktionsweise der Seiten soll wie folgt sein:

## Dashboard
Im Dashboard sollen die wichtigsten statistischen Größen des Skiverleihsystems auf einen Blick zu sehen sein. Diese beinhalten:
- Anzahl verfügbarer Ski
- Anzahl ausgeliehener Ski
- Gesamtanzahl an Ski
- Anzahl registrierter Personen
Über einen Schnellzugriff im Dashboard soll es möglich sein, Ski zu verwalten , neue Personen zu registrieren und neue Ausleihen anzumelden.

## Ski-Verwaltung
In der Ski-Verwaltung sollen alle verfügbaren Ski zeilenweise in tabellarischer Struktur aufgelistet sein. Pro Ski werden die folgenden Attribute gespeichert:
- "Ski-Id" [Format: "S-" gefolgt von einer vierstelligen Zahl]
- "Marke"
- "Modell" 
- "Status" ["verliehen"/"reserviert"/"verfügbar"]
- "Person-Id" [Format: "P-" gefolgt von einer vierstelligen Zahl], soll dargestellt werden als "Mieter-Id"
- "Person-Name", soll dargestellt werden als "Mieter-Name"
Jeder bereits verliehene Ski soll einen roten Button "Leihe beenden" besitzen, über den die Miete erfolgreich beendet wird und der Ski wieder als verfügbar markiert wird.

## Personen
Das Personen-Tab soll aus zwei Komponenten bestehen.
Im oberen Feld soll es möglich sein, über einen Button "Neue Person registrieren" eine neue Person zu registrieren. Durch Anklicken des Buttons soll sich ein Pop-Up Fenster öffnen, das die Attribute "Vorname", "Nachname", "Geburtsdatum", "Ausweis-Nummer", "Land", "PLZ", "Straße", "Hausnummer" und "Handynummer" abfragt. Um die Registrierung abzuschließen, soll es nötig sein, ein Ausweisfoto hochzuladen (png- oder jpeg-Format). Bei Registrierung soll einer Person automatisch eine eindeutige "Person-Id" zugewiesen werden im Format "P-" gefolgt von einer vierstelligen Zahl.

Im unteren Feld des Personen-Tab werden alle registrierten Personen alphabetisch in tabellarischer Struktur aufgelistet. Für jede Person sollen die Attribute
- "Person-Id"
- "Person-Name" (Kombination aus Vor- und Nachname)
- "Anzahl ausgeliehener Ski"
- "Ausgeliehene Ski" (Aufzählund aller Ski, die aktuell an diese Person verliehen sind)
- "Ausweis-Nummer"
- "Ausweis-Foto"
- "Adresse" und "Handynummer"

Über den Button "Ausweis ansehen" kann das hinterlegte Ausweisdokument im
Personen-Tab angezeigt werden. Jede Person besitzt außerdem den Button
"Stammdaten bearbeiten". Die Eingaben werden mit denselben Pflichtfeld-,
Alters-, PLZ- und Fotoformatprüfungen wie bei der Neuanlage validiert.

Für jede Person, die keine Reervierung getätigt hat, soll ein grüner Button "Neue Ausleihe" existieren, über den eine neue Ausleihe aus den verfügbaren Skiern gemacht werden kann.

Für Personen, die einen Ski reserviert haben, soll ein grüner Button "Miete reservierter Ski beginnen" existieren. Darüber hinaus soll ein gelber Button "Reservierung bearbeiten" und ein roter Button "Reservierung löschen" existieren.

Personen, die bereits einen Ski mieten, soll es ebenso erlaubt sein, die Miete weiterer Ski zu beginnen.

# Reservierungen
Das Reservierungen-Tab soll aus zwei Komponenten bestehen.

Im oberen Feld sollen neue Ski-Reservierungen angemeldet werden können.

Im unteren Feld sollen alle Personen, die eine aktive Reservierung haben, zusammen mit den reservierten Skiern aufgelistet werden. Über einen grünen Button "Miete reservierter Ski beginnen" soll die Reservierung beendet und automatisch eine Miete begonnen werden. Über einen roten Button "Reservierung verwerfen" soll eine Reservierung beendet werden ohne eine Miete zu starten. Über einen gelben Button "Reservierung bearbeiten" soll die Reservierung bearbeitet werden.

# Zusätzliche Anforderungen
Jeder Button soll ein Pop-Up Fenster hervorrufen. In diesem sollen (falls nötig) Informationen eingegeben werden können oder die gewählte Aktion bestätigt werden. Zur Versionierung soll ein GitHub-Repo angelegt werden.
