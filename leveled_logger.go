package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
)

// NewLeveledLogger creates a LeveledLogger with given output writer and flag.
// The flag argument defines the logging properties, it is the same as log.New.
// You could use flags from THIS log package to avoid importing standard package log.
func NewLeveledLogger(out io.Writer, flag int) *LeveledLogger {
	return NewLeveledLoggerWithColor(out, flag, IsTerminal(out))
}

// NewLeveledLoggerWithColor is NewLeveledLogger with an additional colored parameter indicating if color is forced.
func NewLeveledLoggerWithColor(out io.Writer, flag int, colored bool) *LeveledLogger {
	return &LeveledLogger{
		debug:        log.New(out, tryPaint("D ", ColorBlue, colored), flag),
		info:         log.New(out, tryPaint("I ", ColorGreen, colored), flag),
		warn:         log.New(out, tryPaint("W ", ColorYellow, colored), flag),
		fata:         log.New(out, tryPaint("F ", ColorRed, colored), flag),
		defaultLevel: INFO,
		outputLevel:  NOTSET,
		depth:        2,
	}
}

func tryPaint(str string, color string, colored bool) string {
	if !colored {
		return str
	}
	return paint(str, color)
}

// LeveledLogger has the ability of logging with different levels.
type LeveledLogger struct {
	debug        *log.Logger
	info         *log.Logger
	warn         *log.Logger
	fata         *log.Logger
	outputLevel  Level
	defaultLevel Level
	depth        int
}

// SetDefaultLevel sets the DefaultLevel atomically.
func (l *LeveledLogger) SetDefaultLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.defaultLevel), int32(level))
}

// DefaultLevel is the level used by Print* methods.
func (l *LeveledLogger) DefaultLevel() Level {
	return Level(atomic.LoadInt32((*int32)(&l.defaultLevel)))
}

// SetOutputLevel sets the OutputLevel atomically.
func (l *LeveledLogger) SetOutputLevel(level Level) {
	atomic.StoreInt32((*int32)(&l.outputLevel), int32(level))
}

// OutputLevel returns the minimal Level of log that will be outputted.
// Levels lower than this will be ignored.
func (l *LeveledLogger) OutputLevel() Level {
	return Level(atomic.LoadInt32((*int32)(&l.outputLevel)))
}

// SetCallerOffset sets the offset used in runtime.Caller(2 + offset)
// while getting file name and line number.
// NOTE: Do not call this while logging, it's not goroutine safe.
func (l *LeveledLogger) SetCallerOffset(offset int) {
	l.depth = offset + 2
}

// Print prints log with DefaultLevel.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Print(a ...interface{}) {
	if l.DefaultLevel() >= l.OutputLevel() {
		l.info.Output(l.depth, fmt.Sprint(a...))
	}
}

// Printf prints log with DefaultLevel.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Printf(format string, a ...interface{}) {
	if l.DefaultLevel() >= l.OutputLevel() {
		l.info.Output(l.depth, fmt.Sprintf(format, a...))
	}
}

// Println prints log with DefaultLevel.
// Arguments are handled in the manner of fmt.Println.
func (l *LeveledLogger) Println(a ...interface{}) {
	if l.DefaultLevel() >= l.OutputLevel() {
		l.info.Output(l.depth, fmt.Sprintln(a...))
	}
}

// Debug prints log with level DEBUG.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Debug(a ...interface{}) {
	if DEBUG >= l.OutputLevel() {
		l.debug.Output(l.depth, fmt.Sprint(a...))
	}
}

// Debugf prints log with level DEBUG.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Debugf(format string, a ...interface{}) {
	if DEBUG >= l.OutputLevel() {
		l.debug.Output(l.depth, fmt.Sprintf(format, a...))
	}
}

// Info prints log with level INFO.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Info(a ...interface{}) {
	if INFO >= l.OutputLevel() {
		l.info.Output(l.depth, fmt.Sprint(a...))
	}
}

// Infof prints log with level INFO.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Infof(format string, a ...interface{}) {
	if INFO >= l.OutputLevel() {
		l.info.Output(l.depth, fmt.Sprintf(format, a...))
	}
}

// Warn prints log with level WARN.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Warn(a ...interface{}) {
	if WARN >= l.OutputLevel() {
		l.warn.Output(l.depth, fmt.Sprint(a...))
	}
}

// Warnf prints log with level WARN.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Warnf(format string, a ...interface{}) {
	if WARN >= l.OutputLevel() {
		l.warn.Output(l.depth, fmt.Sprintf(format, a...))
	}
}

// Fatal prints log with level FATA then os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Fatal(a ...interface{}) {
	if FATA >= l.OutputLevel() {
		l.fata.Output(l.depth, fmt.Sprint(a...))
	}
	os.Exit(1)
}

// Fatalf prints log with level FATA then os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Fatalf(format string, a ...interface{}) {
	if FATA >= l.OutputLevel() {
		l.fata.Output(l.depth, fmt.Sprintf(format, a...))
	}
	os.Exit(1)
}
