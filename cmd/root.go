package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dpm",
	Short: "dpm or Dvip's Package Manager is a cli Package Manager for Dvip's User Repository",
	Long:  "dpm or Dvip's Package Manger is a cli tool for managing Packages for Dvip's User Repository which runs across all the Operating Systems",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing dpm '%s'\n", err)
		os.Exit(1)
	}
}
