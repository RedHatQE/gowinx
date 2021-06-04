// +build windows

package windows

import (
	"syscall"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

var (
	moduleKernel32 = windows.NewLazySystemDLL("kernel32.dll")

	virtualAllocEx = moduleKernel32.NewProc("VirtualAllocEx")
	openProcess    = moduleKernel32.NewProc("OpenProcess")
)

func VirtualAllocEx(hProcess win.HWND, lpAddress, dwSize uintptr, flAllocationType, flProtect uint32) uintptr {
	ret, _, _ := syscall.Syscall6(virtualAllocEx.Addr(), 5,
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(flAllocationType),
		uintptr(flProtect),
		0)

	return ret
}

func OpenProcessAllAccess(inheritHandle bool, processId uint32) win.HWND {
	return execOpenProcess(PROCESS_ALL_ACCESS, inheritHandle, processId)
}

func execOpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) win.HWND {
	inherit := 0
	if inheritHandle {
		inherit = 1
	}
	ret, _, _ := syscall.Syscall(openProcess.Addr(), 3,
		uintptr(desiredAccess),
		uintptr(inherit),
		uintptr(processId))

	return win.HWND(ret)

}
