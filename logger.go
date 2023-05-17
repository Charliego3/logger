package logger

import (
	"io"
	"sync"
)

// Logger is a logger face, it can be using any instance
// see SetFactory set default factory, provider instance
type Logger interface {
	// PLogger is a print logger interface
	// this is print message and key values
	PLogger

	// FLogger is a formatter logger interface
	// this can be format message and args
	FLogger

	GetLevel() Level
	SetTimeFormat(string)
	SetOutput(out io.Writer)
	SetLevel(l Level)
}

type PLogger interface {
	Debug(msg any, keyvals ...any)
	Info(msg any, keyvals ...any)
	Warn(msg any, keyvals ...any)
	Error(msg any, keyvals ...any)
	Fatal(msg any, keyvals ...any)
	Print(msg any, keyvals ...any)
}

type FLogger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Printf(format string, args ...any)
}

var (
	// defaultLogger created by Factory.Default()
	defaultLogger Logger

	mux sync.RWMutex
)

// SetLogger replace default logger
func SetLogger(logger Logger) {
	mux.Lock()
	defer mux.Unlock()
	defaultLogger = logger
}

// With create logger with key value pairs
//
//	With("thread", "worker")
func With(keyvals ...any) Logger {
	return factory.With(keyvals...)
}

// WithPrefix returns logger within prefix
func WithPrefix(prefix string) Logger {
	return factory.WithPrefix(prefix)
}

func Default() Logger {
	mux.RLock()
	defer mux.RUnlock()
	return defaultLogger
}

func GetLevel() Level {
	return Default().GetLevel()
}

// SetLevel change default logger output level
func SetLevel(l Level) {
	Default().SetLevel(l)
}

func SetOutput(writer io.Writer) {
	Default().SetOutput(writer)
}

func SetTimeFormat(format string) {
	Default().SetTimeFormat(format)
}

func Debug(msg interface{}, keyvals ...interface{}) {
	Default().Debug(msg, keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	Default().Info(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	Default().Warn(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	Default().Error(msg, keyvals...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	Default().Fatal(msg, keyvals...)
}

func Print(msg interface{}, keyvals ...interface{}) {
	Default().Print(msg, keyvals...)
}

func Debugf(format string, args ...interface{}) {
	Default().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Default().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Default().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Default().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Default().Fatalf(format, args...)
}

func Printf(format string, args ...interface{}) {
	Default().Printf(format, args...)
}
