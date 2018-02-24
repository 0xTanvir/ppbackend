package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of PP",
	Long:  "All software has versions. This is PP's.",
	RunE: func(cmd *cobra.Command, args []string) error {
		printPpVersion()
		return nil
	},
}

// printPpVersion will be used to print the pp's version
func printPpVersion() {
	fmt.Println("Programmers Playground v0.1")
}
