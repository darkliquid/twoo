package website

import (
	"errors"
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
		header {
			margin-bottom: 1em;
		}

		nav a:first-child {
			text-align:left;
		}

		nav a:last-child {
			text-align:right;
		}

		nav {
			display: flex;
			margin-bottom: 1em;
			background-color: #333;
			padding: 0 0.5em;
		}

		nav a {
			flex: 1;
		}

		main article {
			padding: 1em 0.5em;
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
			text-align: center;
		}
	</style>
	</head>
	<body>
	<header id="profile-header">
		<h1>{{ .UserInfo.DisplayName}}</h1>
		<h2>@{{ .UserInfo.UserName }}</h2>
		<aside>
			<details>
				<summary>Bio</summary>
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
		</aside>
	</header>
	{{ if gt .PageCount 0 }}
	<nav class="pagination">
	{{ if gt .PrevPage 0 }}
	<a href="/page{{ .PrevPage }}">Previous</a>
	{{ end }}
	{{ if lt .NextPage .PageCount }}
	<a href="/page/{{ .NextPage }}">Next</a>
	{{ end }}
	</nav>
	{{ end }}
	<main>
`

var pageIndexTweetTmpl = `
	<article class="tweet">
		<p>{{ fancy_tweet . }}</p>
		<aside>
			<details>
				<summary>meta</summary>
				<p>
					<abbr title="retweets">♻</abbr>  {{ .RetweetCount }} |
					<abbr title="likes">♥</abbr> {{ .FavoriteCount }} |
					<abbr title="posted at">⏲</abbr> <time datetime="{{ .CreatedAt.Format "2006-01-02T15:04:05Z07:00" }}"><a href="/tweet/{{ .ID }}">{{ .CreatedAt.Format "Jan 02, 2006 15:04:05" }}</a></time>
				</p>
			</details>
		</aside>
	</article>
`

var pageIndexFooterTmpl = `
	</main>
	{{ if gt .PageCount 0 }}
	<nav class="pagination">
	{{ if gt .PrevPage 0 }}
	<a href="/page/{{ .PrevPage }}">Previous</a>
	{{ end }}
	{{ if lt .NextPage .PageCount }}
	<a href="/page/{{ .NextPage }}">Next</a>
	{{ end }}
	</nav>
	{{ end }}
	<footer>
		<p>rendered with 🦉<a href="https://github.com/darkliquid/twoo">twoo</a></p>
	</footer>
	</body>
	</html>
`

type pageData struct {
	Profile    *twitwoo.Profile
	UserInfo   twitwoo.UserInfo
	Page       int64
	PageSize   int64
	PrevPage   int64
	NextPage   int64
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

	i := int64(0)
	return render(data, pageData{
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TweetCount: totalTweets,
		PrevPage:   page - 1,
		NextPage:   page + 1,
	}, func(t *twitwoo.Tweet) *twitwoo.Tweet {
		i++
		if i >= (page-1)*pageSize && i < page*pageSize {
			return t
		}
		return nil
	}, w)
}

// Page writes a single item page.
func Page(data *twitwoo.Data, id int64, w io.Writer) error {
	return render(data, pageData{}, func(t *twitwoo.Tweet) *twitwoo.Tweet {
		if t.ID == id {
			return t
		}
		return nil
	}, w)
}

func render(data *twitwoo.Data, pd pageData, fn func(*twitwoo.Tweet) *twitwoo.Tweet, w io.Writer) error {
	m, err := data.Manifest()
	if err != nil {
		return err
	}

	header := template.Must(template.New("header").Funcs(funcMap).Parse(pageIndexHeaderTmpl))

	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	pd.UserInfo = m.UserInfo
	pd.Profile = profiles[0]

	if err := header.Execute(w, pd); err != nil {
		return err
	}

	tweet := template.Must(template.New("tweet").Funcs(funcMap).Parse(pageIndexTweetTmpl))

	// TODO:This is an incredibly naive implementation, but for live serving
	// is probably fine. Maybe find some way to index deeper into the json data
	// for tweets to skip having to read and parse everything before the page you
	// actually want to display.
	if err := data.EachTweet(func(t *twitwoo.Tweet) error {
		t = fn(t)
		if t != nil {
			return tweet.Execute(w, t)
		}

		return nil
	}); err != nil {
		return err
	}

	footer := template.Must(template.New("footer").Funcs(funcMap).Parse(pageIndexFooterTmpl))
	return footer.Execute(w, pageData{
		UserInfo: m.UserInfo,
		Profile:  profiles[0],
	})
}
