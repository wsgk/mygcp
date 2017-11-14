package cmd

import (
	"github.com/spf13/cobra"
)

// storagetransferCmd represents the storagetransfer command
var storagetransferCmd = &cobra.Command{
	Use:   "storagetransfer",
	Short: "Google Storage Transfer",
	Long:  `Google Storage Transfer`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(storagetransferCmd)
}
