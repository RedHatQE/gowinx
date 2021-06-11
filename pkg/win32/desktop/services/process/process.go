// +build windows
package process

import (
	"syscall"

	win32api "github.com/adrianriobo/gowinx/pkg/win32/api"
)

const MEM_COMMIT = 0x1000
const PAGE_READWRITE = 0x04

// Get a process handler for the process holding the window (ux element, represented by its handler)
// This is required in order to run communications with the window.
func GetProcessHandler(windowHandler syscall.Handle) (processHandler syscall.Handle, err error) {
	var tbProcessID uint32
	toolbarThreadId, err := win32api.GetWindowThreadProcessId(windowHandler, &tbProcessID)
	if toolbarThreadId > 0 {
		processHandler, err = win32api.OpenProcessAllAccess(false, tbProcessID)
	}
	return
}

func AllocateMemory(processHandler syscall.Handle, size int) (uintptr, error) {
	return win32api.VirtualAllocEx(processHandler, 0, uintptr(size), MEM_COMMIT, PAGE_READWRITE)
}
