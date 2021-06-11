// +build windows

package main

import (
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/notificationarea"
)

func main() {

	// WORKING

	// if notifyAreaRect, err := notificationarea.GetHiddenNotificationAreaRect(); err == nil {
	// 	x, y := notificationarea.GetIconPosition(notifyAreaRect)
	// 	ux.Click(x, y)

	// 	stopX, stopY := crc.MenuItemPosition(crc.CONTEXT_MENU_ITEM_STOP)
	// 	ux.Click(stopX, stopY)

	// 	if hiddenIcons, err := notificationarea.GetHiddenIconsCount(); err != nil {
	// 		fmt.Printf("error getting hidden icons count %v\n", err)
	// 	} else {
	// 		fmt.Printf("number of hidden icons is %d\n", hiddenIcons)
	// 	}

	// 	// if visibleIcons, err := notificationarea.GetVisibleIconsCount(); err != nil {
	// 	// 	fmt.Printf("error getting visible icons count %v\n", err)
	// 	// } else {
	// 	// 	fmt.Printf("number of visible icons is %d\n", visibleIcons)
	// 	// }
	// }

	// NOT WORKING

	// notificationarea.GetIconByTittle("Codeready Containers")
	notificationarea.ShowHiddenNotificationArea()
	// rect, _ := notificationarea.GetIconRectByTittle("Codeready Containers")
	// ux.ClickOnRect(rect)
	// Get notification icon check tray notifications to system
	// notificationarea.FindChildElement("TrayButton"

	notificationarea.GetButtonsTexts()

}
