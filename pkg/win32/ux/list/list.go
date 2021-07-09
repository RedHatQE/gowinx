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

func GetAllListItems(parentElement *wa.IUIAutomationElement) ([]*wa.IUIAutomationElement, error) {
	var listItems []*wa.IUIAutomationElement
	children, err := win32waf.GetAllChildren(parentElement, wa.UIA_ListItemControlTypeId)
	if err != nil {
		return nil, err
	}
	childrenCount, err := children.Get_Length()
	if err != nil {
		return nil, err
	}
	var i int32
	for i = 0; i < childrenCount; i++ {
		if listItem, err := children.GetElement(i); err == nil {
			listItems = append(listItems, listItem)
		}
	}
	return listItems, nil
}
