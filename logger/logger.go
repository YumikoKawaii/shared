package logger

import (
	"context"
	"sync"
)

// A global variable so that log functions can be directly accessed
var log Logger = DefaultLogger()

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

type LogLvl string

const (
	// Debug has verbose message
	DebugLvl LogLvl = "debug"
	//Info is default log level
	InfoLvl LogLvl = "info"
	// Warn is for logging messages about possible issues
	WarnLvl LogLvl = "warn"
	// Error is for logging errors
	ErrorLvl LogLvl = "error"
	// Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	FatalLvl LogLvl = "fatal"
)

const (
	// Debug has verbose message
	debugLvl = "debug"
	//Info is default log level
	infoLvl = "info"
	// Warn is for logging messages about possible issues
	warnLvl = "warn"
	// Error is for logging errors
	errorLvl = "error"
	// Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	fatalLvl = "fatal"
)

var once sync.Once

// Logger is our contract for the logger
type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...interface{})

	Info(msg string)
	Infof(format string, args ...interface{})

	Warn(msg string)
	Warnf(format string, args ...interface{})

	Error(msg string)
	Errorf(format string, args ...interface{})

	Fatal(msg string)
	Fatalf(format string, args ...interface{})

	Panic(msg string)
	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
	WithPrefix(prefix string) Logger
	WithTraceId(ctx context.Context) Logger

	GetDelegate() interface{}

	CloseWriter() error
}

// Configuration stores the config for the logger
type Configuration struct {
	EnableConsole     bool   `json:"enable_console" mapstructure:"enable_console" yaml:"enable_console"`
	ConsoleJSONFormat bool   `json:"console_json_format" mapstructure:"console_json_format" yaml:"console_json_format"`
	ConsoleLevel      string `json:"console_level" mapstructure:"console_level" yaml:"console_level"`
	EnableFile        bool   `json:"enable_file" mapstructure:"enable_file" yaml:"enable_file"`
	FileJSONFormat    bool   `json:"file_json_format" mapstructure:"file_json_format" yaml:"file_json_format"`
	FileLevel         string `json:"file_level" mapstructure:"file_level" yaml:"file_level"`
	FileLocation      string `json:"file_location" mapstructure:"file_location" yaml:"file_location"`
}

func DefaultConfig() Configuration {
	return Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: false,
		ConsoleLevel:      "info",
		EnableFile:        false,
	}
}

// DefaultLogger creates default logger, which uses zap sugarlogger and outputs to console
func DefaultLogger() Logger {
	cfg := DefaultConfig()
	logger, _ := newZapLogger(cfg)
	return logger
}

// InitLogger returns an instance of logger
func InitLogger(conf Configuration) (Logger, error) {
	var err error
	once.Do(func() {
		log, err = NewLogger(conf)
	})
	return log, err
}

func NewLogger(conf Configuration) (Logger, error) {
	return newZapLogger(conf)
}

func Debug(msg string) {
	log.Debugf(msg)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(msg string) {
	log.Infof(msg)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(msg string) {
	log.Warnf(msg)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(msg string) {
	log.Errorf(msg)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(msg string) {
	log.Fatalf(msg)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panic(msg string) {
	log.Panicf(msg)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}

func WithPrefix(prefix string) Logger {
	return log.WithPrefix(prefix)
}

func WithTraceId(ctx context.Context) Logger {
	return log.WithTraceId(ctx)
}

func Get() Logger {
	return log
}

func GetDelegate() interface{} {
	return log.GetDelegate()
}
