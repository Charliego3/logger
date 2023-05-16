package zerolog

import (
	"fmt"
	"github.com/charliego93/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

type zLog struct {
	*zerolog.Logger
}

func (l *zLog) Debug(msg any, keyvals ...any) {
	l.log(l.Logger.Debug(), false, msg, keyvals...)
}

func (l *zLog) Info(msg any, keyvals ...any) {
	l.log(l.Logger.Info(), false, msg, keyvals...)
}

func (l *zLog) Warn(msg any, keyvals ...any) {
	l.log(l.Logger.Warn(), false, msg, keyvals...)
}

func (l *zLog) Error(msg any, keyvals ...any) {
	l.log(l.Logger.Error(), false, msg, keyvals...)
}

func (l *zLog) Fatal(msg any, keyvals ...any) {
	l.log(l.Logger.Fatal(), false, msg, keyvals...)
}

func (l *zLog) Print(msg any, keyvals ...any) {
	l.log(l.Logger.Log(), false, msg, keyvals...)
}

func (l *zLog) Debugf(format string, args ...any) {
	l.log(l.Logger.Debug(), true, format, args...)
}

func (l *zLog) Infof(format string, args ...any) {
	l.log(l.Logger.Info(), true, format, args...)
}

func (l *zLog) Warnf(format string, args ...any) {
	l.log(l.Logger.Warn(), true, format, args...)
}

func (l *zLog) Errorf(format string, args ...any) {
	l.log(l.Logger.Error(), true, format, args...)
}

func (l *zLog) Fatalf(format string, args ...any) {
	l.log(l.Logger.Fatal(), true, format, args...)
}

func (l *zLog) SetTimeFormat(format string) {
	zerolog.TimeFieldFormat = format
}

func (l *zLog) SetOutput(out io.Writer) {
	*l.Logger = l.Logger.Output(out)
}

func (l *zLog) SetLevel(level logger.Level) {
	parseLevel, err := zerolog.ParseLevel(level.String())
	if err != nil {
		l.Logger.Warn().Err(err).Str("level", level.String()).Msg("failed set logger level")
		parseLevel = zerolog.InfoLevel
	}
	l.Logger.Level(parseLevel)
}

func (l *zLog) log(e *zerolog.Event, format bool, msg any, args ...any) {
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
	if len(keyvals) > 0 {
		fields := make(map[string]any)
		for i := 0; i < len(keyvals); i = i + 2 {
			fields[fmt.Sprintf("%v", keyvals[i])] = keyvals[i+1]
		}
		return fields
	}
	return nil
}

type zFactory struct{}

func (f *zFactory) With(keyvals ...any) logger.Logger {
	l := newContext().Fields(getFields(keyvals)).Logger()
	return &zLog{&l}
}

func (f *zFactory) WithPrefix(prefix string) logger.Logger {
	return f.With("prefix", prefix)
}

func (f *zFactory) Default() logger.Logger {
	l := newContext().Logger()
	return &zLog{&l}
}

func newContext() zerolog.Context {
	return log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With()
}

func init() {
	logger.SetFactory(&zFactory{})
}
