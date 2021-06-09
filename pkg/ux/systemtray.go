// +build windows
package ux

import (
	"fmt"
	"syscall"

	"github.com/adrianriobo/gowinx/pkg/win32"
)

const (
	NIOW_CLASS        string = "NotifyIconOverflowWindow"
	NIOW_TOOLBAR32_ID int32  = 1504
)

func GetNotifyIconOverflowWindowHandler() (syscall.Handle, error) {
	if handler, err := FindWindowByClass(NIOW_CLASS); err == nil {
		return handler, nil
	} else {
		return syscall.Handle(0), fmt.Errorf("No handler found for class %s\n", NIOW_CLASS)
	}
}

func GetNotifyToolbarHandler() (toolbarHandler syscall.Handle, err error) {
	if handler, err := GetNotifyIconOverflowWindowHandler(); err == nil {
		toolbarHandler, err = win32.GetDlgItem(handler, NIOW_TOOLBAR32_ID)
	}
	return
}

func GetButtonsCountONotifyToolbar() (buttonsCount uintptr, err error) {
	if handler, err := GetNotifyToolbarHandler(); err == nil {
		buttonsCount, err = win32.SendMessageW(handler, win32.TB_BUTTONCOUNT, 0, 0)
	}
	return
}

func GetNotifyAreaRect() (rect win32.RECT, err error) {
	// Show notification area (hidden)
	ShowNotifyIconOverflowWindow()
	if toolbarHandler, err := GetNotifyToolbarHandler(); err == nil {
		if _, err = win32.GetWindowRect(toolbarHandler, &rect); err == nil {
			fmt.Printf("Rect for system tray t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
		}
	}
	if err != nil {
		fmt.Printf("error getting notification area rect: %v\n", err)
	}
	return
}

func ShowNotifyIconOverflowWindow() (err error) {
	if handler, err := GetNotifyIconOverflowWindowHandler(); err == nil {
		win32.ShowWindow(handler, win32.SW_SHOWNORMAL)
	}
	return
}

// To implement
// func GetIconPosition(title string) (x, y int32) {
// }

func GetIconPosition(rect win32.RECT) (x, y int32) {
	x = rect.Left + 10
	y = rect.Top + 10
	fmt.Printf("Crc icon will be clicked at x: %d y: %d\n", x, y)
	return
}
