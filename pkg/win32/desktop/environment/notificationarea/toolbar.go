// +build windows
package notificationarea

import (
	"fmt"
	"syscall"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

const (
	TOOLBAR_WINDOW32_CLASS string = "ToolbarWindow32"
)

// The notification area is composed
func findToolbars() ([]syscall.Handle, error) {
	handler, _ := windows.FindWindowByClass(NOTIFICATION_AREA_VISIBLE_WINDOW_CLASS)
	return findElementsbyClass(handler, TOOLBAR_WINDOW32_CLASS)
}

// https://www.codeproject.com/Articles/192/Finding-the-position-and-dimensions-of-the-Windows
func findElementsbyClass(hwndParent syscall.Handle, class string) ([]syscall.Handle, error) {
	var hwnds []syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		elementClassName, err := win32api.GetClassName(h)
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		// fmt.Printf("looking for child elements got: %s\n", elementClassName)
		if elementClassName == class {
			// note the window
			// hwnd = h
			hwnds = append(hwnds, h)
			// return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	win32api.EnumChildWindows(hwndParent, cb, 0)
	printHandlers(hwnds)
	if len(hwnds) == 0 {
		return hwnds, fmt.Errorf("No child element on %s with classname %s\n", NOTIFICATION_AREA_VISIBLE_WINDOW_CLASS, TOOLBAR_WINDOW32_CLASS)
	}
	return hwnds, nil
}

func findElementByTitle(hwndParent syscall.Handle, title string) (syscall.Handle, int32, error) {
	var hwnd syscall.Handle
	var elementIndex int32
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := win32api.GetWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			// ignore the error
			elementIndex++
			return 1 // continue enumeration
		}
		elementTitle := syscall.UTF16ToString(b)
		fmt.Printf("looking for child elements got: %s\n", elementTitle)
		if elementTitle == title {
			// note the window
			hwnd = h
			return 0 // stop enumeration
		}
		elementIndex++
		return 1 // continue enumeration
	})
	win32api.EnumChildWindows(hwndParent, cb, 0)
	if hwnd == 0 {
		fmt.Printf("Error the expected element with title %s\n", title)
		return 0, 0, fmt.Errorf("No window with title '%s' found", title)
	}
	return hwnd, elementIndex, nil
}

func printHandlers(hwnds []syscall.Handle) {
	fmt.Printf("len=%d cap=%d %v\n", len(hwnds), cap(hwnds), hwnds)
}
