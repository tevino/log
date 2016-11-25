package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestFileLine(t *testing.T) {
	var buf bytes.Buffer
	l := NewLeveledLogger(&buf, LstdFlags)
	l.Info("Test file line")

	var exp = "leveled_logger_test.go:21"
	if strings.Contains(buf.String(), exp) {
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
