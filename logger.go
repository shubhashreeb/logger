package logger

import "errors"

type Fields map[string]interface{}

const (
	DebugLevel            = "debug"
	InfoLevel             = "info"
	WarnLevel             = "warn"
	ErrorLevel            = "error"
	FatalLevel            = "fatal"
	InstanceZapLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
	log                      Logger
)

//Logger is our contract for the logger
type Logger interface {
	Debug(format string, args ...interface{})

	Info(format string, args ...interface{})

	Warn(format string, args ...interface{})

	Error(format string, args ...interface{})

	Fatal(format string, args ...interface{})

	Panic(format string, args ...interface{})

	WithContext(keyValues Fields) Logger
}

// Configuration stores the config for the logger
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func NewLogger(config Configuration) (Logger, error) {
	return NewZapLogger(config), nil
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}

func Panic(format string, args ...interface{}) {
	log.Panic(format, args...)
}

func WithContext(keyValues Fields) Logger {
	return log.WithContext(keyValues)
}
