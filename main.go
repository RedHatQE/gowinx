// +build windows

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"

	"github.com/adrianriobo/gowinx/pkg/ux/notify"
	"github.com/adrianriobo/gowinx/pkg/ux/toolbar"
	"github.com/adrianriobo/gowinx/pkg/windows"
	"github.com/lxn/win"
)

const (
	NIOW_CLASS string = "NotifyIconOverflowWindow"
)

type TBBUTTONINFO struct {
	CbSize    uint32
	DwMask    uint32
	IdCommand int32
	IImage    int32
	FsState   byte
	FsStyle   byte
	Cx        uint16
}

// type TBBUTTONINFO struct {
// 	CbSize    uint32
// 	DwMask    uint32
// 	IdCommand int32
// 	IImage    int32
// 	FsState   byte
// 	FsStyle   byte
// 	Cx        uint16
// 	LParam    uintptr
// 	PszText   uintptr
// 	CchText   int32
// }

func decode(b []byte) (*TBBUTTONINFO, error) {
	buf := bytes.NewBuffer(b)

	obj := &TBBUTTONINFO{}

	if err := binary.Read(buf, binary.LittleEndian, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func makeLPARAM(hiword, loword uint16) uintptr {
	return uintptr((hiword << 16) | uint16(loword&0xffff))
}

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
		processHandler := windows.OpenProcessAllAccess(false, tbProcessID)
		fmt.Printf("ProcessHandler is %d \n", processHandler)
		infoBaseAddress := toolbar.GetTBButtonInfoAllocation(processHandler)
		fmt.Printf("Base adrress is %d \n", infoBaseAddress)

		if buttonsCount, err := notify.GetButtonsCountONotifyToolbar(); err != nil {
			os.Exit(1)
		} else {
			fmt.Printf("There are %d buttons on the notify toolbar \n", buttonsCount)
			var i uintptr
			for i = 0; i < buttonsCount; i++ {
				fmt.Printf("Button %d \n", i)
				// Request button
				// win.SendMessage(toolbarHandler, win.TB_GETBUTTONINFO, i, infoBaseAddress)
				// var destination [uint32(unsafe.Sizeof(win.TBBUTTONINFO{}))]byte
				// var numRead uintptr
				// if dataRead := windows.ReadProcessMemory(processHandler, infoBaseAddress,
				// 	uintptr(unsafe.Pointer(&destination[0])),
				// 	uintptr(unsafe.Sizeof(win.TBBUTTONINFO{})),
				// 	&numRead); dataRead == 0 {
				// 	fmt.Print("Nothing read \n")
				// } else {
				// 	if tbButtoninfo, err := decode(destination[:]); err != nil {
				// 		fmt.Printf("Error decoding buttoninfo %v \n", err)
				// 	} else {
				// 		fmt.Printf("Button %d is %s\n", i, tbButtoninfo.IdCommand)
				// 	}
				// }
				// var destination [uint32(unsafe.Sizeof(win.TBBUTTONINFO{}))]byte
				var numRead uintptr

				win.SendMessage(toolbarHandler, win.TB_GETBUTTONTEXT, i, infoBaseAddress)
				var destination [200]byte
				if dataRead := windows.ReadProcessMemory(processHandler, infoBaseAddress,
					uintptr(unsafe.Pointer(&destination[0])),
					100,
					&numRead); dataRead == 0 {
					fmt.Print("Nothing read \n")
				} else {
					fmt.Printf("Button %d is %s\n", i, string(destination[:]))

				}
				if i == 1 {
					win.SendMessage(toolbarHandler, win.WM_COMMAND, i, makeLPARAM(uint16(win.GetDlgItem(toolbarHandler, int32(i))), win.BN_CLICKED))
				}

			}
			// After click pick the menu items
			// https://stackoverflow.com/questions/16271512/how-to-get-handle-for-menu-items-of-an-application-running-in-the-back-ground-fr
		}

	}
}
