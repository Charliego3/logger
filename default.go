package logger

import (
	"github.com/charmbracelet/log"
)

type DefaultLog struct {
	*log.Logger
}

func (l *DefaultLog) SetLevel(level Level) {
	l.Logger.SetLevel(log.ParseLevel(level.String()))
}

func (l *DefaultLog) GetLevel() Level {
	level, _ := LevelString(l.Logger.GetLevel().String())
	return level
}

type defaultFactory struct{}

func (f *defaultFactory) With(keyvals ...any) Logger {
	return &DefaultLog{log.With(keyvals...)}
}

func (f *defaultFactory) WithPrefix(prefix string) Logger {
	return &DefaultLog{log.WithPrefix(prefix)}
}

func (f *defaultFactory) Default() Logger {
	return &DefaultLog{log.Default()}
}

func init() {
	SetFactory(&defaultFactory{})
}
