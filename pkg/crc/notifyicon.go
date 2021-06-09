// +build windows
package crc

import (
	"fmt"

	"github.com/adrianriobo/gowinx/pkg/ux"
	"github.com/adrianriobo/gowinx/pkg/win32"
)

const (
	// OS dependant
	// https://stackoverflow.com/questions/19436860/go-golang-trying-to-get-window-information-via-syscall-as-in-enumwindows-etc
	// https://gist.github.com/EliCDavis/5374fa4947897b16a81f6550d142ab28
	CONTEXT_MENU_CLASS string = "WindowsForms10.Window.20808.app.0.232467a_r7_ad1"
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
	return
}

func iconMenuRect() (rect win32.RECT, err error) {
	// if winHWND := ux.FinWindowByClassAndTitle(CONTEXT_MENU_CLASS, CONTEXT_MENU_TITLE); winHWND > 0 {
	if winHWND, err := ux.FindWindowByTitle(CONTEXT_MENU_TITLE); err == nil {
		if _, err = win32.GetWindowRect(winHWND, &rect); err == nil {
			fmt.Printf("Rect for system tray t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
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
			fmt.Printf("Get button %s X coord at %d\n", menuItemName, x)
			break
		}
		y = y + menuitem.Height
	}
	return
}

// Give a try to click directly sending messages
// func ClickMenuItem(position int) {
// 	// if winHWND := ux.FinWindowByClassAndTitle(CONTEXT_MENU_CLASS, CONTEXT_MENU_TITLE); winHWND > 0 {
// 	if winHWND, err := ux.FindWindowByTitle(CONTEXT_MENU_TITLE); err == nil {
// 		// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu#remarks
// 		if menuHandler := windows.GetMenu(winHWND); menuHandler > 0 {
// 			if menuItemID := win.GetMenuItemID(menuHandler, int32(position)); menuItemID > 0 {
// 				fmt.Printf("We got menu item ID %d", menuItemID)
// 				win.SendMessage(winHWND, win.WM_COMMAND, windows.MakeLPARAM(0, uint16(menuItemID)), 0)
// 			}
// 		}

// 	}
// }
