package cmd

import (
	"github.com/spf13/cobra"
)

// storagetransfertransferJobsCmd represents the transferJobs command
var storagetransfertransferJobsCmd = &cobra.Command{
	Use:   "transferJobs",
	Short: "transferJobs",
	Long:  `transferJobs`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	storagetransferCmd.AddCommand(storagetransfertransferJobsCmd)
}
