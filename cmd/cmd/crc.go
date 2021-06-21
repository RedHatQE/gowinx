package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	rootCmd.AddCommand(crcCmd)
	flagSet := pflag.NewFlagSet("crc", pflag.ExitOnError)
	crcCmd.Flags().AddFlagSet(flagSet)
}

var crcCmd = &cobra.Command{
	Use:   "crc",
	Short: "crc notification icon interactions",
	Long:  "crc notification icon interactions",
}
