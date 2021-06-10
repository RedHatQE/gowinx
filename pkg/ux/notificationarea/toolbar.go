// +build windows
package notificationarea

import (
	"fmt"
	"syscall"

	"github.com/adrianriobo/gowinx/pkg/ux"
	"github.com/adrianriobo/gowinx/pkg/win32"
)

const (
	TOOLBAR_WINDOW32_CLASS string = "ToolbarWindow32"
)

func findToolbars() ([]syscall.Handle, error) {
	handler, _ := ux.FindWindowByClass(NOTIFICATION_AREA_VISIBLE_WINDOW_CLASS)
	return findElementsbyClass(handler, TOOLBAR_WINDOW32_CLASS)
}

// https://www.codeproject.com/Articles/192/Finding-the-position-and-dimensions-of-the-Windows
func findElementsbyClass(hwndParent syscall.Handle, class string) ([]syscall.Handle, error) {
	var hwnds []syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		elementClassName, err := win32.GetClassName(h)
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		fmt.Printf("looking for child elements got: %s\n", elementClassName)
		if elementClassName == class {
			// note the window
			// hwnd = h
			hwnds = append(hwnds, h)
			// return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	win32.EnumChildWindows(hwndParent, cb, 0)
	printHandlers(hwnds)
	if len(hwnds) == 0 {
		return hwnds, fmt.Errorf("No child element on %s with classname %s\n", NOTIFICATION_AREA_VISIBLE_WINDOW_CLASS, TOOLBAR_WINDOW32_CLASS)
	}
	return hwnds, nil
}

func findElementByTitle(hwndParent syscall.Handle, title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := win32.GetWindowText(h, &b[0], int32(len(b)))
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
	win32.EnumChildWindows(hwndParent, cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("No window with title '%s' found", title)
	}
	return hwnd, nil
}

func printHandlers(hwnds []syscall.Handle) {
	fmt.Printf("len=%d cap=%d %v\n", len(hwnds), cap(hwnds), hwnds)
}
