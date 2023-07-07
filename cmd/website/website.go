package website

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var pageIndexHeaderTmpl = `<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>@{{ .UserInfo.UserName }} - {{ .Page }}/{{ .PageCount }}</title>
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
	<nav class="pagination">
	{{ if gt .PrevPage 0 }}
	<a href="?page={{ .PrevPage }}&page_size={{ .PageSize }}">Previous</a>
	{{ end }}
	{{ if lt .NextPage .PageCount }}
	<a href="?page={{ .NextPage }}&page_size={{ .PageSize }}">Next</a>
	{{ end }}
	</nav>
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
					<abbr title="posted at">⏲</abbr> <time datetime="{{ .CreatedAt.Format "2006-01-02T15:04:05Z07:00" }}">{{ .CreatedAt.Format "Jan 02, 2006 15:04:05" }}</time>
				</p>
			</details>
		</aside>
	</article>
`

var pageIndexFooterTmpl = `
	</main>
	<nav class="pagination">
	{{ if gt .PrevPage 0 }}
	<a href="?page={{ .PrevPage }}&page_size={{ .PageSize }}">Previous</a>
	{{ end }}
	{{ if lt .NextPage .PageCount }}
	<a href="?page={{ .NextPage }}&page_size={{ .PageSize }}">Next</a>
	{{ end }}
	</nav>
	<footer>
		<p>rendered with 🦉<a href="https://github.com/darkliquid/twoo">twoo</a></p>
	</footer>
	</body>
	</html>
`

// Page returns a page of data.
func Page(data *twitwoo.Data, page, pageSize int64, w http.ResponseWriter) error {
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

	header := template.Must(template.New("header").Funcs(funcMap).Parse(pageIndexHeaderTmpl))

	profiles, err := data.Profiles()
	if err != nil {
		return err
	}
	if len(profiles) < 1 {
		return errors.New("no profiles found")
	}

	if err := header.Execute(w, struct {
		Profile    *twitwoo.Profile
		UserInfo   twitwoo.UserInfo
		Page       int64
		PageSize   int64
		PrevPage   int64
		NextPage   int64
		PageCount  int64
		TweetCount int64
	}{
		UserInfo:   m.UserInfo,
		Profile:    profiles[0],
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TweetCount: totalTweets,
		PrevPage:   page - 1,
		NextPage:   page + 1,
	}); err != nil {
		return err
	}

	tweet := template.Must(template.New("tweet").Funcs(funcMap).Parse(pageIndexTweetTmpl))

	// TODO:This is an incredibly naive implementation, but for live serving
	// is probably fine. Maybe find some way to index deeper into the json data
	// for tweets to skip having to read and parse everything before the page you
	// actually want to display.
	i := int64(0)
	if err := data.EachTweet(func(t *twitwoo.Tweet) error {
		i++
		if i >= (page-1)*pageSize && i < page*pageSize {
			return tweet.Execute(w, t)
		}

		return nil
	}); err != nil {
		return err
	}

	footer := template.Must(template.New("footer").Funcs(funcMap).Parse(pageIndexFooterTmpl))
	return footer.Execute(w, struct {
		Profile    *twitwoo.Profile
		UserInfo   twitwoo.UserInfo
		Page       int64
		PageSize   int64
		PrevPage   int64
		NextPage   int64
		PageCount  int64
		TweetCount int64
	}{
		UserInfo:   m.UserInfo,
		Profile:    profiles[0],
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TweetCount: totalTweets,
		PrevPage:   page - 1,
		NextPage:   page + 1,
	})
}
