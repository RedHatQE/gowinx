package cmd

import (
	"github.com/spf13/cobra"

	actionCenter "github.com/RedHatQE/gowinx/pkg/app/action-center"
	"github.com/RedHatQE/gowinx/pkg/util/logging"
)

const (
	actionCenterClearAllCmdName string = "clear-all"
)

func init() {
	actionCenterCmd.AddCommand(actionCenterClearAllCmd)
}

var actionCenterClearAllCmd = &cobra.Command{
	Use:   actionCenterClearAllCmdName,
	Short: "clear all notifications from action center",
	Long:  "clear all notifications from action center",
	RunE: func(cmd *cobra.Command, args []string) error {
		logging.Infof("action center clear all ")
		return actionCenter.ClearNotifications()
	},
}
