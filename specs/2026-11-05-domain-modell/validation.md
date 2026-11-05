# Validation: Domain-Modell

Die Phase gilt als erfolgreich und der Branch ist für einen Merge nach `main`
bereit, wenn alle folgenden Kriterien erfüllt sind:

1. `internal/domain/` enthält typsichere Modelle für Ski, Person, Ausleihe und
   Reservierung mit den in Mission und Requirements festgelegten Feldern.
2. Ski- und Personen-IDs werden ausschließlich im Format `S-0000`
   beziehungsweise `P-0000` akzeptiert oder erzeugt.
3. Für Ski sind ausschließlich `verfügbar`, `reserviert` und `verliehen`
   möglich; ungültige und fachlich unzulässige Übergänge liefern explizite
   Fehler.
4. Registrierung, Ausleihe, Rückgabe sowie Anlegen, Bearbeiten, Umwandeln und
   Verwerfen von Reservierungen sind durch Domain-Methoden abbildbar.
5. Eine Person ohne gültiges verpflichtendes PNG-Ausweisfoto kann nicht
   erfolgreich registriert werden.
6. Inkonsistente Zuordnungen, doppelte aktive Reservierungen,
   Rückgaben ohne aktive Ausleihe und Ausleihen nicht verfügbarer Ski werden
   zurückgewiesen.
7. `go test ./...` läuft erfolgreich durch und enthält aussagekräftige
   Unit-Tests für Erfolgs- und Fehlerfälle der Domain.
8. `gofmt` beziehungsweise `go fmt ./...` hinterlässt keine Änderungen;
   zusätzliche Dependencies und HTTP-, Store- oder Frontend-Code wurden nicht
   eingeführt.
9. Die Implementierung bleibt mit
   [README.md](../../README.md), [mission.md](../mission.md) und
   [techstack.md](../techstack.md) konsistent.

## Merge-Entscheidung

Der Merge nach `main` ist freigegeben, sobald alle neun Kriterien durch
Testausgaben, Formatierungsprüfung und eine kurze Diff-Prüfung belegt sind.
Bei einem fehlgeschlagenen Kriterium bleibt der Branch offen, bis die Ursache
behoben und die vollständige Prüfung wiederholt wurde.

