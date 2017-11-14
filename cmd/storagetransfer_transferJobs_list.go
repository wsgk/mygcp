package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storagetransfer/v1"
)

// storagetransfertransferJobslistCmd represents the list command
var storagetransfertransferJobslistCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists transfer jobs.",
	Long:  `Lists transfer jobs.`,
	Run:   storageTransferTransferJobsList,
}

func init() {
	storagetransfertransferJobsCmd.AddCommand(storagetransfertransferJobslistCmd)

	storagetransfertransferJobslistCmd.Flags().StringP("filter", "f", "", "Filter")
	viper.BindPFlag("filter", storagetransfertransferJobslistCmd.Flags().Lookup("filter"))
}

func storageTransferTransferJobsList(*cobra.Command, []string) {
	filter := viper.GetString("filter")
	if filter == "" {
		fmt.Println("filter is required")
		os.Exit(1)
	}

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, storagetransfer.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	storagetransferService, err := storagetransfer.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := storagetransferService.TransferJobs.List()
	req.Filter(filter)
	var transferJobs []*storagetransfer.TransferJob
	if err := req.Pages(ctx, func(page *storagetransfer.ListTransferJobsResponse) error {
		for _, transferJob := range page.TransferJobs {
			transferJobs = append(transferJobs, transferJob)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	dataRaw, _ := json.MarshalIndent(transferJobs, "", "    ")
	fmt.Fprintln(os.Stdout, string(dataRaw))
}
