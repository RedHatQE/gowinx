// +build windows
package ux

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/adrianriobo/gowinx/pkg/win32"
)

const (
	elementClickDelay time.Duration = 300 * time.Millisecond
)

type MOUSE_INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

func Click(x, y int32) error {
	events := [3]uint32{
		win32.MOUSEEVENTF_ABSOLUTE | win32.MOUSEEVENTF_MOVE,
		win32.MOUSEEVENTF_LEFTDOWN,
		win32.MOUSEEVENTF_LEFTUP}
	for _, event := range events {
		time.Sleep(elementClickDelay)
		if err := mouseInput(x, y, uint32(event)); err != nil {
			return err
		}
	}
	return nil
}

func mouseInput(x, y int32, dwFlags uint32) error {
	dx := getDX(x)
	dy := getDY(y)
	fmt.Printf("Click done at x:%d y:%d \n", dx, dy)
	mouseInput := MOUSE_INPUT{
		Type: win32.INPUT_MOUSE,
		Mi: MOUSEINPUT{
			Dx:          dx,
			Dy:          dy,
			MouseData:   uint32(0),
			DwFlags:     dwFlags,
			Time:        uint32(0),
			DwExtraInfo: uintptr(0)}}

	actions := [1]MOUSE_INPUT{mouseInput}

	if success, err := win32.SendInput(uint32(2), unsafe.Pointer(&actions), int32(unsafe.Sizeof(mouseInput))); err == nil {
		fmt.Printf("Input sent successfull returns %d actions\n", success)
		return nil
	} else {
		return err
	}
}

func getDX(x int32) int32 {
	return getDAxisValue(x, win32.SM_CXSCREEN)
}

func getDY(y int32) int32 {
	return getDAxisValue(y, win32.SM_CYSCREEN)
}

func getDAxisValue(axisValue, systemMetricConstant int32) int32 {
	if metric, err := win32.GetSystemMetrics(systemMetricConstant); err == nil {
		return (axisValue * 65536) / metric
	} else {
		fmt.Print(err)
		return 0
	}
}
