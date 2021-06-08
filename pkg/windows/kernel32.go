// +build windows
package windows

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

var (
	moduleKernel32 = windows.NewLazySystemDLL("kernel32.dll")

	virtualAllocEx    = moduleKernel32.NewProc("VirtualAllocEx")
	openProcess       = moduleKernel32.NewProc("OpenProcess")
	readProcessMemory = moduleKernel32.NewProc("ReadProcessMemory")
)

// https://github.com/elastic/go-windows/blob/f97ca94f20d7a0b96d53964653b37da30f9dfbfc/zsyscall_windows.go#L119
func ReadProcessMemory(handle win.HWND, baseAddress uintptr, buffer uintptr, size uintptr, numRead *uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(readProcessMemory.Addr(), 5,
		uintptr(handle),
		uintptr(baseAddress),
		uintptr(buffer),
		uintptr(size),
		uintptr(unsafe.Pointer(numRead)),
		0)

	return ret
}

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
