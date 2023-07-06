package extract

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var tweetFormat string

const defaultTweetFormat = "{{.FullText}}\n"

// tweetsCmd represents the extract command.
var tweetsCmd = &cobra.Command{
	Use:   "tweets FILE|DIR",
	Short: "extract tweets data",
	Long:  `extract the tweets data included in the archive`,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		fs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		tmpl, err := template.New("tweet").Parse(tweetFormat)
		if err != nil {
			return err
		}

		data := twitwoo.New(fs)
		return data.EachTweet(func(t *twitwoo.Tweet) error {
			return tmpl.Execute(os.Stdout, t)
		})
	},
}

func init() {
	tweetsCmd.Flags().StringVarP(&tweetFormat, "format", "f", defaultTweetFormat, "format of extracted tweets")
}
