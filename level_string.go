// Code generated by "enumer -type=Level -output level_string.go -trimprefix=Level"; DO NOT EDIT.

package logger

import (
	"fmt"
	"strings"
)

const _LevelName = "DebugInfoWarnErrorFatal"

var _LevelIndex = [...]uint8{0, 5, 9, 13, 18, 23}

const _LevelLowerName = "debuginfowarnerrorfatal"

func (i Level) String() string {
	if i >= Level(len(_LevelIndex)-1) {
		return fmt.Sprintf("Level(%d)", i)
	}
	return _LevelName[_LevelIndex[i]:_LevelIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _LevelNoOp() {
	var x [1]struct{}
	_ = x[LevelDebug-(0)]
	_ = x[LevelInfo-(1)]
	_ = x[LevelWarn-(2)]
	_ = x[LevelError-(3)]
	_ = x[LevelFatal-(4)]
}

var _LevelValues = []Level{LevelDebug, LevelInfo, LevelWarn, LevelError, LevelFatal}

var _LevelNameToValueMap = map[string]Level{
	_LevelName[0:5]:        LevelDebug,
	_LevelLowerName[0:5]:   LevelDebug,
	_LevelName[5:9]:        LevelInfo,
	_LevelLowerName[5:9]:   LevelInfo,
	_LevelName[9:13]:       LevelWarn,
	_LevelLowerName[9:13]:  LevelWarn,
	_LevelName[13:18]:      LevelError,
	_LevelLowerName[13:18]: LevelError,
	_LevelName[18:23]:      LevelFatal,
	_LevelLowerName[18:23]: LevelFatal,
}

var _LevelNames = []string{
	_LevelName[0:5],
	_LevelName[5:9],
	_LevelName[9:13],
	_LevelName[13:18],
	_LevelName[18:23],
}

// LevelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func LevelString(s string) (Level, error) {
	if val, ok := _LevelNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _LevelNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Level values", s)
}

// LevelValues returns all values of the enum
func LevelValues() []Level {
	return _LevelValues
}

// LevelStrings returns a slice of all String values of the enum
func LevelStrings() []string {
	strs := make([]string, len(_LevelNames))
	copy(strs, _LevelNames)
	return strs
}

// IsALevel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Level) IsALevel() bool {
	for _, v := range _LevelValues {
		if i == v {
			return true
		}
	}
	return false
}