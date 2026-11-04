# Validation: Projektgerüst und lokales Git

Die Phase ist erfolgreich und der Branch kann in `main` gemerged werden, wenn
alle folgenden Kriterien erfüllt sind:

1. Das Repository enthält eine gültige Go-Modulkonfiguration mit dokumentierter
   Go-Version und dokumentiertem Modulnamen.
2. Der Einstiegspunkt startet mit `go run .` ohne Fehler.
3. Der minimale HTTP-Server antwortet am festgelegten lokalen Port auf seinen
   Start- oder Health-Endpunkt erfolgreich.
4. Die geplante Grundstruktur mit Einstiegspunkt und den vorgesehenen
   `internal/`-Bereichen ist vorhanden und enthält keine unnötigen
   Abhängigkeiten.
5. `go test ./...` läuft erfolgreich durch.
6. `gofmt -w .` beziehungsweise `go fmt ./...` erzeugt keine verbleibenden
   Formatierungsänderungen.
7. Der Branch enthält einen lokalen Git-Commit mit den Änderungen dieser Phase.
8. Die Änderungen bleiben auf das Projektgerüst und die lokale Versionierung
   beschränkt; Domain-, API- und Frontend-Funktionalität wird nicht vorgezogen.

## Merge-Entscheidung

Der Merge ist freigegeben, sobald alle acht Kriterien durch Kommandoausgaben,
einen manuellen Health-Check oder eine nachvollziehbare Dateiprüfung belegt sind.
Bei einem fehlgeschlagenen Kriterium bleibt der Branch offen, bis die Ursache
behoben und die vollständige Prüfung erneut ausgeführt wurde.
