package commands

import (
	"fmt"

	// TEMPLATE_TODO: update spec package path
	"github.com/jasonpilz/go-cli/spec"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version",
	Long:  fmt.Sprintf("Show the current version of %s", spec.Repo),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(spec.AppVersion.Complete(&spec.GithubLatestVersioner{}))
	},
}
