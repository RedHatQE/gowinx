// +build windows
package crc

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/adrianriobo/gowinx/pkg/ux"
	"github.com/adrianriobo/gowinx/pkg/windows"
	"github.com/lxn/win"
)

const (
	CONTEXT_MENU_CLASS string = "WindowsForms10.Window.20808.app.0.232467a_r7_ad1"
	CONTEXT_MENU_TITLE string = "crcText"

	// Check if resolution affects
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
	iconMenuRect := iconMenuRect()
	x, y = menuItemRelativePosition(menuItemName)
	x = x + iconMenuRect.Left
	y = y + iconMenuRect.Top
	return
}

func iconMenuRect() win.RECT {
	winHWND := ux.FinWindowByClassAndTitle(CONTEXT_MENU_CLASS, CONTEXT_MENU_TITLE)
	if winHWND > 0 {
		var rect win.RECT
		if win.GetWindowRect(winHWND, &rect) {
			fmt.Printf("Rect for CRC icon menu t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
			return rect
		}
	}
	return win.RECT{}
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

// Deprecated ??
func buttonsCountOnTray() int {
	if handler := ux.FinWindowByClassAndTitle(CONTEXT_MENU_CLASS, CONTEXT_MENU_TITLE); handler > 0 {
		number := int(win.SendMessage(handler, win.TB_BUTTONCOUNT, 0, 0))
		fmt.Printf("Got the handler and number buttons is %d \n", number)
		return number
	} else {
		return 0
	}
}

// Deprecated ??
func clickButtonsOnTray() {
	if toolbarHandler, err := ux.GetNotifyToolbarHandler(); err != nil {
		os.Exit(1)
	} else {
		var numRead uintptr
		winHWND := ux.FinWindowByClassAndTitle(CONTEXT_MENU_CLASS, CONTEXT_MENU_TITLE)
		var tbProcessID uint32
		toolbarThreadId := win.GetWindowThreadProcessId(toolbarHandler, &tbProcessID)
		fmt.Printf("ProcessId is %d ThreadId is %d \n", tbProcessID, toolbarThreadId)
		processHandler := windows.OpenProcessAllAccess(false, tbProcessID)
		fmt.Printf("ProcessHandler is %d \n", processHandler)
		infoBaseAddress := ux.GetTBButtonInfoAllocation(processHandler)
		fmt.Printf("Base adrress is %d \n", infoBaseAddress)
		if buttonsCount := buttonsCountOnTray(); buttonsCount > 0 {
			fmt.Printf("Number of buttons %d \n", buttonsCount)
			for i := 0; i < buttonsCount; i++ {
				win.SendMessage(winHWND, win.TB_GETBUTTONTEXT, uintptr(i), infoBaseAddress)
				var destination [200]byte
				if dataRead := windows.ReadProcessMemory(processHandler, infoBaseAddress,
					uintptr(unsafe.Pointer(&destination[0])),
					100,
					&numRead); dataRead == 0 {
					fmt.Print("Nothing read \n")
				} else {
					fmt.Printf("On windows Button %d is %s\n", i, string(destination[:]))

				}
			}
		}
	}
}

// Deprecated ??
func ClickOn() {
	winHWND := ux.FinWindowByClassAndTitle(CONTEXT_MENU_CLASS, CONTEXT_MENU_TITLE)
	if winHWND > 0 {

		var rect win.RECT
		if win.GetWindowRect(winHWND, &rect) {
			fmt.Printf("Get rect top: %d, left: %d right: %d, bottom: %d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
		}
		// topMargin := uint16(2)
		// menuItemHeight := uint16(22)
		// menuItemWitdh := int32(215)
		// separatorHeight := uint16(6)
		y := uint16(65)
		x := uint16(rect.Right / 2)
		fmt.Printf("Will click at x: %d, y: %d\n", x, y)

		// Test
		// win.SendMessage(winHWND, win.WM_MOUSEMOVE, win.MK_RBUTTON, windows.MakeLPARAM(x, ya))
		// Test2
		//ux.MouseClick(int32(x), int32(y))

	}
	// After click pick the menu items
	// https://stackoverflow.com/questions/16271512/how-to-get-handle-for-menu-items-of-an-application-running-in-the-back-ground-fr
}
