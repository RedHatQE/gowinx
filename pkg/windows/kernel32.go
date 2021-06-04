// +build windows

package windows

import (
	"C"

	"golang.org/x/sys/windows"
)
import (
	"syscall"

	"github.com/lxn/win"
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

var (
	moduleKernel32 = windows.NewLazySystemDLL("kernel32.dll")

	virtualAllocEx = moduleKernel32.NewProc("VirtualAllocEx")
	openProcess    = moduleKernel32.NewProc("OpenProcess")
)

func VirtualAllocEx(hProcess C.HANDLE, lpAddress C.LPVOID, dwSize uintptr, flAllocationType, flProtect uint32) C.LPVOID {
	ret, _, _ := syscall.Syscall6(virtualAllocEx.Addr(), 5,
		uintptr(hProcess),
		uintptr(lpAddress),
		dwSize,
		uintptr(flAllocationType),
		uintptr(flProtect),
		0)

	return C.LPVOID(ret)
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
