package log

import "strings"

// Log levels from low to high.
// NOTE: FATAL is the highest, NOTSET is the lowest.
const (
	NOTSET Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATA
)

// LevelFromString parses string to Level.
// NOTSET is returned if the given string can not be recognized.
func LevelFromString(s string) Level {
	switch strings.TrimSpace(strings.ToUpper(s)) {
	case "DEBUG", "D":
		return DEBUG
	case "INFO", "I":
		return INFO
	case "WARN", "WARNING", "W":
		return WARN
	case "ERROR", "E":
		return ERROR
	case "FATA", "FATAL", "F":
		return FATA
	case "NOTSET", "NOT SET", "N":
		fallthrough
	default:
		return NOTSET
	}
}

// Level represents level of logging.
type Level int32

// String returns the string representation of Level.
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATA:
		return "FATA"
	case NOTSET:
		fallthrough
	default:
		return "NOT SET"
	}
}
