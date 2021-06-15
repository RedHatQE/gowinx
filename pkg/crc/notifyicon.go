// +build windows
package crc

import (
	"fmt"

	win32wam "github.com/adrianriobo/gowinx/pkg/win32/api/windows-and-messages"
	win32windows "github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

const (
	CONTEXT_MENU_TITLE string = "crcText"

	CONTEXT_MENU_MARGIN_TOP            int32 = 2
	CONTEXT_MENU_ITEM_SEPARATOR_HEIGHT int32 = 6
	CONTEXT_MENU_ITEM_HEIGHT           int32 = 22
	CONTEXT_MENU_ITEM_WIDTH            int32 = 215

	CONTEXT_MENU_ITEM_SEPARATOR       string = "separator"
	CONTEXT_MENU_ITEM_STATUS          string = "status"
	CONTEXT_MENU_ITEM_STATUS_AND_LOGS string = "status-and-logs"
	CONTEXT_MENU_ITEM_START           string = "start"
	CONTEXT_MENU_ITEM_STOP            string = "stop"
	CONTEXT_MENU_ITEM_DELETE          string = "delete"
)

type contextMenuItem struct {
	Height int32
	Name   string
}

func separator() contextMenuItem {
	return contextMenuItem{Height: CONTEXT_MENU_ITEM_SEPARATOR_HEIGHT, Name: CONTEXT_MENU_ITEM_SEPARATOR}
}

func menuItem(menuItemName string) contextMenuItem {
	return contextMenuItem{Height: CONTEXT_MENU_ITEM_HEIGHT, Name: menuItemName}
}

var (
	contextMenu = [13]contextMenuItem{
		menuItem(CONTEXT_MENU_ITEM_STATUS),
		separator(),
		menuItem(CONTEXT_MENU_ITEM_STATUS_AND_LOGS),
		separator(),
		menuItem(CONTEXT_MENU_ITEM_START),
		menuItem(CONTEXT_MENU_ITEM_STOP),
		menuItem(CONTEXT_MENU_ITEM_DELETE)}
)

func MenuItemPosition(menuItemName string) (x, y int32) {
	if iconMenuRect, err := iconMenuRect(); err == nil {
		x, y = menuItemRelativePosition(menuItemName)
		x = x + iconMenuRect.Left
		y = y + iconMenuRect.Top
	}
	fmt.Printf("Get button on menu %s X:%d y:%d\n", menuItemName, x, y)
	return
}

func iconMenuRect() (rect win32wam.RECT, err error) {
	winHWND, err := win32windows.FindWindowByTitle(CONTEXT_MENU_TITLE)
	if err == nil {
		if _, err = win32wam.GetWindowRect(winHWND, &rect); err == nil {
			fmt.Printf("Rect for system tray t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
			return
		}
	}
	fmt.Print("error getting icon menu window handler")
	return
}

func menuItemRelativePosition(menuItemName string) (x, y int32) {
	x = int32(CONTEXT_MENU_ITEM_WIDTH) / 2
	y = int32(CONTEXT_MENU_MARGIN_TOP)
	for _, menuitem := range contextMenu {
		if menuitem.Name == menuItemName {
			y = y + (menuitem.Height / 2)
			break
		}
		y = y + menuitem.Height
	}
	fmt.Printf("Get button relative %s X:%d y:%d\n", menuItemName, x, y)
	return
}
