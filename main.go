// +build windows

package main

import (
	"github.com/adrianriobo/gowinx/pkg/crc"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/notificationarea"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/interaction"
)

func main() {

	// Sample testing code

	// Click on icon from notification area
	notificationarea.ShowHiddenNotificationArea()
	if x, y, err := notificationarea.GetIconPositionByTitle("Codeready Containers"); err == nil {
		interaction.Click(int32(x), int32(y))
	}

	// Click on menu
	stopX, stopY := crc.MenuItemPosition(crc.CONTEXT_MENU_ITEM_STOP)
	interaction.Click(stopX, stopY)

	// Try with children
	// if err := crc.FindAllChildren(); err != nil {
	// 	fmt.Printf("error on uiautomation %v\n", err)
	// }

}
