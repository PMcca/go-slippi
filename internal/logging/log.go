package logging

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
)

// NewLogger returns a new Logger.
func NewLogger() zerolog.Logger {
	l := zerolog.New(os.Stdout).With().
		Timestamp().
		Caller().
		Stack().
		Logger()

	// Format output to be "<level> <source> <msg>" only
	l = l.Level(zerolog.InfoLevel).Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		NoColor:      true,
		PartsExclude: []string{"time", "trace"},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(i.(string))
		}})
	return l
}
