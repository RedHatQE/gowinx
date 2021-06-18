package ux

import (
	"fmt"

	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/windows-accesibility-features"
	"github.com/go-ole/go-ole"
	wa "github.com/hnakamur/w32uiautomation"
)

func GetToolbarByName(name string) (*wa.IUIAutomationElement, error) {
	conditionByName, err := win32waf.CreatePropertyCondition(
		wa.UIA_NamePropertyId,
		wa.NewVariantString(name))
	if err != nil {
		return &wa.IUIAutomationElement{}, err
	}
	conditionByType, err := win32waf.CreatePropertyCondition(wa.UIA_ControlTypePropertyId,
		ole.NewVariant(ole.VT_INT, wa.UIA_MenuItemControlTypeId))
	if err != nil {
		return &wa.IUIAutomationElement{}, err
	}
	condition, err := win32waf.CreateAndCondition(conditionByName, conditionByType)
	if err != nil {
		return &wa.IUIAutomationElement{}, err
	}
	if toolbar, err := win32waf.FindFirstFromRoot(wa.TreeScope_Children, condition); err != nil {
		fmt.Printf("error getting toolbar %s: %v", name, err)
		return &wa.IUIAutomationElement{}, err
	} else {
		return toolbar, nil
	}
}
