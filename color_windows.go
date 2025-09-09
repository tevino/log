//go:build windows
// +build windows

package log

import (
	"fmt"
	"io"
)

// Dummy colors.
const (
	ColorBlue    = ""
	ColorGreen   = ""
	ColorYellow  = ""
	ColorRed     = ""
	ColorMegenta = ""

	ColorRST = ""
)

func paint(str string, color string) string {
	return fmt.Sprintf("%s%s%s", color, str, ColorRST)
}

// IsColoredTerminal returns true.
func IsColoredTerminal(_ io.Writer) bool {
	// TODO: color support for Windows.
	return false
}
