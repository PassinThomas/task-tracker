package utils

import (
    "log/slog"
    "os"
)

var Logger *slog.Logger

// Init configure le logger.
// Par défaut (debug=false) : il est muet (Discard).
// Si debug=true : il écrit tout sur Stderr.
func InitLogger(debug bool) {
    if debug {
        // Mode verbeux : On affiche tout sur Stderr
        opts := &slog.HandlerOptions{Level: slog.LevelDebug}
        Logger = slog.New(slog.NewTextHandler(os.Stderr, opts))
    } else {
        // Mode silencieux : On jette les logs à la poubelle (/dev/null)
        // C'est plus performant et plus sûr que de filtrer le niveau Info
        Logger = slog.New(slog.NewTextHandler(os.Discard, nil)) // Discard ne fait rien
    }
}

// Debug est ton SEUL outil de log nécessaire
func Debug(msg string, args ...any) {
    if Logger != nil {
        Logger.Debug(msg, args...)
    }
}