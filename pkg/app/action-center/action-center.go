package action_center

import (
	"fmt"
	"syscall"

	win32wam "github.com/adrianriobo/gowinx/pkg/win32/api/windows-and-messages"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/notificationarea"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/interaction"
)

const (
	ACTION_CENTER_NAME string = "Action Center"
)

func Click() error {
	handler, err := notificationarea.FindTrayButtonByTitle(ACTION_CENTER_NAME)
	if err != nil {
		return err
	}
	rect, err := getActionCenterIconPosition(handler)
	if err != nil {
		return err
	}
	return interaction.ClickOnRect(rect)
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
