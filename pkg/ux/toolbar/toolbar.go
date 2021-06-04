package toolbar

// +build windows

import (
	"unsafe"

	"github.com/adrianriobo/gowinx/pkg/windows"
	"github.com/lxn/win"
)

const MEM_COMMIT = 0x1000
const PAGE_READWRITE = 0x04

func GetTBButtonInfoAllocation(hProcess win.HWND) uintptr {
	return windows.VirtualAllocEx(hProcess, 0, GetTBButtonInfoSize(), MEM_COMMIT, PAGE_READWRITE)
}

func GetTBButtonInfoSize() uintptr {
	var tbButtonInfo win.TBBUTTONINFO
	return unsafe.Sizeof(tbButtonInfo)
}
