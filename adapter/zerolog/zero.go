package zerolog

import (
	"fmt"
	"github.com/charliego3/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

type ZeroLog struct {
	*zerolog.Logger
}

func (l *ZeroLog) Debug(msg any, keyvals ...any) {
	l.log(l.Logger.Debug(), false, msg, keyvals...)
}

func (l *ZeroLog) Info(msg any, keyvals ...any) {
	l.log(l.Logger.Info(), false, msg, keyvals...)
}

func (l *ZeroLog) Warn(msg any, keyvals ...any) {
	l.log(l.Logger.Warn(), false, msg, keyvals...)
}

func (l *ZeroLog) Error(msg any, keyvals ...any) {
	l.log(l.Logger.Error(), false, msg, keyvals...)
}

func (l *ZeroLog) Fatal(msg any, keyvals ...any) {
	l.log(l.Logger.Fatal(), false, msg, keyvals...)
}

func (l *ZeroLog) Print(msg any, keyvals ...any) {
	l.log(l.Logger.Log(), false, msg, keyvals...)
}

func (l *ZeroLog) Debugf(format string, args ...any) {
	l.log(l.Logger.Debug(), true, format, args...)
}

func (l *ZeroLog) Infof(format string, args ...any) {
	l.log(l.Logger.Info(), true, format, args...)
}

func (l *ZeroLog) Warnf(format string, args ...any) {
	l.log(l.Logger.Warn(), true, format, args...)
}

func (l *ZeroLog) Errorf(format string, args ...any) {
	l.log(l.Logger.Error(), true, format, args...)
}

func (l *ZeroLog) Fatalf(format string, args ...any) {
	l.log(l.Logger.Fatal(), true, format, args...)
}

func (l *ZeroLog) SetTimeFormat(format string) {
	zerolog.TimeFieldFormat = format
}

func (l *ZeroLog) SetOutput(out io.Writer) {
	*l.Logger = l.Logger.Output(out)
}

func (l *ZeroLog) SetLevel(level logger.Level) {
	parseLevel, err := zerolog.ParseLevel(level.String())
	if err != nil {
		l.Logger.Warn().Err(err).Str("level", level.String()).Msg("failed set logger level")
		parseLevel = zerolog.InfoLevel
	}
	l.Logger.Level(parseLevel)
}

func (l *ZeroLog) GetLevel() logger.Level {
	level, _ := logger.LevelString(l.Logger.GetLevel().String())
	return level
}

func (l *ZeroLog) log(e *zerolog.Event, format bool, msg any, args ...any) {
	if format {
		e.Msg(fmt.Sprintf(msg.(string), args...))
		return
	}

	if len(args) > 0 {
		e.Fields(getFields(args))
	}
	e.Any("msg", msg).Send()
}

func getFields(keyvals []any) map[string]any {
	length := len(keyvals)
	if length > 0 {
		if length%2 != 0 {
			length -= 1
		}
		fields := make(map[string]any)
		for i := 0; i < length; i = i + 2 {
			fields[fmt.Sprintf("%v", keyvals[i])] = keyvals[i+1]
		}
		return fields
	}
	return nil
}

type zFactory struct{}

func (f *zFactory) With(keyvals ...any) logger.Logger {
	l := newContext().Fields(getFields(keyvals)).Logger()
	return &ZeroLog{&l}
}

func (f *zFactory) WithPrefix(prefix string) logger.Logger {
	return f.With("prefix", prefix)
}

func (f *zFactory) Default() logger.Logger {
	l := newContext().Logger()
	return &ZeroLog{&l}
}

func newContext() zerolog.Context {
	return log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With()
}

func init() {
	logger.SetFactory(&zFactory{})
}
