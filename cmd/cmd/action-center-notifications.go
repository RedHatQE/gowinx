package cmd

import (
	"github.com/spf13/cobra"

	actionCenter "github.com/adrianriobo/gowinx/pkg/app/action-center"
	"github.com/adrianriobo/gowinx/pkg/util/logging"
)

const (
	actionCenterNotificationsCmdName string = "notifications"
)

func init() {
	actionCenterCmd.AddCommand(actionCenterNotificationsCmd)
}

var actionCenterNotificationsCmd = &cobra.Command{
	Use:   actionCenterNotificationsCmdName,
	Short: "get notifications",
	Long:  "get notificatons",
	RunE: func(cmd *cobra.Command, args []string) error {
		logging.Infof("action center notifications")
		return actionCenter.PrintNotifications("CodeReady Containers")
	},
}
