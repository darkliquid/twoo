package extract

import (
	"github.com/spf13/cobra"
)

// extractCmd represents the extract command.
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extract data",
	Long:  `extract the data included in the archive`,
	Args:  cobra.ExactArgs(1),
}

func init() {
	extractCmd.AddCommand(
		manifestCmd,
		tweetsCmd,
		acipCmd,
	)
}

// Command returns the extract command.
func Command() *cobra.Command {
	return extractCmd
}
