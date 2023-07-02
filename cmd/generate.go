package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate static HTML",
	Long:  `Generate a static HTML website of the data included in the archive`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
