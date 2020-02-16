package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	// TEMPLATE_TODO: Update help descriptions
	Use:   "cli",
	Short: "cli is a command line interface for ...",
	Long:  "cli is a command line interface for ...",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
