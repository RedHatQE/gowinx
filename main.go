// +build windows

package main

import (
	"github.com/adrianriobo/gowinx/pkg/crc"
	"github.com/adrianriobo/gowinx/pkg/ux"
)

func main() {
	if notifyAreaRect, err := ux.GetNotifyAreaRect(); err == nil {
		x, y := ux.GetIconPosition(notifyAreaRect)
		ux.Click(x, y)

		stopX, stopY := crc.MenuItemPosition(crc.CONTEXT_MENU_ITEM_STOP)
		ux.Click(stopX, stopY)
	}
}
