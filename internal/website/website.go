package website

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

//go:embed templates
var builtinTmpl embed.FS

// Templates is an optional set of templates to use.
var Templates fs.FS

// EnableSearch enables the search functionality.
var EnableSearch = false

// BuiltInTemplates returns the built-in templates.
func BuiltInTemplates() embed.FS {
	return builtinTmpl
}

var tmplCache *template.Template

func templateCache(data *twitwoo.Data) *template.Template {
	if tmplCache != nil {
		return tmplCache
	}

	m, err := data.Manifest()
	if err != nil {
		panic(err)
	}

	tmplCache, err = template.New("twoo").Funcs(FuncMap(m)).ParseFS(builtinTmpl, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	if Templates != nil {
		tmplCache, err = tmplCache.ParseFS(Templates, "*.tmpl")
		if err != nil {
			panic(err)
		}
	}

	return tmplCache
}

type PageData struct {
	Profile    *twitwoo.Profile
	PrevPage   string
	NextPage   string
	UserInfo   twitwoo.UserInfo
	Page       int64
	PageSize   int64
	PageCount  int64
	TweetCount int64
}

// Stylesheet writes the stylesheet to the given writer.
func Stylesheet(data *twitwoo.Data, w io.Writer) error {
	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	pd := PageData{Profile: profiles[0]}

	return templateCache(data).Lookup("stylesheet.tmpl").Execute(w, pd)
}

// Index write a page of multiple items.
func Index(data *twitwoo.Data, page, pageSize int64, w io.Writer) error {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	m, err := data.Manifest()
	if err != nil {
		return err
	}

	totalTweets := int64(0)
	for _, file := range m.DataTypes.Tweets.Files {
		totalTweets += file.Count
	}

	pageCount := totalTweets / pageSize
	pd := PageData{
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TweetCount: totalTweets,
	}
	if page > 1 {
		pd.PrevPage = fmt.Sprintf("/page/%d", page-1)
	}
	if page < pageCount {
		pd.NextPage = fmt.Sprintf("/page/%d", page+1)
	}

	return render(data, pd, nil, w)
}

// Page writes a single item page.
func Page(data *twitwoo.Data, id int64, w io.Writer) error {
	return render(data, PageData{PageSize: 1}, func(t *twitwoo.Tweet) *twitwoo.Tweet {
		if t.ID == id {
			return t
		}
		return nil
	}, w)
}

type filterFunc func(*twitwoo.Tweet) *twitwoo.Tweet

func renderTweets(data *twitwoo.Data, pd PageData, filter filterFunc, w io.Writer) error {
	tweet := templateCache(data).Lookup("tweet.tmpl")

	// TODO:This is an incredibly naive implementation, but for live serving
	// is probably fine. Maybe find some way to index deeper into the json data
	// for tweets to skip having to read and parse everything before the page you
	// actually want to display.
	i := int64(0)
	cnt := int64(0)
	if err := data.EachTweet(func(t *twitwoo.Tweet) error {
		i++

		// If we care about pages but we're before the offset
		// into the tweets we care about, skip.
		if pd.Page > 0 && i <= (pd.Page-1)*pd.PageSize {
			return nil
		}

		// If we have collected all the tweets we want, stop iterating.
		if pd.PageSize > 0 && cnt >= pd.PageSize {
			return twitwoo.ErrBreak
		}

		// If we don't have a selector function, select everything.
		if filter == nil {
			cnt++
			return tweet.Execute(w, t)
		}

		// If the tweet passes the selector function, use it.
		t = filter(t)
		if t != nil {
			cnt++
			return tweet.Execute(w, t)
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func render(data *twitwoo.Data, pd PageData, fn func(*twitwoo.Tweet) *twitwoo.Tweet, w io.Writer) error {
	m, err := data.Manifest()
	if err != nil {
		return err
	}

	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	pd.UserInfo = m.UserInfo
	pd.Profile = profiles[0]

	header := templateCache(data).Lookup("header.tmpl")
	if err = header.Execute(w, pd); err != nil {
		return err
	}

	if err = renderTweets(data, pd, fn, w); err != nil {
		return err
	}

	footer := templateCache(data).Lookup("footer.tmpl")
	return footer.Execute(w, pd)
}

// Content renders a page containing custom content.
func Content(data *twitwoo.Data, pd PageData, fn contentFunc, w io.Writer) error {
	return wrapper(data, pd, fn, w)
}

// Tweet renders a singular tweet without wrapper HTML for the rest of a page.
func Tweet(data *twitwoo.Data, tw *twitwoo.Tweet, w io.Writer) error {
	return templateCache(data).Lookup("tweet.tmpl").Execute(w, tw)
}

type contentFunc func(data *twitwoo.Data, pd PageData, w io.Writer) error

func wrapper(data *twitwoo.Data, pd PageData, fn contentFunc, w io.Writer) error {
	m, err := data.Manifest()
	if err != nil {
		return err
	}

	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	pd.UserInfo = m.UserInfo
	pd.Profile = profiles[0]

	header := templateCache(data).Lookup("header.tmpl")
	if err = header.Execute(w, pd); err != nil {
		return err
	}

	if err = fn(data, pd, w); err != nil {
		return err
	}

	footer := templateCache(data).Lookup("footer.tmpl")
	return footer.Execute(w, pd)
}
