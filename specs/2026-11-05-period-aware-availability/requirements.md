# Requirements: Zeitraumbezogene Ski-Verfügbarkeit

## Scope Decisions

- Der Zeitraum wird bei einer neuen Ausleihe oder Reservierung zuerst abgefragt.
- Ein konkreter Ski ist nur verfügbar, wenn im gesamten inklusiven Zeitraum
  keine aktive Ausleihe und keine aktive Reservierung desselben Skis liegt.
- Start- und Enddatum zählen inklusiv; jede Überschneidung macht den Ski
  nicht auswählbar.
- Ausleihen beginnen weiterhin ausschließlich am aktuellen Tag.
- Reservierungen dürfen heute oder später beginnen.
- Eine Reservierung darf ausschließlich am gespeicherten Reservierungsbeginn
  in eine Ausleihe umgewandelt werden.
- Wenn kein Ski verfügbar ist, endet der Vorgang mit einer verständlichen
  Meldung ohne Ski- oder Personenauswahl.

## Nicht-Ziele

- Keine Persistenz, Authentifizierung oder neue externe Abhängigkeit.
- Keine Änderung der bestehenden Mengen- und Checkbox-Auswahl.
- Keine Änderung der Preisberechnung oder der Rückgabeabläufe.
