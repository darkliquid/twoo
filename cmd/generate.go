package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

type generateCfg struct {
	OutDir          string
	Verbose         bool
	PageSize        int
	SortOrder       string
	IncludeReplies  bool
	IncludeRetweets bool
	ExtractOnly     bool
	SkipExtract     bool
}

var gencfg generateCfg

const generateHelp = `Generate a static HTML website of the data included in the archive

generate uses a different strategy to serve to build the same kinf of data.
Rather than operating entirely from the archive, generate first extracts every tweet
to disk and then builds a static HTML website using the extracted data.

This approach allows for more flexibility in how the data is presented, but is
more disk intensive as the data is being duplicated.`

func vlog(args ...any) {
	if gencfg.Verbose {
		log.Println(args...)
	}
}

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate static HTML",
	Long:  generateHelp,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Step 1: Open the archive
		fs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		// Step 2: Init the data parser
		data := twitwoo.New(fs)

		if !gencfg.SkipExtract {
			cnt := int64(0)

			// Step 3: Extract the tweets
			if err = data.EachTweet(func(t *twitwoo.Tweet) error {
				// TODO: Filter whether we want to include the current tweet or not.
				if !gencfg.IncludeReplies && t.InReplayToStatusID != "" {
					return nil
				}

				if !gencfg.IncludeRetweets && strings.HasPrefix(t.FullText, "RT ") {
					return nil
				}

				dir := tweetDir(t)
				if err = os.MkdirAll(dir, 0755); err != nil {
					return err
				}

				fn := fmt.Sprintf("%d.json", t.ID)
				fp := filepath.Join(dir, fn)
				f, ferr := os.Create(fp)
				if ferr != nil {
					return ferr
				}
				defer f.Close()

				if err = json.NewEncoder(f).Encode(t); err != nil {
					return err
				}

				vlog("Writing tweet to", fp)

				cnt++
				return nil
			}); err != nil {
				return err
			}

			vlog("Extracted", cnt, "tweets")
		} else {
			vlog("Skipping tweet extraction")
		}

		if gencfg.ExtractOnly {
			vlog("Skipping HTML generation")
			return nil
		}

		// Step 4: Iterate over the tweets on the file system and build the static HTML
		return nil
	},
}

func tweetDir(t *twitwoo.Tweet) string {
	year, month, day := t.CreatedAt.Date()
	yearStr := fmt.Sprint(year)
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	hour, minute, second := t.CreatedAt.Clock()
	hourStr := fmt.Sprintf("%02d", hour)
	minStr := fmt.Sprintf("%02d", minute)
	secStr := fmt.Sprintf("%02d", second)

	return filepath.Join(gencfg.OutDir, yearStr, monthStr, dayStr, hourStr, minStr, secStr)
}

func init() {
	generateCmd.Flags().StringVarP(&gencfg.OutDir, "out", "o", ".", "where to write the static site to")
	generateCmd.Flags().BoolVarP(&gencfg.Verbose, "verbose", "v", false, "enable verbose output")
	generateCmd.Flags().IntVarP(&gencfg.PageSize, "page-size", "p", 50, "how many tweets to include per page")
	generateCmd.Flags().StringVarP(&gencfg.SortOrder, "sort", "s", "desc", "sort order for tweets (asc or desc)")
	generateCmd.Flags().BoolVarP(&gencfg.IncludeReplies, "include-replies", "r", false, "include replies in the output")
	generateCmd.Flags().BoolVarP(&gencfg.IncludeRetweets, "include-retweets", "t", false, "include retweets in the output")
	generateCmd.Flags().BoolVarP(&gencfg.ExtractOnly, "extract-only", "e", false, "only extract the tweets, don't build the static site")
	generateCmd.Flags().BoolVarP(&gencfg.SkipExtract, "skip-extract", "k", false, "skip the extraction step and only build the static site")

	rootCmd.AddCommand(generateCmd)
}
