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
		erro:         log.New(out, tryPaint("E ", ColorMegenta, colored), flag),
		fata:         log.New(out, tryPaint("F ", ColorRed, colored), flag),
		defaultLevel: INFO,
		outputLevel:  NOTSET,
		depth:        3,
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
	erro         *log.Logger
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

// SetCallerOffset sets the offset used in runtime.Caller(3 + offset)
// while getting file name and line number.
// NOTE: Do not call this while logging, it's not goroutine safe.
func (l *LeveledLogger) SetCallerOffset(offset int) {
	l.depth = offset + 3
}

// Print prints log with DefaultLevel.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Print(a ...interface{}) {
	l.output(l.DefaultLevel(), a...)
}

// Printf prints log with DefaultLevel.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Printf(format string, a ...interface{}) {
	l.outputf(l.DefaultLevel(), format, a...)
}

// Println prints log with DefaultLevel.
// Arguments are handled in the manner of fmt.Println.
func (l *LeveledLogger) Println(a ...interface{}) {
	l.outputln(l.DefaultLevel(), a...)
}

// PrintDepth acts as Print but uses depth to determine which call frame to log
// PrintDepth(0, "msg") is the same as Print("msg")
func (l *LeveledLogger) PrintDepth(depth int, a ...interface{}) {
	l.outputDepth(depth, l.DefaultLevel(), a...)
}

// PrintfDepth acts as Printf but uses depth to determine which call frame to log
// PrintfDepth(0, "msg") is the same as Printf("msg")
func (l *LeveledLogger) PrintfDepth(depth int, format string, a ...interface{}) {
	l.outputfDepth(depth, l.DefaultLevel(), format, a...)
}

// PrintlnDepth acts as Printfln but uses depth to determine which call frame to log
// PrintflnDepth(0, "msg") is the same as Printfln("msg")
func (l *LeveledLogger) PrintlnDepth(depth int, a ...interface{}) {
	l.outputlnDepth(depth, l.DefaultLevel(), a...)
}

// Debug prints log with level DEBUG.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Debug(a ...interface{}) {
	l.output(DEBUG, a...)
}

// Debugf prints log with level DEBUG.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Debugf(format string, a ...interface{}) {
	l.outputf(DEBUG, format, a...)
}

// DebugDepth acts as Debug but uses depth to determine which call frame to log
// DebugDepth(0, "msg") is the same as Debug("msg")
func (l *LeveledLogger) DebugDepth(depth int, a ...interface{}) {
	l.outputDepth(depth, DEBUG, a...)
}

// DebugfDepth acts as Debugf but uses depth to determine which call frame to log
// DebugfDepth(0, "msg") is the same as Debugf("msg")
func (l *LeveledLogger) DebugfDepth(depth int, format string, a ...interface{}) {
	l.outputfDepth(depth, DEBUG, format, a...)
}

// Info prints log with level INFO.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Info(a ...interface{}) {
	l.output(INFO, a...)
}

// Infof prints log with level INFO.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Infof(format string, a ...interface{}) {
	l.outputf(INFO, format, a...)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log
// InfoDepth(0, "msg") is the same as Info("msg")
func (l *LeveledLogger) InfoDepth(depth int, a ...interface{}) {
	l.outputDepth(depth, INFO, a...)
}

// InfofDepth acts as Infof but uses depth to determine which call frame to log
// InfofDepth(0, "msg") is the same as Infof("msg")
func (l *LeveledLogger) InfofDepth(depth int, format string, a ...interface{}) {
	l.outputfDepth(depth, INFO, format, a...)
}

// Warn prints log with level WARN.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Warn(a ...interface{}) {
	l.output(WARN, a...)
}

// Warnf prints log with level WARN.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Warnf(format string, a ...interface{}) {
	l.outputf(WARN, format, a...)
}

// WarnDepth acts as Warn but uses depth to determine which call frame to log
// WarnDepth(0, "msg") is the same as Warn("msg")
func (l *LeveledLogger) WarnDepth(depth int, a ...interface{}) {
	l.outputDepth(depth, WARN, a...)
}

// WarnfDepth acts as Warnf but uses depth to determine which call frame to log
// WarnfDepth(0, "msg") is the same as Warnf("msg")
func (l *LeveledLogger) WarnfDepth(depth int, format string, a ...interface{}) {
	l.outputfDepth(depth, WARN, format, a...)
}

// Error prints log with level ERROR.
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Error(a ...interface{}) {
	l.output(ERROR, a...)
}

// Errorf prints log with level ERROR.
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Errorf(format string, a ...interface{}) {
	l.outputf(ERROR, format, a...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log
// ErrorDepth(0, "msg") is the same as Error("msg")
func (l *LeveledLogger) ErrorDepth(depth int, a ...interface{}) {
	l.outputDepth(depth, ERROR, a...)
}

// ErrorfDepth acts as Errorf but uses depth to determine which call frame to log
// ErrorfDepth(0, "msg") is the same as Errorf("msg")
func (l *LeveledLogger) ErrorfDepth(depth int, format string, a ...interface{}) {
	l.outputfDepth(depth, ERROR, format, a...)
}

// Fatal prints log with level FATA then os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *LeveledLogger) Fatal(a ...interface{}) {
	l.output(FATA, a...)
}

// Fatalf prints log with level FATA then os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *LeveledLogger) Fatalf(format string, a ...interface{}) {
	l.outputf(FATA, format, a...)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log
// FatalDepth(0, "msg") is the same as Fatal("msg")
func (l *LeveledLogger) FatalDepth(depth int, a ...interface{}) {
	l.outputDepth(depth, FATA, a...)
}

// FatalfDepth acts as Fatalf but uses depth to determine which call frame to log
// FatalfDepth(0, "msg") is the same as Fatalf("msg")
func (l *LeveledLogger) FatalfDepth(depth int, format string, a ...interface{}) {
	l.outputfDepth(depth, FATA, format, a...)
}

func (l *LeveledLogger) output(level Level, a ...interface{}) {
	l.outputDepth(1, level, a...)
}

func (l *LeveledLogger) outputDepth(depth int, level Level, a ...interface{}) {
	if level < l.OutputLevel() {
		return
	}

	logger := l.getOutputTarget(level)
	if logger != nil {
		logger.Output(l.depth+depth, fmt.Sprint(a...))
	}
	if level == FATA {
		os.Exit(1)
	}
}

func (l *LeveledLogger) outputln(level Level, a ...interface{}) {
	l.outputlnDepth(1, level, a...)
}

func (l *LeveledLogger) outputlnDepth(depth int, level Level, a ...interface{}) {
	if level < l.OutputLevel() {
		return
	}

	logger := l.getOutputTarget(level)
	if logger != nil {
		logger.Output(l.depth+depth, fmt.Sprintln(a...))
	}
	if level == FATA {
		os.Exit(1)
	}
}

func (l *LeveledLogger) outputf(level Level, format string, a ...interface{}) {
	l.outputfDepth(1, level, format, a...)
}

func (l *LeveledLogger) outputfDepth(depth int, level Level, format string, a ...interface{}) {
	if level < l.OutputLevel() {
		return
	}

	logger := l.getOutputTarget(level)
	if logger != nil {
		logger.Output(l.depth+depth, fmt.Sprintf(format, a...))
	}
	if level == FATA {
		os.Exit(1)
	}
}

func (l *LeveledLogger) getOutputTarget(level Level) (logger *log.Logger) {
	switch level {
	case DEBUG:
		logger = l.debug
	case INFO:
		logger = l.info
	case WARN:
		logger = l.warn
	case ERROR:
		logger = l.erro
	case FATA:
		logger = l.fata
	}
	return
}
