package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

type generateCfg struct {
	OutDir          string
	SortOrder       string
	PageSize        int
	Verbose         bool
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

func genExtractTweets(data *twitwoo.Data) ([]string, error) {
	var files []string
	replies := int64(0)
	retweets := int64(0)

	if err := data.EachTweet(func(t *twitwoo.Tweet) error {
		if t.InReplyToStatusID != "" {
			// TODO: handle threads separately.
			replies++
			if !gencfg.IncludeReplies {
				vlog("Skipping reply", t.ID)
				return nil
			}
		}

		if strings.HasPrefix(t.FullText, "RT ") {
			retweets++
			if !gencfg.IncludeRetweets {
				vlog("Skipping retweet", t.ID)
				return nil
			}
		}

		dir := tweetDir(t)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		// ensure the tweet ID is 20 characters long for easier sorting
		fn := fmt.Sprintf("%020d.json", t.ID)
		fp := filepath.Join(dir, fn)
		f, ferr := os.Create(fp)
		if ferr != nil {
			return ferr
		}
		defer f.Close()

		if err := json.NewEncoder(f).Encode(t); err != nil {
			return err
		}

		vlog("Writing tweet to", fp)

		files = append(files, fp)
		return nil
	}); err != nil {
		return nil, err
	}

	vlog("Extracted", len(files), "tweets")

	if gencfg.IncludeReplies {
		vlog("Included", replies, "replies")
	} else {
		vlog("Excluded", replies, "replies")
	}

	if gencfg.IncludeRetweets {
		vlog("Included", retweets, "retweets")
	} else {
		vlog("Excluded", retweets, "retweets")
	}

	return files, nil
}

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate static HTML",
	Long:  generateHelp,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Step 1: Open the archive
		afs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		// Step 2: Init the data parser
		data := twitwoo.New(afs)

		// Step 3: Extract the tweets
		var files []string
		if !gencfg.SkipExtract {
			if files, err = genExtractTweets(data); err != nil {
				return err
			}
		} else {
			vlog("Skipping tweet extraction")
		}

		if gencfg.ExtractOnly {
			vlog("Skipping HTML generation")
			return nil
		}

		// Step 4: Iterate over the tweets on the file system if we haven't already
		// determined them via extraction.
		if len(files) == 0 {
			if err = filepath.WalkDir(gencfg.OutDir, func(path string, d fs.DirEntry, err error) error {
				if path == gencfg.OutDir {
					return nil
				}

				if d.IsDir() {
					return nil
				}

				if strings.HasSuffix(path, ".json") {
					files = append(files, path)
				}
				return nil
			}); err != nil {
				return err
			}
		}

		// Step 5: Sort the files based on sort order (they are lexographically sorted by default
		// which will result is ascending chronological order due to the naming conventions).
		if gencfg.SortOrder == "desc" {
			sort.Sort(sort.Reverse(sort.StringSlice(files)))
		}

		// Step 6: Create detail pages for each tweet
		for _, f := range files {
			if err = genDetailsPage(f); err != nil {
				return err
			}
		}

		// Step 7: Group files into pages based on page size
		pageNum := int64(1)
		for len(files) > 0 {
			var page []string
			if len(files) > gencfg.PageSize {
				page = files[:gencfg.PageSize]
				files = files[gencfg.PageSize:]
			} else {
				page = files
				files = []string{}
			}

			// Step 8: Generate the HTML for the main feed pages
			if err := genIndexPage(page, pageNum, len(files) == 0); err != nil {
				return err
			}
			pageNum++
		}

		return nil
	},
}

func genDetailsPage(fn string) error {
	vlog("Generating details page for", fn)
	return nil
}

func genIndexPage(fns []string, pageNum int64, hasNext bool) error {
	vlog("Generating index page", pageNum, "with", len(fns), "tweets")
	return nil
}

func tweetDir(t *twitwoo.Tweet) string {
	year, month, day := t.CreatedAt.Date()
	yearStr := fmt.Sprint(year)
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)

	return filepath.Join(gencfg.OutDir, yearStr, monthStr, dayStr)
}

func init() {
	generateCmd.Flags().StringVarP(&gencfg.OutDir, "out", "o", ".", "where to write the static site to")
	generateCmd.Flags().BoolVarP(&gencfg.Verbose, "verbose", "v", false, "enable verbose output")
	generateCmd.Flags().IntVarP(&gencfg.PageSize, "page-size", "p", 20, "how many tweets to include per page")
	generateCmd.Flags().StringVarP(&gencfg.SortOrder, "sort", "s", "desc", "sort order for tweets (asc or desc)")
	generateCmd.Flags().BoolVarP(&gencfg.IncludeReplies, "include-replies", "r", false, "include replies in the output")
	generateCmd.Flags().BoolVarP(&gencfg.IncludeRetweets, "include-retweets", "t", false, "include retweets in the output")
	generateCmd.Flags().BoolVarP(&gencfg.ExtractOnly, "extract-only", "e", false, "only extract the tweets, don't build the static site")
	generateCmd.Flags().BoolVarP(&gencfg.SkipExtract, "skip-extract", "k", false, "skip the extraction step and only build the static site")

	rootCmd.AddCommand(generateCmd)
}
