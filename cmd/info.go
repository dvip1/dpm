package cmd

import (
	"fmt"
	"github.com/dvip1/dpm/internal/client"
	"github.com/spf13/cobra"
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

		resp, err := apiClient.GetPackageByName(pkg)
		if err != nil {
			fmt.Println("Error fetching for the package %s", pkg)
			return
		}
		if resp == (client.Package{}) {
			fmt.Printf("There is no such package")
			return
		}

		fmt.Println("---")
		fmt.Printf("Package:     %s\n", resp.Name)
		fmt.Printf("Description: %s\n", resp.Description)
		fmt.Printf("Repository:  %s\n", resp.RepoURL)
		fmt.Printf("Homepage:    %s\n", resp.Homepage)
		fmt.Printf("License:     %s\n", resp.License)
		fmt.Printf("Version:     %s\n", resp.LatestVersion.Version)
		fmt.Printf("Published:   %s\n", resp.LatestVersion.PublishedAt)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
