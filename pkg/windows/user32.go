// +build windows

package windows

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

var (
	moduleUser32 = windows.NewLazySystemDLL("user32.dll")

	findWindowEx = moduleUser32.NewProc("FindWindowExW")
)

// https://social.msdn.microsoft.com/Forums/vstudio/en-US/82cf3f2b-b661-47c5-854d-dcd42b0d45c4/how-to-click-toolbar-button-in-another-application-using-api?forum=csharpgeneral

func FindWindowEx(hWndParent, hWndChildAfter win.HWND, lpszClass, lpszWindow *uint16) win.HWND {
	ret, _, _ := syscall.Syscall6(findWindowEx.Addr(), 4,
		uintptr(hWndParent),
		uintptr(hWndChildAfter),
		uintptr(unsafe.Pointer(lpszClass)),
		uintptr(unsafe.Pointer(lpszWindow)),
		0,
		0)

	return win.HWND(ret)
}
