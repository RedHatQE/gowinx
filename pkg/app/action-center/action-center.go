// +build windows

package action_center

import (
	"fmt"
	"syscall"

	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-accesibility-features"
	win32wam "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-and-messages"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/notificationarea"
	"github.com/adrianriobo/gowinx/pkg/win32/interaction"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/button"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

const (
	icon_name        string = "Action Center"
	window_title     string = "Action Center"
	clear_all_button string = "Clear all notifications"
)

func ClickNotifyButton() error {
	handler, err := notificationarea.FindTrayButtonByTitle(icon_name)
	if err != nil {
		return err
	}
	rect, err := getActionCenterIconPosition(handler)
	if err != nil {
		return err
	}
	return interaction.ClickOnRect(rect)
}

func ClearNotifications() error {
	// Initialize base elements
	intialize()

	actionCenterWindow, err := windows.GetActiveWindow(window_title)
	if err != nil {
		return err
	}
	clearAllButton, err := button.GetButton(actionCenterWindow, clear_all_button)
	if err != nil {
		return err
	}

	if err := button.Click(clearAllButton); err != nil {
		return err
	}
	finalize()
	return nil
}

func getActionCenterIconPosition(handler syscall.Handle) (win32wam.RECT, error) {
	var rect win32wam.RECT
	if succeed, err := win32wam.GetWindowRect(handler, &rect); succeed {
		fmt.Printf("Rect for action center icon is t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
		return rect, nil
	} else {
		return win32wam.RECT{}, err
	}
}

func intialize() error {
	// Initialize context
	win32waf.Initalize()
	// Click notifiy button to expand action center
	return ClickNotifyButton()
}

func finalize() {
	// Finalize context
	win32waf.Finalize()
}
