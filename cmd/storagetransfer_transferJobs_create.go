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

// storagetransfertransferJobscreateCmd represents the create command
var storagetransfertransferJobscreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a transfer job that runs periodically.",
	Long:  `Creates a transfer job that runs periodically.`,
	Run:   storageTransferTransferJobsCreate,
}

func init() {
	storagetransfertransferJobsCmd.AddCommand(storagetransfertransferJobscreateCmd)

	storagetransfertransferJobscreateCmd.Flags().StringP("body", "b", "", "Request body")
	viper.BindPFlag("body", storagetransfertransferJobscreateCmd.Flags().Lookup("body"))
}

func storageTransferTransferJobsCreate(*cobra.Command, []string) {
	body := viper.GetString("body")
	if body == "" {
		fmt.Println("body is required")
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

	var transferJob storagetransfer.TransferJob
	json.Unmarshal([]byte(body), &transferJob)
	rb := &transferJob

	resp, err := storagetransferService.TransferJobs.Create(rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	dataRaw, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Fprintln(os.Stdout, string(dataRaw))
}
