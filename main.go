package main

import (
	"fmt"
	"os"

	"github.com/adrianriobo/gowinx/pkg/ux/notify"
	"github.com/lxn/win"
)

const (
	NIOW_CLASS string = "NotifyIconOverflowWindow"
)

func main() {
	//Show notification area (hidden)
	notify.ShowNotifyIconOverflowWindow()
	if tbHWND, err := notify.GetNotifyToolbarHandler(); err != nil {
		os.Exit(1)
	} else {
		var rect win.RECT
		if win.GetWindowRect(tbHWND, &rect) {
			fmt.Printf("Get rect top: %d, left: %d \n", rect.Top, rect.Left)
		}
		if buttonsCount, err := notify.GetButtonsCountONotifyToolbar(); err != nil {
			os.Exit(1)
		} else {
			fmt.Printf("There are %d buttons on the notify toolbar \n", buttonsCount)
		}
	}
}
