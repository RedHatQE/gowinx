// +build windows
package windows_accesibility_features

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	wa "github.com/hnakamur/w32uiautomation"
)

var (
	once    sync.Once
	manager *wa.IUIAutomation
)

func Initalize() {
	once.Do(func() {
		ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
		if waManager, err := wa.NewUIAutomation(); err != nil {
			fmt.Printf("Error initializing ui automation framework: %v", err)
			os.Exit(1)
		} else {
			manager = waManager
		}
	})
}

func Finalize() {
	ole.CoUninitialize()
}

// https://docs.microsoft.com/en-us/windows/win32/api/uiautomationclient/nf-uiautomationclient-iuiautomation-createpropertycondition
// HRESULT CreatePropertyCondition(
// 	PROPERTYID             propertyId,
// 	VARIANT                value,
// 	IUIAutomationCondition **newCondition
//  );
func CreatePropertyCondition(propertyId wa.PROPERTYID, value ole.VARIANT) (*wa.IUIAutomationCondition, error) {
	var newCondition *wa.IUIAutomationCondition
	hr, _, er1 := syscall.Syscall6(
		manager.VTable().CreatePropertyCondition,
		4,
		uintptr(unsafe.Pointer(manager)),
		uintptr(propertyId),
		uintptr(unsafe.Pointer(&value)),
		uintptr(unsafe.Pointer(&newCondition)),
		0,
		0)
	if hr != 0 {
		return nil, error(er1)
	}
	return newCondition, nil
}

func CreateAndCondition(condition1, condition2 *wa.IUIAutomationCondition) (newCondition *wa.IUIAutomationCondition, err error) {
	return manager.CreateAndCondition(condition1, condition2)
}

// https://docs.microsoft.com/en-us/windows/win32/api/uiautomationclient/nf-uiautomationclient-iuiautomationelement-findfirst
// HRESULT FindFirst(
// 	TreeScope              scope,
// 	IUIAutomationCondition *condition,
// 	IUIAutomationElement   **found
// );
func FindFirst(elem *wa.IUIAutomationElement, scope wa.TreeScope, condition *wa.IUIAutomationCondition) (found *wa.IUIAutomationElement, err error) {
	return wa.WaitFindFirst(manager, elem, scope, condition)
}

func GetRootElement() (root *wa.IUIAutomationElement, err error) {
	return manager.GetRootElement()
}
