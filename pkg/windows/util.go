// +build windows
package windows

func MakeLPARAM(hiword, loword uint16) uintptr {
	return uintptr((hiword << 16) | uint16(loword&0xffff))
}
