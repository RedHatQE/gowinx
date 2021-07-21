package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adrianriobo/gowinx/pkg/app/crc"
	"github.com/adrianriobo/gowinx/pkg/util/logging"
)

const (
	crcStatusCmdName string = "status"
)

func init() {
	crcCmd.AddCommand(crcStatusCmd)
}

var crcStatusCmd = &cobra.Command{
	Use:   crcStatusCmdName,
	Short: "get status",
	Long:  "get status",
	RunE: func(cmd *cobra.Command, args []string) error {
		logging.Infof("crc get status")
		return crc.GetStatus()
	},
}
