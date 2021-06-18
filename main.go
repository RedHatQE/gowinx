// +build windows

package main

import (
	"fmt"

	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/windows-accesibility-features"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/menu"
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

	win32waf.Initalize()
	crcMenu, err := menu.GetMenuFromRoot("crc")
	if err != nil {
		fmt.Printf("Error with %v", err)
	}
	menuitem, err := menu.GetMenuItem(crcMenu, "crc-delete")
	if err != nil {
		fmt.Printf("Error with %v", err)
	}
	menuitemPosition, err := menu.GetMenuItemRect(menuitem)
	if err != nil {
		fmt.Printf("Error with %v", err)
	}
	interaction.ClickOnRect(*menuitemPosition)
	win32waf.Finalize()
}
