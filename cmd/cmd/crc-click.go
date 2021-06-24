package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adrianriobo/gowinx/pkg/app/crc"
	"github.com/adrianriobo/gowinx/pkg/util/logging"
)

const (
	crcClickCmdName string = "click"
)

func init() {
	crcCmd.AddCommand(crcClickCmd)
}

var crcClickCmd = &cobra.Command{
	Use:   crcClickCmdName,
	Short: "click on crc notification icon element",
	Long:  "click on crc notification icon element",
	RunE: func(cmd *cobra.Command, args []string) error {
		logging.Infof("crc click on %v", args)
		return crc.Click(args)
	},
}
