package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/dvip1/dpm/internal/client"
	"log"
	"io"
)

var infoCmd = &cobra.Command{
	Use:   "info [package]",
	Short: fmt.Sprintf("Show package info"),
	Long:  fmt.Sprintf("Show package info"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := client.NewClient(client.AppRoutes.BaseUrl)
		pkg := args[0]
		fmt.Printf("Fetching Information for: %s\n", pkg)
		resp, err := apiClient.GetPackageByName(pkg);
		if err!=nil {
			log.Fatalf("Error fetching for the package %s", pkg )
		}
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		fmt.Println("--- Package Info ---")
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
