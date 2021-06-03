package windows

import (
	"C"

	"golang.org/x/sys/windows"
)
import (
	"syscall"
)

var (
	module = windows.NewLazySystemDLL("kernel32.dll")

	virtualAllocEx = module.NewProc("VirtualAllocEx")
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
