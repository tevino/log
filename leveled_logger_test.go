package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestFileLine(t *testing.T) {
	var buf bytes.Buffer
	l := NewLeveledLogger(&buf, Lshortfile)
	l.Info("Test file line")

	var exp = "leveled_logger_test.go:12"
	if !strings.Contains(buf.String(), exp) {
		t.Errorf("Expected filename and line number '%s' not found in: '%s'", exp, buf.String())
	}
}

func TestCallerOffset(t *testing.T) {
	var buf bytes.Buffer
	l := NewLeveledLogger(&buf, Lshortfile)
	l.SetCallerOffset(1)
	l.Info("Test file line")

	var exp = "testing.go"
	if !strings.Contains(buf.String(), exp) {
		t.Errorf("Expected filename and line number '%s' not found in: '%s'", exp, buf.String())
	}

	l.SetCallerOffset(0)
	l.Info("Test file line")

	exp = "leveled_logger_test.go:32"
	if !strings.Contains(buf.String(), exp) {
		t.Errorf("Expected filename and line number '%s' not found in: '%s'", exp, buf.String())
	}
}

func TestOutputLevel(t *testing.T) {
	var buf bytes.Buffer
	l := NewLeveledLogger(&buf, LstdFlags)
	l.SetOutputLevel(WARN)
	l.Debug("DEBUG Log")
	l.Info("INFO Log")
	l.Warn("WARN Log")
	if strings.Contains(buf.String(), "DEBUG") {
		t.Errorf("DEBUG log is NOT expected: '%s'", buf.String())
	}
	if strings.Contains(buf.String(), "INFO") {
		t.Errorf("INFO log is NOT expected: '%s'", buf.String())
	}
	if !strings.Contains(buf.String(), "WARN") {
		t.Errorf("WARN log is expected: '%s'", buf.String())
	}
}

func TestDynamicCallDepth(t *testing.T) {
	var buf bytes.Buffer
	l := NewLeveledLogger(&buf, Lshortfile)

	assertFileAndLine := func(t *testing.T, data string, exp string) {
		t.Helper()
		if !strings.Contains(data, exp) {
			t.Errorf("Expected filename and line number '%s' not found in: '%s'", exp, data)
		}
	}

	t.Run("debug", func(t *testing.T) {
		buf.Reset()
		debug := func(a ...interface{}) {
			l.DebugDepth(1, a...)
		}
		debug("DEBUG log")
		exp := "leveled_logger_test.go:74"
		assertFileAndLine(t, buf.String(), exp)
	})

	t.Run("info", func(t *testing.T) {
		buf.Reset()
		info := func(a ...interface{}) {
			l.InfoDepth(1, a...)
		}
		info("INFO log")
		exp := "leveled_logger_test.go:84"
		assertFileAndLine(t, buf.String(), exp)
	})

	t.Run("warn", func(t *testing.T) {
		buf.Reset()
		warn := func(a ...interface{}) {
			l.WarnDepth(1, a...)
		}
		warn("WARN log")
		exp := "leveled_logger_test.go:94"
		assertFileAndLine(t, buf.String(), exp)
	})

	t.Run("error", func(t *testing.T) {
		buf.Reset()
		err := func(a ...interface{}) {
			l.WarnDepth(1, a...)
		}
		err("ERROR log")
		exp := "leveled_logger_test.go:104"
		assertFileAndLine(t, buf.String(), exp)
	})
}
