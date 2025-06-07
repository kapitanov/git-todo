package logutil

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ConfigureLogger(verbose bool) {
	if verbose {
		log.Logger = log.Output(
			zerolog.NewConsoleWriter(
				func(w *zerolog.ConsoleWriter) {
					w.Out = os.Stderr
				},
			),
		).
			Level(zerolog.DebugLevel)
	} else {
		log.Logger = log.Level(zerolog.Disabled)
	}
}

func WithTestLogger(t *testing.T, fn func()) {
	originalLogger := log.Logger
	defer func() {
		log.Logger = originalLogger
	}()

	log.Logger = log.Output(zerolog.NewConsoleWriter(zerolog.ConsoleTestWriter(t))).Level(zerolog.DebugLevel)

	fn()
}
