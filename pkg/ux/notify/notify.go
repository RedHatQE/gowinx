package notify

import (
	"fmt"
	"syscall"

	"github.com/lxn/win"
)

const (
	NIOW_CLASS        string = "NotifyIconOverflowWindow"
	NIOW_TOOLBAR32_ID int32  = 1504
)

func GetNotifyIconOverflowWindowHandler() (win.HWND, error) {
	if handler := getWindowHandlerByClass(NIOW_CLASS); handler > 0 {
		return handler, nil
	}
	return win.HWND(0), fmt.Errorf("No handler found for class %s", NIOW_CLASS)
}

func ShowNotifyIconOverflowWindow() error {
	if handler, err := GetNotifyIconOverflowWindowHandler(); err != nil {
		return err
	} else {
		win.ShowWindow(handler, win.SW_SHOWNORMAL)
	}
	return nil
}

func GetNotifyToolbarHandler() (win.HWND, error) {
	if handler, err := GetNotifyIconOverflowWindowHandler(); err != nil {
		return win.HWND(0), err
	} else {
		if toolbarHandler := win.GetDlgItem(handler, NIOW_TOOLBAR32_ID); toolbarHandler > 0 {
			return toolbarHandler, nil
		}
	}
	return win.HWND(0), fmt.Errorf("Error getting NotifyToolbarHandler")
}

func GetButtonsCountONotifyToolbar() (uintptr, error) {
	if handler, err := GetNotifyToolbarHandler(); err != nil {
		return 0, err
	} else {
		return win.SendMessage(handler, win.TB_BUTTONCOUNT, 0, 0), nil
	}
}

func getWindowHandlerByClass(class string) win.HWND {
	z := uint16(0)
	return win.FindWindow(syscall.StringToUTF16Ptr(class), &z)
}
