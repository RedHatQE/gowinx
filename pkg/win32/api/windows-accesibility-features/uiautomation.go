// +build windows
package windows_accesibility_features

import (
	"syscall"
)

var (
	elementFromHandle = uiautomationClient.MustFindProc("ElementFromHandle")
)

// https://docs.microsoft.com/en-us/windows/win32/api/uiautomationclient/nf-uiautomationclient-iuiautomation-elementfromhandle
// HRESULT ElementFromHandle(
// 	UIA_HWND             hwnd,
// 	IUIAutomationElement **element
// );
func ElementFromHandle(hwnd syscall.Handle, element uintptr) (successedEventsNumber int32, err error) {
	// r0, _, e1 := syscall.Syscall(elementFromHandle.Addr(), 2,
	// 	uintptr(hwnd),
	// 	uintptr(element),
	// 	0)
	// success, err = evalSyscallBool(r0, e1)
	return
}
