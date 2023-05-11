package logger

import "github.com/charmbracelet/log"

type Factory interface {
	Default() Logger
	With(keyvals ...any) Logger
	WithPrefix(prefix string) Logger
}

var defaultFactory Factory = &factory{}

type factory struct{}

func SetFactory(factory Factory) {
	defaultFactory = factory
}

func (f *factory) With(keyvals ...any) Logger {
	return log.With(keyvals...)
}

func (f *factory) WithPrefix(prefix string) Logger {
	return log.WithPrefix(prefix)
}

func (f *factory) Default() Logger {
	return log.Default()
}
