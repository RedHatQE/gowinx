// +build windows
package notificationarea

import (
	"fmt"
	"syscall"

	"github.com/adrianriobo/gowinx/pkg/win32"
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
		buttonsCount, _ := win32.SendMessage(toolbarHandler, win32.TB_BUTTONCOUNT, 0, 0)
		return int32(buttonsCount), nil
	}
	return 0, err
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

func GetIconByTittle(title string) syscall.Handle {
	toolbarHandlers, _ := findToolbars()
	for _, toolbarHandler := range toolbarHandlers {
		iconHandler, err := findElementByTitle(toolbarHandler, title)
		if err != nil {
			fmt.Printf("We found the icon for %s\n", title)
			return iconHandler
		}
	}
	return 0
}
