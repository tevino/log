package log

import "io"

// NewLogger creates a Logger with given output and flag.
// The flag argument defines the logging properties, it is the same as log.New.
// You could use flags from THIS log package to avoid importing standard package log.
func NewLogger(out io.Writer, flag int) Logger {
	return NewLeveledLogger(out, flag)
}

// NewLoggerWithColor is NewLogger with an additional colored parameter indicating if color is forced.
func NewLoggerWithColor(out io.Writer, flag int, colored bool) Logger {
	return NewLeveledLoggerWithColor(out, flag, colored)
}
