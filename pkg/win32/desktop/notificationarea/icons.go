// +build windows
package notificationarea

import (
	"fmt"
	"syscall"

	win32wam "github.com/adrianriobo/gowinx/pkg/win32/api/windows-and-messages"
	win32windows "github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

func GetHiddenIconsCount() (int32, error) {
	return getIconsCountByWindowClass(NOTIFICATION_AREA_HIDDEN_WINDOW_CLASS)
}

func getIconsCountByWindowClass(className string) (int32, error) {
	var err error
	if toolbarHandler, err := getNotificationAreaToolbarByWindowClass(className); err == nil {
		buttonsCount, _ := win32wam.SendMessage(toolbarHandler, win32wam.TB_BUTTONCOUNT, 0, 0)
		return int32(buttonsCount), nil
	}
	return 0, err
}

func GetIconPosition(rect win32wam.RECT) (x, y int32) {
	x = rect.Left + 10
	y = rect.Top + 10
	fmt.Printf("Crc icon will be clicked at x: %d y: %d\n", x, y)
	return
}

func GetIconByTittle(title string) syscall.Handle {
	toolbarHandlers, _ := findToolbars()
	for i, toolbarHandler := range toolbarHandlers {
		fmt.Printf("Looking for %s at toolbar index %d\n", title, i)
		iconHandler, iconIndex, err := win32windows.FindChildWindowByTitle(toolbarHandler, title)
		if err == nil {
			fmt.Printf("We found the icon for %s at index %d\n", title, iconIndex)
			return iconHandler
		}
	}
	return 0
}

func GetIconRectByTittle(title string) (rect win32wam.RECT, err error) {
	toolbarHandlers, _ := findToolbars()
	for i, toolbarHandler := range toolbarHandlers {
		fmt.Printf("Looking for %s at toolbar index %d\n", title, i)
		iconHandler, iconIndex, err := win32windows.FindChildWindowByTitle(toolbarHandler, title)
		if err == nil {
			fmt.Printf("We found the icon for %s at index %d\n", title, iconIndex)
			rect, err = getControlRect(iconHandler)
		}
	}
	return
}

func getControlRect(controlHandler syscall.Handle) (rect win32wam.RECT, err error) {
	if _, err = win32wam.GetWindowRect(controlHandler, &rect); err == nil {
		fmt.Printf("Rect for control t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
	} else {
		fmt.Printf("error getting control area rect: %v\n", err)
	}
	return
}
