package cmd

import (
	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/internal/server"
)

var (
	srvbind  string
	cachedir string
)

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a website of the archive",
	Long:  `Serve a website of the data in the twitter archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		fs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		return server.Serve(srvbind, cachedir, fs)
	},
}

func init() {
	serveCmd.Flags().StringVar(&srvbind, "bind", "localhost:3000", "host:port to bind to")
	serveCmd.Flags().StringVar(&cachedir, "cache", "", "directory to cache pages in (cache off by default)")
	rootCmd.AddCommand(serveCmd)
}
