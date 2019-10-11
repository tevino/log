package log

// PrintLogger represents a logger with Print* APIs.
type PrintLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	PrintDepth(int, ...interface{})
	PrintfDepth(int, string, ...interface{})
	PrintlnDepth(int, ...interface{})
}

// DebugLogger represents a logger with Debug* APIs.
type DebugLogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	DebugDepth(int, ...interface{})
	DebugfDepth(int, string, ...interface{})
}

// InfoLogger represents a logger with Info* APIs.
type InfoLogger interface {
	Info(...interface{})
	Infof(string, ...interface{})

	InfoDepth(int, ...interface{})
	InfofDepth(int, string, ...interface{})
}

// WarnLogger represents a logger with Warn* APIs.
type WarnLogger interface {
	Warn(...interface{})
	Warnf(string, ...interface{})

	WarnDepth(int, ...interface{})
	WarnfDepth(int, string, ...interface{})
}

// WarnLogger represents a logger with Error* APIs.
type ErrorLogger interface {
	Error(...interface{})
	Errorf(string, ...interface{})

	ErrorDepth(int, ...interface{})
	ErrorfDepth(int, string, ...interface{})
}

// FatalLogger represents a logger with Fatal* APIs.
type FatalLogger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})

	FatalDepth(int, ...interface{})
	FatalfDepth(int, string, ...interface{})
}

// Leveler contains level-related APIs.
type Leveler interface {
	DefaultLevel() Level
	SetDefaultLevel(Level)

	OutputLevel() Level
	SetOutputLevel(Level)
}

// CallerOffsetter provides the ability of setting caller offset.
type CallerOffsetter interface {
	SetCallerOffset(int)
}

// Logger represents a full-featured logger.
type Logger interface {
	DebugLogger
	PrintLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	FatalLogger

	Leveler
	CallerOffsetter
}
