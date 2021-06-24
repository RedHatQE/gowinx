// +build windows
package crc

import (
	"fmt"

	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/windows-accesibility-features"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/menu"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/notificationarea"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/interaction"
	wa "github.com/hnakamur/w32uiautomation"
)

const (
	notification_icon_id string = "Codeready Containers"

	menu_id string = "crc"

	ACTION_START       string = "start"
	menuitem_start_id  string = "crc-start"
	ACTION_STOP        string = "stop"
	menuitem_stop_id   string = "crc-stop"
	ACTION_DELETE      string = "delete"
	menuitem_delete_id string = "crc-delete"
	ACTION_EXIT        string = "exit"
	menuitem_exit_id   string = "crc-exit"

	ACTION_COPY_OC_COMMAND                string = "oc-command"
	menuitem_copy_oc_command_id           string = "crc-oc-login-menu"
	ACTION_COPY_OC_COMMAND_DEVELOPER      string = "developer"
	menuitem_copy_oc_command_developer_id string = "crc-oc-login-developer"
	ACTION_COPY_OC_COMMAND_KUBEADMIN      string = "kubeadmin"
	menuitem_copy_oc_command_kubeadmin_id string = "crc-oc-login-kubeadmin"
)

var clickActions map[string]string

var crcMenu *wa.IUIAutomationElement

func init() {
	clickActions = map[string]string{
		ACTION_START:                     menuitem_start_id,
		ACTION_STOP:                      menuitem_stop_id,
		ACTION_DELETE:                    menuitem_delete_id,
		ACTION_EXIT:                      menuitem_exit_id,
		ACTION_COPY_OC_COMMAND:           menuitem_copy_oc_command_id,
		ACTION_COPY_OC_COMMAND_DEVELOPER: menuitem_copy_oc_command_developer_id,
		ACTION_COPY_OC_COMMAND_KUBEADMIN: menuitem_copy_oc_command_kubeadmin_id}
}

func Click(actions []string) error {
	// Initialize base elements
	intialize()

	// Click action
	clickSimpleAction(actions[0], crcMenu)

	if len(actions) == 2 {
		subMenuitem_id, ok := clickActions[actions[0]]
		if !ok {
			return fmt.Errorf("No action defined %s", actions[0])
		}
		subMenu, err := menu.GetMenuItem(crcMenu, subMenuitem_id)
		if err != nil {
			return err
		}
		clickSimpleAction(actions[1], subMenu)
	}
	// Finalize context
	finalize()
	return nil
}

func intialize() error {
	// Initialize context
	win32waf.Initalize()

	// Show notification icon
	notificationarea.ShowHiddenNotificationArea()
	if x, y, err := notificationarea.GetIconPositionByTitle(notification_icon_id); err == nil {
		interaction.Click(int32(x), int32(y))
	}

	// Get crc menu element
	menu, err := menu.GetMenuFromRoot(menu_id)
	if err != nil {
		return err
	}
	crcMenu = menu
	return nil
}

func finalize() {
	// Finalize context
	win32waf.Finalize()
}

func clickSimpleAction(action string, actionMenu *wa.IUIAutomationElement) error {
	menuitem_id, ok := clickActions[action]
	if !ok {
		return fmt.Errorf("No action defined %s", action)
	}
	menuitem, err := menu.GetMenuItem(actionMenu, menuitem_id)
	if err != nil {
		return err
	}
	menuitemPosition, err := menu.GetMenuItemRect(menuitem)
	if err != nil {
		return err
	}
	interaction.ClickOnRect(*menuitemPosition)
	return nil
}
