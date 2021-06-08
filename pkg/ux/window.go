// +build windows
package ux

import (
	"syscall"

	"github.com/lxn/win"
)

func FinWindowByClassAndTitle(class, title string) win.HWND {
	return win.FindWindow(syscall.StringToUTF16Ptr(class), syscall.StringToUTF16Ptr(title))
}
