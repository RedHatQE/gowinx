package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	actionCenterCmdName string = "action-center"
)

func init() {
	rootCmd.AddCommand(actionCenterCmd)
	flagSet := pflag.NewFlagSet(actionCenterCmdName, pflag.ExitOnError)
	actionCenterCmd.Flags().AddFlagSet(flagSet)
}

var actionCenterCmd = &cobra.Command{
	Use:   actionCenterCmdName,
	Short: "manage action center from notification area",
	Long:  "manage action center from notification area",
}
