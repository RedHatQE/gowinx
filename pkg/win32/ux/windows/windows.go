// +build windows
package windows

import (
	"fmt"
	"syscall"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
)

func FindWindowByTitle(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := win32api.GetWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		if syscall.UTF16ToString(b) == title {
			// note the window
			hwnd = h
			return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	win32api.EnumWindows(cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("No window with title '%s' found", title)
	}
	return hwnd, nil
}

func FindWindowByClass(class string) (syscall.Handle, error) {
	z := uint16(0)
	return win32api.FindWindowW(syscall.StringToUTF16Ptr(class), &z)
}

func FindWindowExByClass(parentHandler syscall.Handle, class string) (syscall.Handle, error) {
	z := uint16(0)
	return win32api.FindWindowEx(parentHandler, syscall.Handle(0), syscall.StringToUTF16Ptr(class), &z)
}
