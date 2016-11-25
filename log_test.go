package log

import "os"

func ExampleLogger() {
	l := NewLoggerWithColor(os.Stdout, 0, false)
	// Color and time-related flags should be disabled in testing.
	// While you may want to use the following instead:
	// l := NewLeveledLogger(os.Stdout, LstdFlags)

	l.SetOutputLevel(INFO)
	l.Debug("Output level is INFO.")
	l.Info("Thus debug is not printed.")

	l.SetOutputLevel(DEBUG)
	l.Debug("The above line enables debug.")
	// Output:
	// I Thus debug is not printed.
	// D The above line enables debug.
}
