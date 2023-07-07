package cmd

import (
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/cmd/website"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var srvbind string

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a website of the archive",
	Long:  `Serve a website of the data in the twitter archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		data := twitwoo.New(fs)

		http.ListenAndServe(srvbind, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			page, _ := strconv.ParseInt(q.Get("page"), 10, 64)
			pageSize, _ := strconv.ParseInt(q.Get("page_size"), 10, 64)

			if err := website.Page(data, page, pageSize, w); err != nil {
				log.Fatal(err)
			}
		}))

		return nil
	},
}

func init() {
	serveCmd.Flags().StringVar(&srvbind, "bind", "localhost:3000", "host:port to bind to")
	rootCmd.AddCommand(serveCmd)
}
