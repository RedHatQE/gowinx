// +build windows

package list

import (
	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-accesibility-features"
	wa "github.com/openstandia/w32uiautomation"
)

func GetList(parentElement *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	return win32waf.GetElementFromParent(parentElement, name, wa.UIA_ListControlTypeId)
}

func GetListItem(parentElement *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	return win32waf.GetElementFromParent(parentElement, name, wa.UIA_ListItemControlTypeId)
}
