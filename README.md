# coco-logger

Ein schlanker File-Logger für Go mit täglicher Ordnerstruktur und einfacher API.

## Features

- Schreiben in Logdateien mit automatischer, datumsbasierter Ordnerstruktur (YYYY/MM/DD)
- Thread-sicheres Logging (Mutex-geschützt)
- Einfache Methoden: Info, Warning, Error, Close
- Konfigurierbarer Log-Verzeichnis-Pfad

## Installation

Nutze Go Modules und füge das Modul deinem Projekt hinzu. Passe den Modulpfad an dein tatsächliches Repository an.

```bash
go get github.com/your-org/coco-logger
```

Importiere den Logger aus dem `src/logger` Paket:

```go
import logger "github.com/your-org/coco-logger/src/logger"
```

## Quick Start

```go
package main

import (
    logger "github.com/your-org/coco-logger/src/logger"
)

func main() {
    // Lege dein Log-Verzeichnis fest (wird nach Datum strukturiert)
    log, err := logger.NewLogger("app.log", "./logs")
    if err != nil {
        panic(err)
    }
    defer log.Close()

    log.Info("service starting: version %s", "0.1.0")
    log.Warning("slower than expected: duration_ms=%d", 120)
    log.Error("failed to do work: err=%s", "some error")
}
```

Die Logs werden in einem Pfad wie `./logs/2026/02/14/app.log` geschrieben.

## API

- `NewLogger(fileName string, logDir string) (Logger, error)`
  - Erstellt einen neuen Logger, der in `logDir/YYYY/MM/DD/fileName` schreibt
  - `logDir` muss gesetzt werden (kein Fallback)
- `Info(msg string, args ...interface{})`
- `Warning(msg string, args ...interface{})`
- `Error(msg string, args ...interface{})`
- `Close()` schließt die Dateiressource

### Format

Jeder Logeintrag wird als einfache Textzeile geschrieben:

```
[YYYY-MM-DD HH:MM:SS] [LEVEL] message\n
```

Beispiel:

```
[2026-02-14 10:23:45] [INFO] service starting: version 0.1.0
```

## Best Practices

- Immer `Close()` aufrufen (z. B. mit `defer`), um Dateihandles zu schließen
- Wähle ein sinnvolles `logDir` (z. B. `./logs` oder `/var/log/myapp`)
- Gib strukturierte Informationen als Teil der Nachricht aus (key=value), z. B. `duration_ms=120`
- Vermeide das Loggen sensibler Daten (PII/Secrets)

## Tests

Falls Tests vorhanden sind, ausführen mit:

```bash
go test ./...
```

## Versionierung

Dieses Projekt folgt semantischer Versionierung, bis v1.0 können Breaking Changes auftreten.

## Lizenz

Dieses Projekt ist unter der MIT-Lizenz veröffentlicht. Siehe `LICENSE`.
