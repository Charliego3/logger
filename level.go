package logger

//go:generate enumer -type=Level -output level_string.go -trimprefix=Level
type Level uint

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)
