package cmd

import (
	"archive/zip"
	"os"
	"text/template"

	"github.com/spf13/afero"
	"github.com/spf13/afero/zipfs"
	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var tweetFormat string

const defaultTweetFormat = "{{.FullText}}\n"

// extractCmd represents the extract command.
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extract data",
	Long:  `extract the data included in the archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		fi, err := os.Stat(args[0])
		if err != nil {
			return err
		}

		var fs afero.Fs
		if fi.IsDir() {
			fs = afero.NewBasePathFs(afero.NewOsFs(), args[0])
		} else {
			r, err := zip.OpenReader(args[0])
			if err != nil {
				return err
			}
			defer r.Close()

			fs = zipfs.New(&r.Reader)
		}

		tmpl, err := template.New("tweet").Parse(tweetFormat)
		if err != nil {
			return err
		}

		data := twitwoo.New(fs)
		data.EachTweet(func(t twitwoo.Tweet) error {
			return tmpl.Execute(os.Stdout, t)
		})

		return nil
	},
}

func init() {
	extractCmd.Flags().StringVarP(&tweetFormat, "format", "f", defaultTweetFormat, "format of extracted tweets")
	rootCmd.AddCommand(extractCmd)
}
