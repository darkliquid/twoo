package website

import (
	"fmt"
	"html/template"
	"path"
	"regexp"
	"strings"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var urlRE = regexp.MustCompile(`(?i)(^|[^>"])(http|https|ftp):\/\/(\S+)([^<"]|$)`)

const (
	linkSubstitution = " <a href=\"$2://$3\">$2://$3</a> "
)

func hashtagLink(tag string) string {
	return "<a href=\"https://twitter.com/hashtag/" + tag + "\">#" + tag + "</a>"
}

func expandedLink(link twitwoo.Link) string {
	return "<a href=\"" + link.ExpandedURL + "\">" + link.DisplayURL + "</a>"
}

func mentionLink(mention twitwoo.Mention) string {
	return "<a href=\"https://twitter.com/" + mention.ScreenName + "\">@" + mention.ScreenName + "</a>"
}

func fancyTweetMedia(t *twitwoo.Tweet) string {
	includedMedia := 0
	mediaList := ""
	if len(t.Media) > 0 {
		mediaList = "<ul>"
		for _, media := range t.Media {
			if media.SourceStatusID > 0 && media.SourceStatusID != t.ID {
				continue
			}
			mediaList += "<li>"
			file := TweetMediaPath(t.ID, media)
			switch media.Type {
			case "photo":
				mediaList += fmt.Sprintf(`<img src="%s">`, file)
				includedMedia++
			case "video", "animated_gif":
				includedMedia++
				mediaList += fmt.Sprintf(`<video controls><source src="%s" type="video/mp4"></video>`, file)
			}
			mediaList += "</li>"
		}
		mediaList += "</ul>"
	}

	if includedMedia > 0 {
		return mediaList
	}

	return ""
}

func fancyTweet(t *twitwoo.Tweet) template.HTML {
	text := "<p>" + t.FullText
	text = strings.ReplaceAll(text, "\n", "<br>")

	for _, tag := range t.Hashtags {
		text = strings.ReplaceAll(text, "#"+tag, hashtagLink(tag))
	}

	for url, link := range t.URLMap {
		text = strings.ReplaceAll(text, url, expandedLink(link))
	}

	for _, mention := range t.Mentions {
		text = strings.ReplaceAll(text, "@"+mention.ScreenName, mentionLink(mention))
	}

	text = urlRE.ReplaceAllString(text, linkSubstitution)

	text += "</p>"

	text += fancyTweetMedia(t)

	return template.HTML(text) //nolint:gosec // input is trusted
}

// FuncMap returns a template.FuncMap for use in the website templates.
func FuncMap(m *twitwoo.Manifest) template.FuncMap {
	return template.FuncMap{
		"fancy_tweet": fancyTweet,
		"profile_header_url": func(p *twitwoo.Profile) string {
			if p.Header == "" {
				return ""
			}
			return ProfileMediaPath(m.UserInfo.AccountID, p.Header) + ".jpg"
		},
		"profile_avatar_url": func(p *twitwoo.Profile) string {
			if p.Avatar == "" {
				return ""
			}
			return ProfileMediaPath(m.UserInfo.AccountID, p.Avatar)
		},
		"tweet_url": func(t *twitwoo.Tweet) string {
			y, m, d := t.CreatedAt.Date()
			return fmt.Sprintf("/%d/%02d/%02d/%020d", y, m, d, t.ID)
		},
		"search_js": func() template.HTML {
			if !EnableSearch {
				return ""
			}

			return template.HTML(`
  <script type="module">
    import { search, default as init } from '/tinysearch_engine.js';
    window.search = search;

    async function run() {
      await init('/tinysearch_engine_bg.wasm');
    }

    run();
  </script>

  <script>
    function doSearch() {
      let value = document.getElementById("search").value;
      const results = search(value, 5);
      let ul = document.getElementById("search-results");
      ul.innerHTML = "";
	  ul.style.display = "none";

      for (i = 0; i < results.length; i++) {
        var li = document.createElement("li");

        let [title, url] = results[i];
        let elemlink = document.createElement('a');
        elemlink.innerHTML = title;
        elemlink.setAttribute('href', url);
        li.appendChild(elemlink);

        ul.appendChild(li);
      }
	  if (results.length > 0) {
		ul.style.display = "block";
	  }
    }
  </script>
			`) //nolint:gosec // input is trusted
		},
		"searchbox": func() template.HTML {
			if !EnableSearch {
				return ""
			}

			return template.HTML(`
<input type="text" id="search" onkeyup="doSearch()" placeholder="Search...">
`)
		},
		"search_results": func() template.HTML {
			if !EnableSearch {
				return ""
			}

			return template.HTML(`<ul id="search-results"></ul>`)
		},
	}
}

// TweetMediaPath returns the path to the media file for the given tweet ID and media URL.
func TweetMediaPath(tweetID int64, media twitwoo.Media) string {
	return fmt.Sprintf("/data/tweets_media/%d-%s", tweetID, path.Base(media.MediaURL))
}

// ProfileMediaPath returns the path to the media for the given profile ID and media URL.
func ProfileMediaPath(accountID int64, mediaURL string) string {
	return fmt.Sprintf("/data/profile_media/%d-%s", accountID, path.Base(mediaURL))
}
