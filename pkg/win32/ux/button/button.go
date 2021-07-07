// +build windows

package button

import (
	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/user-interface/windows-accesibility-features"
	"github.com/adrianriobo/gowinx/pkg/win32/interaction"
	wa "github.com/openstandia/w32uiautomation"
)

func GetButton(parentElement *wa.IUIAutomationElement, name string) (*wa.IUIAutomationElement, error) {
	return win32waf.GetElementFromParent(parentElement, name, wa.UIA_ButtonControlTypeId)
}

func Click(button *wa.IUIAutomationElement) error {
	position, err := win32waf.GetElementRect(button)
	if err != nil {
		return err
	}
	return interaction.ClickOnRect(*position)
}
