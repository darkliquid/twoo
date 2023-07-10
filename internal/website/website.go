package website

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var pageIndexHeaderTmpl = `<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>@{{ .UserInfo.UserName }}{{ if gt .PageCount 0 }} - {{ .Page }}/{{ .PageCount }}{{ end }}</title>
	<link rel="stylesheet" href="http://markdowncss.github.io/retro/css/retro.css">
	<style>
		body {
			margin: 0;
			max-width: 100%;
			padding: 0;
		}

		header {
			{{ $profile_header_url := profile_header_url .Profile }}
			{{ with $profile_header_url }}
			background-image: url({{ . }});
			{{ end }}
			background-repeat: no-repeat;
			background-size: cover;
			background-position: center;
			margin-bottom: 1em;
			padding-top: 5em;
		}

		header h1 {
			max-width: 48rem;
			background-color: #222222dd;
			margin: 0 auto;
		}

		header aside {
			max-width: 48rem;
			background-color: #222222dd;
			margin: 0 auto;
		}

		header aside figure {
			display: flex;
			align-items: flex-start;
			margin: 0;
		}

		header aside figure figcaption {
			margin: 0 1em 1em;
		}

		nav {
			max-width: 48rem;
			display: flex;
			margin-bottom: 1em;
			background-color: #333;
			margin: 0 auto;
		}

		nav a {
			flex: 1;
			padding: 0 0.5em;
		}

		nav a.prev {
			text-align:left;
		}

		nav a.next {
			text-align:right;
		}

		main {
			max-width: 48rem;
			background-color: #222222dd;
			margin: 0 auto;
		}

		main article {
			padding: 1em 0.5em;
		}

		main article > p {
			margin: 0;
		}

		main article ul {
			display: flex;
			list-style: none;
			flex-wrap: wrap;
			padding: 0;
			margin: 0;
		}

		main article ul li {
			text-align: center;
			flex: 1;
			padding: 1em;
			box-sizing: border-box;
			margin: 0;
		}

		main article img {
			min-width: 200px;
			display: block;
		}

		main article+article {
			border-top: 1px solid #333;
		}

		main article aside {
			text-align: right;
			color: #333;
		}

		main article aside details p {
			margin: 0;
		}

		main article aside details abbr {
			text-decoration: none;
		}

		main article aside details time a {
			color: #333 !important;
			text-decoration: none;
		}

		main article aside details time a:hover {
			text-decoration: underline;
		}

		footer {
			max-width: 48rem;
			text-align: center;
			margin: 0 auto;
		}
	</style>
	</head>
	<body>
	<header>
		<h1>
			@{{ .UserInfo.UserName }}
		</h1>
		<aside>
			<figure>
				{{ $profile_avatar_url := profile_avatar_url .Profile }}
				{{ with $profile_avatar_url }}
				<img src="{{ . }}" alt="{{ $.UserInfo.UserName }} Avatar">
				{{ end }}

				<figcaption>
					<details>
						<summary>Bio</summary>
						<strong>{{ .UserInfo.DisplayName }}</strong>
						<p>{{ .Profile.Description.Bio }}</p>
					</details>
					<details>
						<summary>Website</summary>
						<p>
							<a href="{{ .Profile.Description.Website }}">
								{{ .Profile.Description.Website }}
							</a>
						</p>
					</details>
					<details>
						<summary>Location</summary>
						<p>{{ .Profile.Description.Location }}</p>
					</details>
				</figcaption>
			</figure>
		</aside>
	</header>
	{{ if or (.PrevPage) (.NextPage) }}
	<nav>
	{{ if .PrevPage }}
	<a class="prev" href="{{ .PrevPage }}">Previous</a>
	{{ end }}
	{{ if .NextPage }}
	<a class="next" href="{{ .NextPage }}">Next</a>
	{{ end }}
	</nav>
	{{ end }}
	<main>
`

var pageIndexTweetTmpl = `
	<article class="tweet">
		{{ fancy_tweet . }}
		<aside>
			<details>
				<summary>meta</summary>
				<p>
					<abbr title="retweets">♻</abbr>  {{ .RetweetCount }} |
					<abbr title="likes">♥</abbr> {{ .FavoriteCount }} |
					<abbr title="posted at">⏲</abbr>
					<time datetime="{{ .CreatedAt.Format "2006-01-02T15:04:05Z07:00" }}">
						<a href="{{ tweet_url . }}">
							{{ .CreatedAt.Format "Jan 02, 2006 15:04:05" }}
						</a>
					</time>
				</p>
			</details>
		</aside>
	</article>
`

var pageIndexFooterTmpl = `
	</main>
	{{ if or (.PrevPage) (.NextPage) }}
	<nav>
	{{ if .PrevPage}}
	<a class="prev" href="{{ .PrevPage }}">Previous</a>
	{{ end }}
	{{ if .NextPage }}
	<a class="next" href="{{ .NextPage }}">Next</a>
	{{ end }}
	</nav>
	{{ end }}
	<footer>
		<p>rendered with 🦉<a href="https://github.com/darkliquid/twoo">twoo</a></p>
	</footer>
	</body>
	</html>
`

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

func renderTweets(data *twitwoo.Data, pd PageData, filter filterFunc, funcMap template.FuncMap, w io.Writer) error {
	tweet := template.Must(template.New("tweet").Funcs(funcMap).Parse(pageIndexTweetTmpl))

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

	funcMap := FuncMap(m)

	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	pd.UserInfo = m.UserInfo
	pd.Profile = profiles[0]

	header := template.Must(template.New("header").Funcs(funcMap).Parse(pageIndexHeaderTmpl))
	if err = header.Execute(w, pd); err != nil {
		return err
	}

	if err = renderTweets(data, pd, fn, funcMap, w); err != nil {
		return err
	}

	footer := template.Must(template.New("footer").Funcs(funcMap).Parse(pageIndexFooterTmpl))
	return footer.Execute(w, pd)
}

// Content renders a page containing custom content.
func Content(data *twitwoo.Data, pd PageData, fn contentFunc, w io.Writer) error {
	return wrapper(data, pd, fn, w)
}

// Tweet renders a singular tweet without wrapper HTML for the rest of a page.
func Tweet(data *twitwoo.Data, tw *twitwoo.Tweet, w io.Writer) error {
	m, err := data.Manifest()
	if err != nil {
		return err
	}

	funcMap := FuncMap(m)

	return template.Must(template.New("tweet").Funcs(funcMap).Parse(pageIndexTweetTmpl)).Execute(w, tw)
}

type contentFunc func(data *twitwoo.Data, pd PageData, w io.Writer) error

func wrapper(data *twitwoo.Data, pd PageData, fn contentFunc, w io.Writer) error {
	m, err := data.Manifest()
	if err != nil {
		return err
	}

	funcMap := FuncMap(m)

	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	pd.UserInfo = m.UserInfo
	pd.Profile = profiles[0]

	header := template.Must(template.New("header").Funcs(funcMap).Parse(pageIndexHeaderTmpl))
	if err = header.Execute(w, pd); err != nil {
		return err
	}

	if err = fn(data, pd, w); err != nil {
		return err
	}

	footer := template.Must(template.New("footer").Funcs(funcMap).Parse(pageIndexFooterTmpl))
	return footer.Execute(w, pd)
}
