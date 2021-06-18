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

// func GetMenuByName(menuName string) win32wam.RECT {
// 	condVal := wa.NewVariantString(menuName)
// 	condition, err := win32waf.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
// 	crcMenu, _ := win32waf.FindFirstFromRoot(wa.TreeScope_Children, condition)
// 	name, err := crcMenu.Get_CurrentName()
// 	if err != nil {
// 		fmt.Printf("menu error %v", err)
// 	} else {
// 		fmt.Printf("menu name %v", name)
// 	}
// 	condValDel := wa.NewVariantString("crc-delete")
// 	conditionDel, err := win32waf.CreatePropertyCondition(wa.UIA_NamePropertyId, condValDel)
// 	if err != nil {
// 		fmt.Printf("error in condition %v", err)
// 	}
// 	buttonDelete, err := win32waf.FindFirst(crcMenu, wa.TreeScope_Children, conditionDel)
// 	name, err = buttonDelete.Get_CurrentName()
// 	if err != nil {
// 		fmt.Printf("button delete error %v", err)
// 	} else {
// 		fmt.Printf("button delete name %v", name)
// 	}

// 	// buttonDelete.GetCurrentPattern(wa.UIA_SelectionPatternId)
// 	// wa.Invoke()
// 	// wa.UIA_ValuePatternId wa.UIA_TextPatternId

// 	// return win32wam.RECT{}

// 	// condValMenu := wa.
// 	// UIA_MenuItemControlTypeId
// 	// UIA_ControlTypePropertyId
// 	// condMenuitem := ole.NewVariant(ole.VT_INT, wa.UIA_MenuItemControlTypeId)
// 	// conditionDel, err := win32waf.CreatePropertyCondition(wa.UIA_ControlTypePropertyId, condMenuitem)
// 	// if err != nil {
// 	// 	fmt.Printf("error in condition %v", err)
// 	// }
// 	// _, err = win32waf.FindFirst(crcMenu, wa.TreeScope_Children, conditionDel)
// 	// // name, err := buttonDelete.Get_CurrentName()
// 	// if err != nil {
// 	// 	fmt.Printf("crc error %v", err)
// 	// }

// 	// if _, err := win32waf.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal); err != nil {
// 	// 	// return nil, err
// 	// } else {
// 	// 	fmt.Print("Condition was created")
// 	// if menu, err := win32waf.FindFirstFromRoot(wa.TreeScope_Children, condition); err != nil {
// 	// if menu, err := ; err != nil {

// 	// return nil, err
// 	// } else {
// 	// return menu, nil
// 	// }
// 	// menu, err := win32waf.FindFirstFromRoot(wa.TreeScope_Children, condition)
// 	// ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
// 	// defer ole.CoUninitialize()

// 	// auto, err := wa.NewUIAutomation()
// 	// if err != nil {
// 	// 	fmt.Printf("auto error %v", err)
// 	// }
// 	// condition2, err := win32waf.CreatePropertyCondition(wa.UIA_NamePropertyId, condVal)
// 	// root, err := auto.GetRootElement()
// 	// if err != nil {
// 	// 	fmt.Printf("auto error %v", err)
// 	// }
// 	// defer root.Release()

// 	// menu, err := wa.WaitFindFirst(auto, root, wa.TreeScope_Children, condition2)
// 	// if err != nil {
// 	// } else {
// 	// 	name, err := menu.Get_CurrentName()
// 	// 	if err != nil {
// 	// 		fmt.Printf("crc error %v", err)
// 	// 	} else {
// 	// 		fmt.Printf("crc name %v", name)
// 	// 	}
// 	// }
// 	// }
// }
