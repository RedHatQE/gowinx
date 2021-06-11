// +build windows
package notificationarea

import (
	"fmt"
	"syscall"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

// systemtray aka notification area, it is composed of notifications icons (offering display the status and various functions)
// distributerd across:
// * visible area on the right side of the taskbar (class: Shell_TrayWnd)
// * hidden area as overflowwindow ( class: NotifyIconOverflowWindow)

const (
	NOTIFICATION_AREA_VISIBLE_WINDOW_CLASS string = "Shell_TrayWnd"
	NOTIFICATION_AREA_HIDDEN_WINDOW_CLASS  string = "NotifyIconOverflowWindow"
	TOOLBARWINDOWS32_ID                    int32  = 1504
)

func GetHiddenNotificationAreaRect() (rect win32api.RECT, err error) {
	// Show notification area (hidden)
	if err = ShowHiddenNotificationArea(); err == nil {
		if toolbarHandler, err := getNotificationAreaToolbarByWindowClass(NOTIFICATION_AREA_HIDDEN_WINDOW_CLASS); err == nil {
			if _, err = win32api.GetWindowRect(toolbarHandler, &rect); err == nil {
				fmt.Printf("Rect for system tray t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
			}
		}
	}
	if err != nil {
		fmt.Printf("error getting hidden notification area rect: %v\n", err)
	}
	return
}

func ShowHiddenNotificationArea() (err error) {
	if handler, err := getNotificationAreaWindowByClass(NOTIFICATION_AREA_HIDDEN_WINDOW_CLASS); err == nil {
		win32api.ShowWindow(handler, win32api.SW_SHOWNORMAL)
	}
	return
}

func getNotificationAreaWindowByClass(className string) (handler syscall.Handle, err error) {
	if handler, err = windows.FindWindowByClass(className); err != nil {
		fmt.Printf("error getting handler on notification area for windows class: %s, error: %v\n", className, err)
	}
	return
}

func getNotificationAreaToolbarByWindowClass(className string) (handler syscall.Handle, err error) {
	if windowHandler, err := getNotificationAreaWindowByClass(className); err == nil {
		if handler, err = win32api.GetDlgItem(windowHandler, TOOLBARWINDOWS32_ID); err != nil {
			fmt.Printf("error getting toolbar handler on notification area for windows class: %s, error: %v\n", className, err)
		}
	}
	return
}
