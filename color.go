package log

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"unsafe"
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

// IsTerminal returns true if the given writer supports colored output.
func IsTerminal(w io.Writer) bool {
	var fd int
	switch w {
	case os.Stdout:
		fd = syscall.Stdout
	case os.Stderr:
		fd = syscall.Stderr
	default:
		return false
	}
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioCtlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
