// +build linux

package log

import "syscall"

const ioCtlReadTermios = syscall.TCGETS
