package test

import (
	"github.com/charliego93/logger"
	_ "github.com/charliego93/logger/adapter/zerolog"
	"testing"
)

func TestInfo(t *testing.T) {
	l := logger.WithPrefix("hahahha")
	l = logger.With("aaaaaa", "bbbbbbbb")
	l.Debug("this is debug message", "key", "value", "age", 34)
	l.Info("this is info message", "key", "value", "age", 34)
	l.Warn("this is warn message", "key", "value", "age", 34)
	l.Error("this is error message", "key", "value", "age", 34)
	l.Fatal("this is fatal message", "key", "value", "age", 34)
}
