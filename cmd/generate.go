package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gosimple/slug"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/internal/website"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

type generateCfg struct {
	OutDir          string
	SortOrder       string
	TemplateDir     string
	SubDir          string
	PageSize        int
	IncludeRetweets bool
	IncludeReplies  bool
	ExtractOnly     bool
	SkipExtract     bool
	SkipDetails     bool
	SkipCleanup     bool
	SearchIndex     bool
	MakeTagPages    bool
	Verbose         bool
}

var gencfg generateCfg

const (
	generateHelp = `Generate a static HTML website of the data included in the archive

generate uses a different strategy to serve to build the same kinf of data.
Rather than operating entirely from the archive, generate first extracts every tweet
to disk and then builds a static HTML website using the extracted data.

This approach allows for more flexibility in how the data is presented, but is
more disk intensive as the data is being duplicated.`
	defaultPageSize = 20
	outDirMode      = 0o755
)

func vlog(args ...any) {
	if gencfg.Verbose {
		log.Println(args...)
	}
}

func extractTweet(searchIdx afero.File, t *twitwoo.Tweet) (int64, int64, string, error) {
	replies := int64(0)
	retweets := int64(0)
	if t.InReplyToStatusID > 0 {
		// TODO: handle threads separately.
		replies++
		if !gencfg.IncludeReplies {
			vlog("Skipping reply", t.ID)
			return replies, retweets, "", nil
		}
	}

	if strings.HasPrefix(t.FullText, "RT ") {
		retweets++
		if !gencfg.IncludeRetweets {
			vlog("Skipping retweet", t.ID)
			return replies, retweets, "", nil
		}
	}

	dir := tweetDir(t)
	if err := os.MkdirAll(dir, outDirMode); err != nil {
		return replies, retweets, "", err
	}

	// ensure the tweet ID is 20 characters long for easier sorting
	fn := fmt.Sprintf("%020d.json", t.ID)
	fp := filepath.Join(dir, fn)
	f, ferr := os.Create(fp)
	if ferr != nil {
		return replies, retweets, "", ferr
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(t); err != nil {
		return replies, retweets, "", err
	}

	if gencfg.SearchIndex {
		if err := json.NewEncoder(searchIdx).Encode(t.SearchIndex()); err != nil {
			return replies, retweets, "", err
		}
		if _, err := searchIdx.WriteString(","); err != nil {
			return replies, retweets, "", err
		}
	}

	fp = strings.TrimPrefix(fp, gencfg.OutDir)
	vlog("Writing tweet to", fp)

	return replies, retweets, fp, nil
}

func genExtractTweets(outfs afero.Fs, data *twitwoo.Data) (map[string][]string, []string, error) {
	var files []string
	tagPages := make(map[string][]string)
	replies := int64(0)
	retweets := int64(0)

	var searchIdx afero.File
	if gencfg.SearchIndex {
		website.EnableSearch = true
		var err error
		searchIdx, err = outfs.Create("search.json")
		if err != nil {
			return nil, nil, err
		}
		defer searchIdx.Close()

		if _, err = searchIdx.WriteString("["); err != nil {
			return nil, nil, err
		}
	}

	if err := data.EachTweet(func(t *twitwoo.Tweet) error {
		reply, retweet, fp, err := extractTweet(searchIdx, t)
		replies += reply
		retweets += retweet
		if fp != "" {
			files = append(files, fp)
			if gencfg.MakeTagPages {
				for _, tag := range t.Hashtags {
					tag = slug.Make(tag)
					tagPages[tag] = append(tagPages[tag], fp)
				}
			}
		}
		return err
	}); err != nil {
		return nil, nil, err
	}

	if gencfg.SearchIndex {
		if _, err := searchIdx.Seek(-1, io.SeekCurrent); err != nil {
			return nil, nil, err
		}
		if _, err := searchIdx.WriteString("]"); err != nil {
			return nil, nil, err
		}
	}

	vlogExtracted(int64(len(files)), replies, retweets)

	return tagPages, files, nil
}

func vlogExtracted(tweets, replies, retweets int64) {
	vlog("Extracted", tweets, "tweets")

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
}

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate static HTML",
	Long:  generateHelp,
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		// Open the archive
		afs, closer, err := util.Open(args[0])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		if err = os.MkdirAll(gencfg.OutDir, outDirMode); err != nil {
			return err
		}

		// Setup the output directory
		outfs := afero.NewBasePathFs(afero.NewOsFs(), gencfg.OutDir)

		// Init the data parser
		data := twitwoo.New(afs)

		// Extract the tweets
		var files []string
		var tagPages map[string][]string
		if !gencfg.SkipExtract {
			if tagPages, files, err = genExtractTweets(outfs, data); err != nil {
				return err
			}
		} else {
			vlog("Skipping tweet extraction")
		}

		if gencfg.ExtractOnly {
			vlog("Skipping HTML generation and media extraction")
			return nil
		}

		if gencfg.TemplateDir != "" {
			tfs := os.DirFS(gencfg.TemplateDir)
			website.Templates = tfs
		}

		if gencfg.MakeTagPages {
			website.UseTagIndex = true
		}

		if gencfg.SubDir != "" {
			website.SubDir = path.Join("/", gencfg.SubDir)
		}

		// Generate the Stylesheet
		sf, err := outfs.Create("/stylesheet.css")
		if err != nil {
			return err
		}
		defer sf.Close()

		if err = website.Stylesheet(data, sf); err != nil {
			return err
		}

		// Extract Profile Media
		m, err := data.Manifest()
		if err != nil {
			return err
		}

		p, err := data.Profiles()
		if err != nil {
			return err
		}

		header := website.ProfileMediaPath(m.UserInfo.AccountID, p[0].Header)
		if header != "" {
			if err = copyPath(afs, outfs, header+".jpg", header+".jpg"); err != nil {
				return err
			}
		}

		avatar := website.ProfileMediaPath(m.UserInfo.AccountID, p[0].Avatar)
		if avatar != "" {
			if err = copyPath(afs, outfs, avatar, avatar); err != nil {
				return err
			}
		}

		// Iterate over the tweets on the file system if we haven't already
		// determined them via extraction.
		if len(files) == 0 {
			if err = filepath.WalkDir(gencfg.OutDir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if path == gencfg.OutDir {
					return nil
				}

				if d.IsDir() {
					return nil
				}

				if strings.HasSuffix(path, ".json") {
					files = append(files, strings.TrimPrefix(path, gencfg.OutDir))
				}
				return nil
			}); err != nil {
				return err
			}
		}

		// Sort the files based on sort order (they are lexographically sorted by default
		// which will result is ascending chronological order due to the naming conventions).
		if gencfg.SortOrder == "desc" {
			sort.Sort(sort.Reverse(sort.StringSlice(files)))
		}

		// Create detail pages for each tweet
		if !gencfg.SkipDetails {
			for i, f := range files {
				var prev, next string
				if i > 0 {
					prev = files[i-1]
				}
				if i < len(files)-1 {
					next = files[i+1]
				}
				if err = genDetailsPage(outfs, data, f, prev, next); err != nil {
					return err
				}
			}
		} else {
			vlog("Skipping detail page generation")
		}

		// Generate tag index pages
		if gencfg.MakeTagPages {
			for tag, pages := range tagPages {
				if err = genIndexPages(afs, outfs, data, pages, path.Join("tag", tag)); err != nil {
					return err
				}
			}
		}

		if err = genIndexPages(afs, outfs, data, files, ""); err != nil {
			return err
		}

		// Cleanup the output directory
		if !gencfg.SkipCleanup {
			for _, f := range files {
				vlog("Removing", f)
				if err = outfs.Remove(f); err != nil {
					vlog("Failed to remove", f, ":", err)
				}
			}
		}

		return nil
	},
}

func genIndexPages(afs, outfs afero.Fs, data *twitwoo.Data, files []string, prefix string) error {
	// Group files into pages based on page size
	pageFiles := make([]string, len(files))
	copy(pageFiles, files)

	pageNum := int64(1)
	for len(pageFiles) > 0 {
		var page []string
		if len(pageFiles) > gencfg.PageSize {
			page = pageFiles[:gencfg.PageSize]
			pageFiles = pageFiles[gencfg.PageSize:]
		} else {
			page = pageFiles
			pageFiles = []string{}
		}

		// Generate the HTML for the main feed pages
		if err := genIndexPage(afs, outfs, data, page, prefix, pageNum, len(pageFiles) > 0); err != nil {
			return err
		}
		pageNum++
	}
	return nil
}

func decodeTweet(fs afero.Fs, fn string) (*twitwoo.Tweet, error) {
	var tweetFile afero.File
	tweetFile, err := fs.Open(fn)
	if err != nil {
		return nil, err
	}
	defer tweetFile.Close()

	var tweet twitwoo.Tweet
	if err = json.NewDecoder(tweetFile).Decode(&tweet); err != nil {
		return nil, err
	}
	return &tweet, nil
}

func genDetailsPage(fs afero.Fs, data *twitwoo.Data, fn, fnPrev, fnNext string) error {
	vlog("Generating details page for", fn)

	tweet, err := decodeTweet(fs, fn)
	if err != nil {
		return err
	}

	pd := website.PageData{}

	if fnPrev != "" {
		vlog("Previous tweet:", fnPrev)
		var prevTweet *twitwoo.Tweet
		prevTweet, err = decodeTweet(fs, fnPrev)
		if err != nil {
			return err
		}
		y, m, d := prevTweet.CreatedAt.Date()
		pd.PrevPage = path.Join("/", gencfg.SubDir, fmt.Sprintf("/%d/%02d/%02d/%020d", y, m, d, prevTweet.ID))
	}

	if fnNext != "" {
		vlog("Next tweet:", fnNext)
		var nextTweet *twitwoo.Tweet
		nextTweet, err = decodeTweet(fs, fnNext)
		if err != nil {
			return err
		}
		y, m, d := nextTweet.CreatedAt.Date()
		pd.NextPage = path.Join("/", gencfg.SubDir, fmt.Sprintf("/%d/%02d/%02d/%020d", y, m, d, nextTweet.ID))
	}

	y, m, d := tweet.CreatedAt.Date()
	outname := filepath.Join(
		"/",
		fmt.Sprintf("%d", y),
		fmt.Sprintf("%02d", m),
		fmt.Sprintf("%02d", d),
		fmt.Sprintf("%020d", tweet.ID),
		"index.html",
	)

	if err = fs.MkdirAll(filepath.Dir(outname), outDirMode); err != nil {
		return err
	}

	vlog("Creating details file", outname)

	f, err := fs.Create(outname)
	if err != nil {
		return err
	}
	defer f.Close()

	vlog("Rendering details page for tweet", tweet.ID)

	return website.Content(data, pd, func(data *twitwoo.Data, _ website.PageData, w io.Writer) error {
		return website.Tweet(data, tweet, w)
	}, f)
}

func copyPath(srcfs, destfs afero.Fs, src, dest string) error {
	srcf, err := srcfs.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()

	if err = destfs.MkdirAll(filepath.Dir(dest), outDirMode); err != nil {
		return err
	}

	destf, err := destfs.Create(dest)
	if err != nil {
		return err
	}
	defer destf.Close()

	_, err = io.Copy(destf, srcf)
	return err
}

func genExtractMedia(srcfs, destfs afero.Fs, tweet *twitwoo.Tweet) error {
	for _, m := range tweet.Media {
		if m.SourceStatusID > 0 && m.SourceStatusID != tweet.ID {
			vlog("Skipping media entry", m.ID, "because it's from a retweet")
			continue
		}
		vlog("Extracting media entry", m.ID)
		path := website.TweetMediaPath(tweet.ID, m)
		if err := copyPath(srcfs, destfs, path, path); err != nil {
			return err
		}
	}
	return nil
}

func genIndexPage(
	srcfs, fs afero.Fs,
	data *twitwoo.Data,
	fns []string,
	prefix string,
	pageNum int64,
	hasNext bool,
) error {
	vlog("Generating index page", prefix, pageNum, "with", len(fns), "tweets")

	outname := filepath.Join(prefix, "index.html")
	if pageNum > 1 {
		outname = filepath.Join(prefix, "page", strconv.FormatInt(pageNum, 10), "index.html")
	}

	if err := fs.MkdirAll(filepath.Dir(outname), outDirMode); err != nil {
		return err
	}

	f, err := fs.Create(outname)
	if err != nil {
		return err
	}
	defer f.Close()

	pd := website.PageData{
		Page:     pageNum,
		PageSize: int64(gencfg.PageSize),
	}

	if pageNum == 2 { //nolint:gomnd // 2 is the second page
		pd.PrevPage = path.Join("/", gencfg.SubDir, prefix, "/")
	} else if pageNum > 1 {
		pd.PrevPage = path.Join("/", gencfg.SubDir, prefix, fmt.Sprintf("/page/%d", pageNum-1))
	}

	if hasNext {
		pd.NextPage = path.Join("/", gencfg.SubDir, prefix, fmt.Sprintf("/page/%d", pageNum+1))
	}

	return website.Content(data, pd, func(data *twitwoo.Data, _ website.PageData, w io.Writer) error {
		for _, fn := range fns {
			var tweet *twitwoo.Tweet
			tweet, err = decodeTweet(fs, fn)
			if err != nil {
				return err
			}

			// Extract tweet media
			if err = genExtractMedia(srcfs, fs, tweet); err != nil {
				return err
			}

			// Render tweet
			if err = website.Tweet(data, tweet, w); err != nil {
				return err
			}
		}
		return nil
	}, f)
}

func tweetDir(t *twitwoo.Tweet) string {
	year, month, day := t.CreatedAt.Date()
	yearStr := fmt.Sprint(year)
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)

	return filepath.Join(gencfg.OutDir, yearStr, monthStr, dayStr)
}

func init() {
	initGenFlags()
	initGenSubcommands()
	rootCmd.AddCommand(generateCmd)
}

func initGenFlags() {
	generateCmd.Flags().StringVarP(
		&gencfg.OutDir,
		"out",
		"o",
		".",
		"where to write the static site to",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.Verbose,
		"verbose",
		"v",
		false,
		"enable verbose output",
	)
	generateCmd.Flags().IntVarP(
		&gencfg.PageSize,
		"page-size",
		"p",
		defaultPageSize,
		"how many tweets to include per page",
	)
	generateCmd.Flags().StringVarP(
		&gencfg.SortOrder,
		"sort",
		"s",
		"desc",
		"sort order for tweets (asc or desc)",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.IncludeReplies,
		"include-replies",
		"r",
		false,
		"include replies in the output",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.IncludeRetweets,
		"include-retweets",
		"t",
		false,
		"include retweets in the output",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.ExtractOnly,
		"extract-only",
		"e",
		false,
		"only extract the tweets, don't build the static site",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.SkipExtract,
		"skip-extract",
		"k",
		false,
		"skip the extraction step and only build the static site",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.SkipCleanup,
		"skip-cleanup",
		"c",
		false,
		"skip cleaning up the output directory after generating",
	)
	generateCmd.Flags().StringVarP(
		&gencfg.TemplateDir,
		"template-dir",
		"m",
		"",
		"directory containing custom templates",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.SearchIndex,
		"search-index",
		"i",
		false,
		"generate a tinysearch index",
	)
	generateCmd.Flags().BoolVarP(
		&gencfg.MakeTagPages,
		"tag-pages",
		"g",
		false,
		"generate hashtag indexes",
	)
	generateCmd.Flags().StringVarP(
		&gencfg.SubDir,
		"sub-dir",
		"d",
		"",
		"sub dir for the generated site",
	)
}

func initGenSubcommands() {
	generateCmd.AddCommand(&cobra.Command{
		Use: "templates [template output dir]",
		Long: generateCmd.UsageString() +
			"\nAvailable templates:\n  header.tmpl\n  tweet.tmpl\n  footer.tmpl\n  stylesheet.tmpl\n",
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				err := cmd.Usage()
				fmt.Println(cmd.Long)
				return err
			}

			// Write out the embedded templates so they can be used as a base to customise.
			efs := website.BuiltInTemplates()
			err := fs.WalkDir(efs, ".", func(fp string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if d.IsDir() {
					return nil
				}

				f, err := efs.Open(fp)
				if err != nil {
					return err
				}
				defer f.Close()

				outfile, err := os.Create(filepath.Join(args[0], path.Base(fp)))
				if err != nil {
					return err
				}
				defer outfile.Close()

				_, err = io.Copy(outfile, f)
				return err
			})

			return err
		},
	})
}
