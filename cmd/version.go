package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of UP",
	Long:  "All software has versions. This is UP's.",
	RunE: func(cmd *cobra.Command, args []string) error {
		printUpVersion()
		return nil
	},
}

// printUpVersion will be used to print the up's version
func printUpVersion() {
	fmt.Println("UP a uiu programmers platform v0.1")
}
