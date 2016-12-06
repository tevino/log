// +build darwin freebsd openbsd netbsd dragonfly

package log

import "syscall"

const ioCtlReadTermios = syscall.TIOCGETA
