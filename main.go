// +build windowsls

package main

import (
	"fmt"
	"os"

	"github.com/adrianriobo/gowinx/pkg/ux/notify"
	"github.com/adrianriobo/gowinx/pkg/windows"
	"github.com/lxn/win"
)

const (
	NIOW_CLASS string = "NotifyIconOverflowWindow"
)

func main() {
	//Show notification area (hidden)
	notify.ShowNotifyIconOverflowWindow()
	if toolbarHandler, err := notify.GetNotifyToolbarHandler(); err != nil {
		os.Exit(1)
	} else {
		var rect win.RECT
		if win.GetWindowRect(toolbarHandler, &rect) {
			fmt.Printf("Get rect top: %d, left: %d \n", rect.Top, rect.Left)
		}
		var tbProcessID uint32
		toolbarThreadId := win.GetWindowThreadProcessId(toolbarHandler, &tbProcessID)
		fmt.Printf("ProcessId is %d ThreadId is %d \n", tbProcessID, toolbarThreadId)
		processHandler := windows.OpenProcessAllAccess(false, toolbarThreadId)
		fmt.Printf("ProcessHandler is %d \n", processHandler)

		if buttonsCount, err := notify.GetButtonsCountONotifyToolbar(); err != nil {
			os.Exit(1)
		} else {
			fmt.Printf("There are %d buttons on the notify toolbar \n", buttonsCount)
		}
	}
}
