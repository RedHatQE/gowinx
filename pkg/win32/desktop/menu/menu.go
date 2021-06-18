// +build windows
package menu

import (
	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/windows-accesibility-features"
	win32wam "github.com/adrianriobo/gowinx/pkg/win32/api/windows-and-messages"
	"github.com/go-ole/go-ole"
	wa "github.com/hnakamur/w32uiautomation"
)

func GetMenuFromRoot(name string) (*wa.IUIAutomationElement, error) {
	root, err := win32waf.GetRootElement()
	defer root.Release()
	if err != nil {
		return nil, err
	} else {
		return GetMenu(root, name)
	}
}

func GetMenu(parentElement *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	condition, err := win32waf.CreatePropertyCondition(
		wa.UIA_NamePropertyId,
		wa.NewVariantString(name))
	if err != nil {
		return nil, err
	}
	menu, err := win32waf.FindFirst(parentElement, wa.TreeScope_Children, condition)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func GetMenuItem(parentElement *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	conditionByName, err := win32waf.CreatePropertyCondition(
		wa.UIA_NamePropertyId,
		wa.NewVariantString(name))
	if err != nil {
		return nil, err
	}
	conditionByType, err := win32waf.CreatePropertyCondition(
		wa.UIA_ControlTypePropertyId,
		ole.NewVariant(ole.VT_INT, wa.UIA_MenuItemControlTypeId))
	if err != nil {
		return nil, err
	}
	condition, err := win32waf.CreateAndCondition(conditionByName, conditionByType)
	if err != nil {
		return nil, err
	}
	menuitem, err := win32waf.FindFirst(parentElement, wa.TreeScope_Children, condition)
	if err != nil {
		return nil, err
	}
	return menuitem, nil
}

func GetMenuItemRect(menuItemElement *wa.IUIAutomationElement) (*win32wam.RECT, error) {
	rect, err := menuItemElement.Get_CurrentBoundingRectangle()
	if err != nil {
		return nil, err
	}
	return &win32wam.RECT{Top: int32(rect.Top),
		Right:  int32(rect.Right),
		Bottom: int32(rect.Bottom),
		Left:   int32(rect.Left)}, nil
}
