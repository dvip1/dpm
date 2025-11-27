package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"bufio"
	"github.com/dvip1/dpm/internal/client"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: fmt.Sprintf("Installs a package"),
	Long:  fmt.Sprintf("Installs a package"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := client.NewClient(client.AppRoutes.BaseUrl)
		pkg := args[0]
		fmt.Printf("Fetching Information for: %s\n", pkg)
		resp, err := apiClient.GetPackageByName(pkg);

		if err!=nil {
			fmt.Println("Error fetching for the package %s", pkg )
			return;
		}

		if resp ==(client.Package{}){
			fmt.Println("There is no such package");
			return;
		}
		fmt.Println("---")
		fmt.Printf("Package:     %s\n", resp.Name)
		fmt.Printf("Description: %s\n", resp.Description)
		fmt.Printf("Repository:  %s\n", resp.RepoURL)
		fmt.Printf("Homepage:    %s\n", resp.Homepage)
		fmt.Printf("License:     %s\n", resp.License)
		fmt.Printf("Version:     %s\n", resp.LatestVersion.Version)
		fmt.Printf("Published:   %s\n", resp.LatestVersion.PublishedAt)
		fmt.Printf("Source:      %s\n", resp.LatestVersion.SourceURL)
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nDo you want to install this package? (y/N): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" || input == "yes"  || input =="Y"{
			fmt.Println("Installing...")
			// call your install logic here, e.g., apiClient.InstallPackage(pkg)
		} else {
			fmt.Println("Installation cancelled.")
			return
		}
		
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
