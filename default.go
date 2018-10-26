package log

import "os"

var defaultLogger Logger

func init() {
	defaultLogger = NewLogger(os.Stdout, LstdFlags|Lshortfile)
	defaultLogger.SetCallerOffset(1)
}

// DefaultLogger returns the default logger which writes to os.Stdout.
func DefaultLogger() Logger {
	return defaultLogger
}

// Print calls the same method on the default logger.
func Print(a ...interface{}) {
	defaultLogger.Print(a...)
}

// Printf calls the same method on the default logger.
func Printf(f string, a ...interface{}) {
	defaultLogger.Printf(f, a...)
}

// Println calls the same method on the default logger.
func Println(a ...interface{}) {
	defaultLogger.Println(a...)
}

// Debug calls the same method on the default logger.
func Debug(a ...interface{}) {
	defaultLogger.Debug(a...)
}

// Debugf calls the same method on the default logger.
func Debugf(f string, a ...interface{}) {
	defaultLogger.Debugf(f, a...)
}

// Info calls the same method on the default logger.
func Info(a ...interface{}) {
	defaultLogger.Info(a...)
}

// Infof calls the same method on the default logger.
func Infof(f string, a ...interface{}) {
	defaultLogger.Infof(f, a...)
}

// Warn calls the same method on the default logger.
func Warn(a ...interface{}) {
	defaultLogger.Warn(a...)
}

// Warnf calls the same method on the default logger.
func Warnf(f string, a ...interface{}) {
	defaultLogger.Warnf(f, a...)
}

// Error calls the same method on the default logger.
func Error(a ...interface{}) {
	defaultLogger.Warn(a...)
}

// Errorf calls the same method on the default logger.
func Errorf(f string, a ...interface{}) {
	defaultLogger.Warnf(f, a...)
}

// Fatal calls the same method on the default logger.
func Fatal(a ...interface{}) {
	defaultLogger.Fatal(a...)
}

// Fatalf calls the same method on the default logger.
func Fatalf(f string, a ...interface{}) {
	defaultLogger.Fatalf(f, a...)
}

// DefaultLevel calls the same method on the default logger.
func DefaultLevel() Level {
	return defaultLogger.DefaultLevel()
}

// SetDefaultLevel calls the same method on the default logger.
func SetDefaultLevel(l Level) {
	defaultLogger.SetDefaultLevel(l)
}

// OutputLevel calls the same method on the default logger.
func OutputLevel() Level {
	return defaultLogger.OutputLevel()
}

// SetOutputLevel calls the same method on the default logger.
func SetOutputLevel(l Level) {
	defaultLogger.SetOutputLevel(l)
}

// SetCallerOffset calls the same method on the default logger.
func SetCallerOffset(offset int) {
	defaultLogger.SetCallerOffset(offset)
}
