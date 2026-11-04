# Skiverleih-Prototyp

Das Projekt ist ein leicht verständlicher Prototyp für die spätere
Skiverleih-Webanwendung. Die erste Phase stellt ein minimales Go-Backend mit
Health-Endpunkt und der vorgesehenen Verzeichnisstruktur bereit.

## Technischer Stand

- Go-Modul: `skirental`
- Unterstützte Go-Version: `1.22`
- Lokaler Server-Port: `8080`
- Health-Endpunkt: `GET http://localhost:8080/healthz`
- Speicherung und fachliche Endpunkte folgen in späteren Phasen.

## Entwicklung

```text
go run .
go test ./...
gofmt -w .
```

Der Server meldet sich unter `http://localhost:8080/healthz` mit
`{"status":"ok"}`. Die Pakete unter `internal/domain/`, `internal/store/` und
`internal/handler/` sind für die folgenden Entwicklungsphasen vorbereitet.

# Ziel

Ziel ist es ein Verleihsystem für Ski zu entwickeln. Es soll als Webanwendung implementiert werden. 
Über Reiter "Dashboard", "Ski-Verwaltung", "Personen" und "Reservierungen" soll es möglich sein, innerhalb der Webanwendung zwischen verschiedenen Seiten zu wechseln. Die Funktionsweise der Seiten soll wie folgt sein:

# Dashboard
Im Dashboard sollen die wichtigsten statistischen Größen des Skiverleihsystems auf einen Blick zu sehen sein. Diese beinhalten:
- Anzahl verfügbarer Ski
- Anzahl ausgeliehener Ski
- Gesamtanzahl an Ski
- Anzahl registrierter Personen
Über einen Schnellzugriff im Dashboard soll es möglich sein, Ski zu verwalten , neue Personen zu registrieren und neue Ausleihen anzumelden.

# Ski-Verwaltung
In der Ski-Verwaltung sollen alle verfügbaren Ski zeilenweise in tabellarischer Struktur aufgelistet sein. Pro Ski werden die folgenden Attribute gespeichert:
- "Ski-Id" [Format: "S-" gefolgt von einer vierstelligen Zahl]
- "Marke"
- "Modell" 
- "Status" ["verliehen"/"reserviert"/"verfügbar"]
- "Person-Id" [Format: "P-" gefolgt von einer vierstelligen Zahl], soll dargestellt werden als "Mieter-Id"
- "Person-Name", soll dargestellt werden als "Mieter-Name"
Jeder bereits verliehene Ski soll einen roten Button "Leihe beenden" besitzen, über den die Miete erfolgreich beendet wird und der Ski wieder als verfügbar markiert wird.

# Personen
Das Personen-Tab soll aus zwei Komponenten bestehen.
Im oberen Feld soll es möglich sein, über einen Button "Neue Person registrieren" eine neue Person zu registrieren. Durch Anklicken des Buttons soll sich ein Pop-Up Fenster öffnen, das die Attribute "Vorname", "Nachname",  "Geburtsdatum", "Ausweis-Nummer" abfragt. Um die Registrierung abzuschließen, soll es nötig sein, ein Ausweisfoto hochzuladen (png-Format). Bei Registrierung soll einer Person automatisch eine eindeutige "Person-Id" zugewiesen werden im Format "P-" gefolgt von einer vierstelligen Zahl.

Im unteren Feld des Personen-Tab werden alle registrierten Personen alphabetisch in tabellarischer Struktur aufgelistet. Für jede Person sollen die Attribute
- "Person-Id"
- "Person-Name" (Kombination aus Vor- und Nachname)
- "Anzahl ausgeliehener Ski"
- "Ausgeliehene Ski" (Aufzählund aller Ski, die aktuell an diese Person verliehen sind)
- "Ausweis-Nummer"
- "Ausweis-Foto"

Für jede Person, die keine Reervierung getätigt hat, soll ein grüner Button "Neue Ausleihe" existieren, über den eine neue Ausleihe aus den verfügbaren Skiern gemacht werden kann.

Für Personen, die einen Ski reserviert haben, soll ein grüner Button "Miete reservierter Ski beginnen" existieren. Darüber hinaus soll ein gelber Button "Reservierung bearbeiten" und ein roter Button "Reservierung löschen" existieren.

Personen, die bereits einen Ski mieten, soll es ebenso erlaubt sein, die Miete weiterer Ski zu beginnen.

# Reservierungen
Das Reservierungen-Tab soll aus zwei Komponenten bestehen.

Im oberen Feld sollen neue Ski-Reservierungen angemeldet werden können.

Im unteren Feld sollen alle Personen, die eine aktive Reservierung haben, zusammen mit den reservierten Skiern aufgelistet werden. Über einen grünen Button "Miete reservierter Ski beginnen" soll die Reservierung beendet und automatisch eine Miete begonnen werden. Über einen roten Button "Reservierung verwerfen" soll eine Reservierung beendet werden ohne eine Miete zu starten. Über einen gelben Button "Reservierung bearbeiten" soll die Reservierung bearbeitet werden.

# Zusätzliche Anforderungen
Die Miete oder Reservierungen beginnen stets sofort. Die App soll keine Preise verwalten. Jeder Button soll ein Pop-Up Fenster hervorrufen. In diesem sollen (falls nötig) Informationen eingegeben werden können oder die gewählte Aktion bestätigt werden. Zur Versionierung soll ein GitHub-Repo angelegt werden.

