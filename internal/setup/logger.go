package setup

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type LogLevel string

const (
	TraceLogLevel LogLevel = "trace"
	DebugLogLevel LogLevel = "debug"
	InfoLogLevel  LogLevel = "info"
	WarnLogLevel  LogLevel = "warn"
	ErrorLogLevel LogLevel = "error"
	FatalLogLevel LogLevel = "fatal"
	PanicLogLevel LogLevel = "panic"
)

func SetupLogger(config Configuration) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var level zerolog.Level
	switch config.LogLevel {
	case TraceLogLevel:
		level = zerolog.TraceLevel
	case DebugLogLevel:
		level = zerolog.DebugLevel
	case InfoLogLevel:
		level = zerolog.InfoLevel
	case WarnLogLevel:
		level = zerolog.WarnLevel
	case ErrorLogLevel:
		level = zerolog.ErrorLevel
	case FatalLogLevel:
		level = zerolog.FatalLevel
	case PanicLogLevel:
		level = zerolog.PanicLevel
	default:
		level = zerolog.InfoLevel
	}

	var writer io.Writer = zerolog.ConsoleWriter{Out: os.Stdout}

	logger := zerolog.New(writer).With().Timestamp().Caller().Logger().Level(level)
	return logger
}
