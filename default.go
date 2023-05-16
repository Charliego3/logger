package logger

import (
	"github.com/charmbracelet/log"
)

type charmbraceletLog struct {
	*log.Logger
}

func (l *charmbraceletLog) SetLevel(level Level) {
	l.Logger.SetLevel(log.ParseLevel(level.String()))
}

type defaultFactory struct{}

func (f *defaultFactory) With(keyvals ...any) Logger {
	return &charmbraceletLog{log.With(keyvals...)}
}

func (f *defaultFactory) WithPrefix(prefix string) Logger {
	return &charmbraceletLog{log.WithPrefix(prefix)}
}

func (f *defaultFactory) Default() Logger {
	return &charmbraceletLog{log.Default()}
}

func init() {
	SetFactory(&defaultFactory{})
}
