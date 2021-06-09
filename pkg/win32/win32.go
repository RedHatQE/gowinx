// +build windows
package win32

import (
	"syscall"
)

var (
	kernel32 = syscall.MustLoadDLL("kernel32.dll")
	user32   = syscall.MustLoadDLL("user32.dll")
)
