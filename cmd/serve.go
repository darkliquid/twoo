package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a website of the archive",
	Long:  `Serve a website of the data in the twitter archive`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
