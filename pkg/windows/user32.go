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
	clipCursor   = moduleUser32.NewProc("ClipCursor")
	getMenu      = moduleUser32.NewProc("GetMenu")
)

// https://github.com/allendang/w32/blob/ad0a36d80adc/kernel32.go#L321
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

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-clipcursor?redirectedfrom=MSDN
func ClipCursor(lpRect *win.RECT) uintptr {
	ret, _, _ := syscall.Syscall(clipCursor.Addr(), 1,
		uintptr(unsafe.Pointer(lpRect)),
		0,
		0)
	// Change this to bool value
	return ret
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenu
// HMENU GetMenu(
// 	HWND hWnd
//   );
func GetMenu(hWnd win.HWND) win.HMENU {
	ret, _, _ := syscall.Syscall(getMenu.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return win.HMENU(ret)
}
