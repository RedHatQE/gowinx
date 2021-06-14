// +build windows
package notificationarea

import (
	"fmt"
	"syscall"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
	win32windows "github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

func GetHiddenIconsCount() (int32, error) {
	return getIconsCountByWindowClass(NOTIFICATION_AREA_HIDDEN_WINDOW_CLASS)
}

// FIXME how to identify ToolBar32 always instance 3?
// func GetVisibleIconsCount() (int32, error) {
// 	handler, _ := ux.FindWindowByClass(NOTIFICATION_AREA_VISIBLE_WINDOW_CLASS)
// 	toolbarHandler, _ := findElementsbyClass(handler, "ToolbarWindow32")
// 	buttonsCount, _ := win32.SendMessage(toolbarHandler, win32.TB_BUTTONCOUNT, 0, 0)
// 	return int32(buttonsCount), nil
// }

func getIconsCountByWindowClass(className string) (int32, error) {
	var err error
	if toolbarHandler, err := getNotificationAreaToolbarByWindowClass(className); err == nil {
		buttonsCount, _ := win32api.SendMessage(toolbarHandler, win32api.TB_BUTTONCOUNT, 0, 0)
		return int32(buttonsCount), nil
	}
	return 0, err
}

// To implement
// func GetIconPosition(title string) (x, y int32) {
// }

func GetIconPosition(rect win32api.RECT) (x, y int32) {
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

func GetIconRectByTittle(title string) (rect win32api.RECT, err error) {
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

func getControlRect(controlHandler syscall.Handle) (rect win32api.RECT, err error) {
	if _, err = win32api.GetWindowRect(controlHandler, &rect); err == nil {
		fmt.Printf("Rect for control t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
	} else {
		fmt.Printf("error getting control area rect: %v\n", err)
	}
	return
}

// func GetButtonsTexts() {
// 	// var err error
// 	if toolbarHandler, err := getNotificationAreaToolbarByWindowClass(NOTIFICATION_AREA_HIDDEN_WINDOW_CLASS); err == nil {
// 		buttonsCount, _ := win32api.SendMessage(toolbarHandler, win32api.TB_BUTTONCOUNT, 0, 0)
// 		for i := 0; i < int(buttonsCount); i++ {
// 			text, _ := win32toolbar.GetButtonText(toolbarHandler, i)
// 			index := win32toolbar.GetButtonIndex(toolbarHandler, i)
// 			fmt.Printf("The name of the button at index %d, is %s\n", index, text)
// 		}
// 	}
// }
