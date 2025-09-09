//go:build linux || darwin || freebsd || openbsd || netbsd || dragonfly
// +build linux darwin freebsd openbsd netbsd dragonfly

package log

import (
	"fmt"
	"io"
	"os"
	"syscall"

	"golang.org/x/term"
)

// Colors.
const (
	ColorBlue    = "\x1b[0;34m"
	ColorGreen   = "\x1b[0;32m"
	ColorYellow  = "\x1b[0;33m"
	ColorRed     = "\x1b[0;31m"
	ColorMegenta = "\x1b[0;35m"

	ColorRST = "\x1b[0;m"
)

func paint(str string, color string) string {
	return fmt.Sprintf("%s%s%s", color, str, ColorRST)
}

// IsColoredTerminal returns true if the given writer supports colored output.
func IsColoredTerminal(w io.Writer) bool {
	var fd int
	switch w {
	case os.Stdout:
		fd = syscall.Stdout
	case os.Stderr:
		fd = syscall.Stderr
	default:
		return false
	}
	// NOTE: modern terminals support ANSI escape codes.
	return term.IsTerminal(fd)
}
