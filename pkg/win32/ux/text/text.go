// +build windows

package text

import (
	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-accesibility-features"
	wa "github.com/openstandia/w32uiautomation"
)

// Find the first element of type text and try to get its value applying the value pattern
func GetTextElementValue(parentElement *wa.IUIAutomationElement) (string, error) {
	if textElement, err := win32waf.GetElementFromParentByType(parentElement, wa.UIA_TextControlTypeId); err != nil {
		return "", err
	} else {
		return win32waf.GetElementText(textElement)
	}
}
