package logger

import "sync"

type Factory interface {
	Default() Logger
	With(keyvals ...any) Logger
	WithPrefix(prefix string) Logger
}

var (
	// factory create Logger instance
	// default using github.com/charmbracelet/log
	factory Factory

	// fmux lock
	fmux sync.Mutex
)

func SetFactory(f Factory) {
	fmux.Lock()
	defer fmux.Unlock()
	factory = f
	SetLogger(factory.Default())
}
