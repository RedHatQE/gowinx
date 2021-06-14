// +build windows
package toolbar

import (
	"fmt"
	"math"
	"syscall"
	"unsafe"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
	win32process "github.com/adrianriobo/gowinx/pkg/win32/services/process"
	win32windows "github.com/adrianriobo/gowinx/pkg/win32/ux/windows"
)

// Toolbars are a way to group commands for efficient access.

const (
	// Review this 256
	BUTTONTEXT_MAX_SIZE   = 256
	BUTTON_DEFAULT_HEIGHT = 40
	BUTTON_DEFAULT_WIDHT  = 40

	TOOLBAR_WINDOW32_CLASS string = "ToolbarWindow32"
)

// Get an array of toolbars children from a parent window
func FindToolbars(parentHandler syscall.Handle) ([]syscall.Handle, error) {
	return win32windows.FindChildWindowsbyClass(parentHandler, TOOLBAR_WINDOW32_CLASS)
}

// Get coordenates to click on a button identified by its text or error if button
// does not exist on toolbar
// func GetButtonClickPosition(toolbarHandler syscall.Handle, text string) (x, y int32, err error) {

// }

// Find the index of a button identified by its tests on a toolbar return error in case
// the button is not on the toolbar
func GetButtonClickablePosition(toolbarHandler syscall.Handle, text string) (int, int, error) {
	buttonsCount := getButtonsCount(toolbarHandler)
	for n := 0; n < buttonsCount; n++ {
		buttonText, _ := getButtonText(toolbarHandler, n)
		if buttonText == text {
			return calculateButtonPosition(toolbarHandler, n)
		}
	}
	return -1, -1, fmt.Errorf("toolbar does not contain the button")
}

func calculateButtonPosition(toolbarHandler syscall.Handle, commandId int) (int, int, error) {
	index := getButtonIndex(toolbarHandler, commandId)
	toolbarRect, err := getToolbarRect(toolbarHandler)
	if err != nil {
		elementsPerRow := int(toolbarRect.Left-toolbarRect.Right) / BUTTON_DEFAULT_WIDHT
		row := index / elementsPerRow
		indexOnRow := int(math.Mod(float64(index), float64(elementsPerRow)))
		x := (indexOnRow * BUTTON_DEFAULT_WIDHT) + int(toolbarRect.Left) + (BUTTON_DEFAULT_WIDHT / 2)
		y := (row * BUTTON_DEFAULT_HEIGHT) + int(toolbarRect.Top) + (BUTTON_DEFAULT_HEIGHT / 2)
		return x, y, nil
	}
	return -1, -1, fmt.Errorf("can not get button position")

}
func getElementsPerRow(rect win32api.RECT) int {
	return int(rect.Left-rect.Right) / BUTTON_DEFAULT_WIDHT
}

func getRowForIndex(index, elementsPerRow int) int {
	return index / elementsPerRow
}

func getRelativeIndex(index, elementsPerRow int) int {
	return int(math.Mod(float64(index), float64(elementsPerRow)))
}

// Get the toolbar rectangle
func getToolbarRect(toolbarHandler syscall.Handle) (win32api.RECT, error) {
	var rect win32api.RECT
	if succeed, err := win32api.GetWindowRect(toolbarHandler, &rect); succeed {
		fmt.Printf("Rect for toolbar is t:%d,l:%d,r:%d,b:%d\n", rect.Top, rect.Left, rect.Right, rect.Bottom)
		return rect, nil
	} else {
		return win32api.RECT{}, err
	}
}

func getButtonsCount(toolbarHandler syscall.Handle) int {
	buttonsCount, _ := win32api.SendMessage(toolbarHandler, win32api.TB_BUTTONCOUNT, 0, 0)
	return int(buttonsCount)
}

// Buttons are indentified by commandId on a toolbar, the posistion can be re arranged
// and it is defined by the index. This function retrieves the button index based on its
// command id
func getButtonIndex(toolbarHandler syscall.Handle, commandId int) int {
	index, _ := win32api.SendMessage(
		toolbarHandler,
		win32api.TB_COMMANDTOINDEX,
		uintptr(commandId),
		0)
	return int(index)
}

// To get button text, the communication requires read / write on a memory address, to request
// the space to OS we need to create a handler process from the thread which created the toolbar
func getButtonText(toolbarHandler syscall.Handle, commandId int) (string, error) {
	processHandler, err := win32process.GetProcessHandler(toolbarHandler)
	if processHandler > 0 {
		infoBaseAddress, err := win32process.AllocateMemory(processHandler, BUTTONTEXT_MAX_SIZE)
		if infoBaseAddress > 0 {
			length := requestButtonText(toolbarHandler, commandId, infoBaseAddress)
			text, err := readButtonText(processHandler, infoBaseAddress, 2*length)
			// should define how to treat memory and process free errors.
			win32process.FreeMemory(processHandler, infoBaseAddress)
			win32process.CloseProcessHandler(processHandler)
			return text, err
		}
		return "", err
	}
	return "", err
}

// Tell toolbar handler to put the text of the button identified by the command id, on a memory address
func requestButtonText(toolbarHandler syscall.Handle, commandId int, memoryBaseAddress uintptr) (textlength int) {
	length, _ := win32api.SendMessage(
		toolbarHandler,
		win32api.TB_GETBUTTONTEXT,
		uintptr(commandId),
		memoryBaseAddress)
	textlength = int(length)
	return
}

// Read text of the button on a memory address
func readButtonText(processHandler syscall.Handle, memoryBaseAddress uintptr, length int) (text string, err error) {
	n := make([]byte, BUTTONTEXT_MAX_SIZE)
	var numRead uintptr
	p := &n[0]
	dataRead, err := win32api.ReadProcessMemory(processHandler, memoryBaseAddress,
		uintptr(unsafe.Pointer(p)),
		uintptr(length),
		&numRead)
	if dataRead {
		text = string(n[:numRead])
		fmt.Printf("Button with is %s\n", string(n[:numRead]))
	}
	return
}
