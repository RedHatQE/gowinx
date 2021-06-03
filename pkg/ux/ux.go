package ux

import (
	"syscall"

	"github.com/lxn/win"
)

func GetWindowHandlerByClass(class string) win.HWND {
	z := uint16(0)
	return win.FindWindow(syscall.StringToUTF16Ptr(class), &z)
}
