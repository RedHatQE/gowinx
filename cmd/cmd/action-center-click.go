package cmd

import (
	"github.com/spf13/cobra"

	actionCenter "github.com/RedHatQE/gowinx/pkg/app/action-center"
	"github.com/RedHatQE/gowinx/pkg/util/logging"
)

const (
	actionCenterClickCmdName string = "click"
)

func init() {
	actionCenterCmd.AddCommand(actionCenterClickCmd)
}

var actionCenterClickCmd = &cobra.Command{
	Use:   actionCenterClickCmdName,
	Short: "click on action center button from notification area",
	Long:  "click on action center button from notification area",
	RunE: func(cmd *cobra.Command, args []string) error {
		logging.Infof("action center click ")
		return actionCenter.ClickNotifyButton()
	},
}
