// +build windows
package crc

import (
	"fmt"

	win32waf "github.com/adrianriobo/gowinx/pkg/win32/api/windows-accesibility-features"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/menu"
	"github.com/adrianriobo/gowinx/pkg/win32/desktop/notificationarea"
	"github.com/adrianriobo/gowinx/pkg/win32/ux/interaction"
)

const (
	notification_icon_id string = "Codeready Containers"

	menu_id string = "crc"

	ACTION_START       string = "START"
	menuitem_start_id  string = "crc-start"
	ACTION_STOP        string = "STOP"
	menuitem_stop_id   string = "crc-stop"
	ACTION_DELETE      string = "DELETE"
	menuitem_delete_id string = "crc-delete"
	ACTION_EXIT        string = "EXIT"
	menuitem_exit_id   string = "crc-exit"

	ACTION_COPY_OC_COMMAND                string = "COPY_OC_COMMAND"
	menuitem_copy_oc_command_id           string = "copy-oc-command"
	ACTION_COPY_OC_COMMAND_DEVELOPER      string = "COPY_OC_COMMAND_DEVELOPER"
	menuitem_copy_oc_command_developer_id string = "copy-oc-command-developer"
	ACTION_COPY_OC_COMMAND_KUBEADMIN      string = "COPY_OC_COMMAND_KUBEADMIN"
	menuitem_copy_oc_command_kubeadmin_id string = "copy-oc-command-kubeadmin"
)

var simpleClickActions map[string]string
var doubleClickActions map[string]string

func init() {
	simpleClickActions = map[string]string{
		ACTION_START:                     menuitem_start_id,
		ACTION_STOP:                      menuitem_stop_id,
		ACTION_DELETE:                    menuitem_delete_id,
		ACTION_EXIT:                      menuitem_exit_id,
		ACTION_COPY_OC_COMMAND_DEVELOPER: menuitem_copy_oc_command_developer_id,
		ACTION_COPY_OC_COMMAND_KUBEADMIN: menuitem_copy_oc_command_kubeadmin_id}

	doubleClickActions = map[string]string{
		ACTION_COPY_OC_COMMAND: menuitem_copy_oc_command_id}
}

func Click(action string) {
	// Click on icon from notification area
	notificationarea.ShowHiddenNotificationArea()
	if x, y, err := notificationarea.GetIconPositionByTitle(notification_icon_id); err == nil {
		interaction.Click(int32(x), int32(y))
	}

	// With the menu displed click on action
	if checkDoubleAction(action) {
		clickDoubleAction(action)
	} else {
		clickSimpleAction(action)
	}
}

func clickSimpleAction(action string) error {
	menuitem_id, ok := simpleClickActions[action]
	if !ok {
		return fmt.Errorf("No action defined %s", action)
	}
	win32waf.Initalize()
	crcMenu, err := menu.GetMenuFromRoot(menu_id)
	if err != nil {
		return err
	}
	menuitem, err := menu.GetMenuItem(crcMenu, menuitem_id)
	if err != nil {
		return err
	}
	menuitemPosition, err := menu.GetMenuItemRect(menuitem)
	if err != nil {
		return err
	}
	interaction.ClickOnRect(*menuitemPosition)
	win32waf.Finalize()
	return nil
}

func clickDoubleAction(action string) {

}

func checkDoubleAction(action string) bool {
	_, ok := doubleClickActions[action]
	return ok
}
