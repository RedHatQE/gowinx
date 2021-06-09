// +build windows
package win32

import (
	"syscall"
	"unsafe"
)

var (
	enumWindows         = user32.MustFindProc("EnumWindows")
	getWindowTextW      = user32.MustFindProc("GetWindowTextW")
	findWindowW         = user32.MustFindProc("FindWindowW")
	findWindowEx        = user32.MustFindProc("FindWindowExW")
	getForegroundWindow = user32.MustFindProc("GetForegroundWindow")
	getClassNameW       = user32.MustFindProc("GetClassNameW")
	sendMessageW        = user32.MustFindProc("SendMessageW")
	getSystemMetrics    = user32.MustFindProc("GetSystemMetrics")
	sendInput           = user32.MustFindProc("SendInput")
	getWindowRect       = user32.MustFindProc("GetWindowRect")
	getDlgItem          = user32.MustFindProc("GetDlgItem")
	showWindow          = user32.MustFindProc("ShowWindow")
)

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
// BOOL EnumWindows(
// 	WNDENUMPROC lpEnumFunc,
// 	LPARAM      lParam
// );
func EnumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r0, _, e1 := syscall.Syscall(enumWindows.Addr(), 2,
		uintptr(enumFunc),
		uintptr(lparam),
		0)
	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
// int GetWindowTextW(
// 	HWND   hWnd,
// 	LPWSTR lpString,
// 	int    nMaxCount
// );
func GetWindowTextW(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(getWindowTextW.Addr(), 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(str)),
		uintptr(maxCount))
	len, err = evalSyscallInt32(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
// HWND FindWindowW(
// 	LPCWSTR lpClassName,
// 	LPCWSTR lpWindowName
// );
func FindWindowW(lpClassName, lpWindowName *uint16) (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(findWindowW.Addr(), 2,
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0)
	hwnd, err = evalSyscallHandler(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindowexw
// HWND FindWindowExW(
// 	HWND    hWndParent,
// 	HWND    hWndChildAfter,
// 	LPCWSTR lpszClass,
// 	LPCWSTR lpszWindow
// );
func FindWindowEx(hwndParent, hwndChildAfter syscall.Handle, lpszClass, lpszWindow *uint16) (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(findWindowEx.Addr(), 4,
		uintptr(hwndParent),
		uintptr(hwndChildAfter),
		uintptr(unsafe.Pointer(lpszClass)),
		uintptr(unsafe.Pointer(lpszWindow)),
		0,
		0)
	hwnd, err = evalSyscallHandler(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getforegroundwindow
// HWND GetForegroundWindow();
func GetForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(getForegroundWindow.Addr(), 0,
		0,
		0,
		0)
	hwnd, err = evalSyscallHandler(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassnamew
// int GetClassNameW(
// 	HWND   hWnd,
// 	LPWSTR lpClassName,
// 	int    nMaxCount
// );
func GetClassName(hWnd syscall.Handle) (name string, err error) {
	n := make([]uint16, 256)
	p := &n[0]
	r0, _, e1 := syscall.Syscall(getClassNameW.Addr(), 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(p)),
		uintptr(len(n)))
	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
		return
	}
	name = syscall.UTF16ToString(n)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagew
// LRESULT SendMessageW(
// 	HWND   hWnd,
// 	UINT   Msg,
// 	WPARAM wParam,
// 	LPARAM lParam
// );
func SendMessageW(hWnd syscall.Handle, msg uint32, wParam, lParam uintptr) (lResult uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(sendMessageW.Addr(), 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)
	lResult = r0
	if e1 != 0 {
		err = error(e1)
	} else {
		err = syscall.EINVAL
	}
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
// int GetSystemMetrics(
// 	int nIndex
// );
func GetSystemMetrics(nIndex int32) (metric int32, err error) {
	r0, _, e1 := syscall.Syscall(getSystemMetrics.Addr(), 1,
		uintptr(nIndex),
		0,
		0)
	metric, err = evalSyscallInt32(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendinput
// UINT SendInput(
// 	UINT    cInputs,
// 	LPINPUT pInputs,
// 	int     cbSize
// );
func SendInput(cInputs uint32, pInputs unsafe.Pointer, cbSize int32) (successedEventsNumber int32, err error) {
	r0, _, e1 := syscall.Syscall(sendInput.Addr(), 3,
		uintptr(cInputs),
		uintptr(pInputs),
		uintptr(cbSize))
	successedEventsNumber, err = evalSyscallInt32(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowrect
// BOOL GetWindowRect(
// 	HWND   hWnd,
// 	LPRECT lpRect
// );
func GetWindowRect(hWnd syscall.Handle, rect *RECT) (success bool, err error) {
	r0, _, e1 := syscall.Syscall(getWindowRect.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)
	success, err = evalSyscallBool(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdlgitem
// HWND GetDlgItem(
// 	HWND hDlg,
// 	int  nIDDlgItem
// );
func GetDlgItem(hDlg syscall.Handle, nIDDlgItem int32) (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(getDlgItem.Addr(), 2,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		0)
	hwnd, err = evalSyscallHandler(r0, e1)
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
// BOOL ShowWindow(
// 	HWND hWnd,
// 	int  nCmdShow
// );
func ShowWindow(hWnd syscall.Handle, nCmdShow int32) (hidden bool) {
	ret, _, _ := syscall.Syscall(showWindow.Addr(), 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)
	return ret != 0
}
