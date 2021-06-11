// +build windows
package toolbar

import (
	"fmt"
	"syscall"
	"unsafe"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
	win32process "github.com/adrianriobo/gowinx/pkg/win32/desktop/services/process"
)

const (
	// Review this 256
	BUTTONTEXT_MAX_SIZE = 256
)

func GetCommandIndex(toolbarHandler syscall.Handle, commandId int) int {
	index, _ := win32api.SendMessage(
		toolbarHandler,
		win32api.TB_COMMANDTOINDEX,
		uintptr(commandId),
		0)
	return int(index)
}

func requestButtonText(toolbarHandler syscall.Handle, commandId int, memoryBaseAddress uintptr) int {
	length, _ := win32api.SendMessage(
		toolbarHandler,
		win32api.TB_GETBUTTONTEXT,
		uintptr(commandId),
		memoryBaseAddress)
	return int(length)
}

func readButtonText(processHandler syscall.Handle, memoryBaseAddress uintptr, length int) {
	var numRead uintptr
	if dataRead, err := win32api.ReadProcessMemory(processHandler, memoryBaseAddress,
		uintptr(unsafe.Pointer(p)),
		uintptr(length),
		&numRead); !dataRead {
		fmt.Print("Nothing read \n")
		return "", nil
	} else {
		text := string(n[:numRead])
		fmt.Printf("Button with is %s\n", string(n[:numRead]))
		return text, nil
	}
}

func GetCommandButtonText(toolbarHandler syscall.Handle, commandId int) (string, error) {
	n := make([]byte, BUTTONTEXT_MAX_SIZE)
	processHandler, err := win32process.GetProcessHandler(toolbarHandler)
	if processHandler > 0 {
		infoBaseAddress, err := win32process.AllocateMemory(processHandler, BUTTONTEXT_MAX_SIZE)
		if infoBaseAddress > 0 {
			length := requestButtonText(toolbarHandler, commandId, infoBaseAddress)
			if length > 0 {
				var numRead uintptr
				if dataRead, err := win32api.ReadProcessMemory(processHandler, infoBaseAddress,
					uintptr(unsafe.Pointer(p)),
					length*2,
					&numRead); !dataRead {
					fmt.Print("Nothing read \n")
					return "", nil
				} else {
					text := string(n[:numRead])
					fmt.Printf("Button with is %s\n", string(n[:numRead]))
					return text, nil
				}
			}
			return "", err
		}
		return "", err
	}
	return "", err
}
