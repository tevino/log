package log

// PrintLogger represents a logger with Print* APIs.
type PrintLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

// DebugLogger represents a logger with Debug* APIs.
type DebugLogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}

// InfoLogger represents a logger with Info* APIs.
type InfoLogger interface {
	Info(...interface{})
	Infof(string, ...interface{})
}

// WarnLogger represents a logger with Warn* APIs.
type WarnLogger interface {
	Warn(...interface{})
	Warnf(string, ...interface{})
}

// FatalLogger represents a logger with Fatal* APIs.
type FatalLogger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// Leveler contains level-related APIs.
type Leveler interface {
	DefaultLevel() Level
	SetDefaultLevel(Level)

	OutputLevel() Level
	SetOutputLevel(Level)
}

// Logger represents a full-featured logger.
type Logger interface {
	DebugLogger
	PrintLogger
	InfoLogger
	WarnLogger
	FatalLogger

	Leveler
}
