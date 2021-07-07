// +build windows

package menu

import (
	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-accesibility-features"
	win32wam "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-and-messages"
	wa "github.com/openstandia/w32uiautomation"
)

func GetActiveMenu(name string) (*wa.IUIAutomationElement, error) {
	return win32waf.GetActiveElement(name, wa.UIA_MenuControlTypeId)
}

func GetMenuItem(menu *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	return win32waf.GetElementFromParent(menu, name, wa.UIA_MenuItemControlTypeId)
}

func GetMenuItemRect(menuItem *wa.IUIAutomationElement) (*win32wam.RECT, error) {
	return win32waf.GetElementRect(menuItem)
}
