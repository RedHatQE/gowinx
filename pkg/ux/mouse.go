// +build windows
package ux

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/lxn/win"
)

const (
	elementClickDelay time.Duration = 300 * time.Millisecond
)

func getDX(x int32) int32 {
	return (x * 65536) / win.GetSystemMetrics(win.SM_CXSCREEN)
}

func getDY(y int32) int32 {
	return (y * 65536) / win.GetSystemMetrics(win.SM_CYSCREEN)
}

func MouseInput(x, y int32, dwFlags uint32) error {
	dx := getDX(x)
	dy := getDY(y)
	fmt.Printf("Click done at x:%d y:%d \n", dx, dy)
	mouseInput := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			Dx:          dx,
			Dy:          dy,
			MouseData:   uint32(0),
			DwFlags:     dwFlags,
			Time:        uint32(0),
			DwExtraInfo: uintptr(0)}}

	actions := [1]win.MOUSE_INPUT{mouseInput}

	if success := win.SendInput(uint32(2), unsafe.Pointer(&actions), int32(unsafe.Sizeof(mouseInput))); success > 0 {
		fmt.Printf("Input sent successfull returns %d actions\n", success)
		return nil
	} else {
		fmt.Printf("Failed clicking\n")
		return fmt.Errorf("failed to mouse input at x:%d y:%d \n", dx, dy)
	}
}

// Evaluate why this is not working
func Click(x, y int32) error {
	events := [3]uint32{
		win.MOUSEEVENTF_ABSOLUTE | win.MOUSEEVENTF_MOVE,
		win.MOUSEEVENTF_LEFTDOWN,
		win.MOUSEEVENTF_LEFTUP}
	for _, event := range events {
		time.Sleep(elementClickDelay)
		if err := MouseInput(x, y, uint32(event)); err != nil {
			return err
		}
	}
	return nil
}

// func Click(x, y int32) error {
// 	if err := MouseInput(x, y, win.MOUSEEVENTF_ABSOLUTE|win.MOUSEEVENTF_MOVE); err != nil {
// 		return err
// 	}
// 	if err := MouseInput(x, y, win.MOUSEEVENTF_LEFTDOWN); err != nil {
// 		return err
// 	}
// 	if err := MouseInput(x, y, win.MOUSEEVENTF_LEFTUP); err != nil {
// 		return err
// 	}
// 	return nil
// }
