// +build windows
package win32

import (
	"syscall"
	"unsafe"
)

var (
	virtualAllocEx    = kernel32.MustFindProc("VirtualAllocEx")
	readProcessMemory = kernel32.MustFindProc("ReadProcessMemory")
)

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-virtualallocex
// LPVOID VirtualAllocEx(
// 	HANDLE hProcess,
// 	LPVOID lpAddress,
// 	SIZE_T dwSize,
// 	DWORD  flAllocationType,
// 	DWORD  flProtect
// );
func VirtualAllocEx(hProcess syscall.Handle, lpAddress, dwSize uintptr, flAllocationType, flProtect uint32) (baseAddress uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(virtualAllocEx.Addr(), 5, uintptr(hProcess), lpAddress, dwSize,
		uintptr(flAllocationType), uintptr(flProtect), 0)
	baseAddress = r0
	if e1 != 0 {
		err = error(e1)
	} else {
		err = syscall.EINVAL
	}
	return
}

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory
// BOOL ReadProcessMemory(
// 	HANDLE  hProcess,
// 	LPCVOID lpBaseAddress,
// 	LPVOID  lpBuffer,
// 	SIZE_T  nSize,
// 	SIZE_T  *lpNumberOfBytesRead
// );
func ReadProcessMemory(hProcess syscall.Handle, lpBaseAddress, lpBuffer, nSize uintptr, numRead *uintptr) (success bool, err error) {
	r0, _, e1 := syscall.Syscall6(readProcessMemory.Addr(), 5,
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(lpBuffer),
		uintptr(nSize),
		uintptr(unsafe.Pointer(numRead)),
		0)
	success, err = evalSyscallBool(r0, e1)
	return
}
