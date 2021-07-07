// +build windows

package group

import (
	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-accesibility-features"
	wa "github.com/openstandia/w32uiautomation"
)

func GetGroup(parentElement *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	return win32waf.GetElementFromParent(parentElement, name, wa.UIA_GroupControlTypeId)
}
