package website

import (
	"fmt"
	"html/template"
	"path"
	"regexp"
	"strings"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var urlRE = regexp.MustCompile(`(?i)[^>"](http|https|ftp):\/\/(\S*)[^<"]`)

const linkSubstitution = "<a href=\"$1://$2\">$1://$2</a>"

func FuncMap(m *twitwoo.Manifest) template.FuncMap {
	return template.FuncMap{
		"fancy_tweet": func(t *twitwoo.Tweet) template.HTML {
			text := t.FullText
			text = strings.ReplaceAll(text, "\n", "<br>")

			for _, tag := range t.Hashtags {
				text = strings.ReplaceAll(text, "#"+tag, "<a href=\"https://twitter.com/hashtag/"+tag+"\">#"+tag+"</a>")
			}

			for url, link := range t.URLMap {
				text = strings.ReplaceAll(text, url, "<a href=\""+link.ExpandedURL+"\">"+link.DisplayURL+"</a>")
			}

			for _, mention := range t.Mentions {
				text = strings.ReplaceAll(text, "@"+mention.ScreenName, "<a href=\"https://twitter.com/"+mention.ScreenName+"\">@"+mention.ScreenName+"</a>")
			}

			text = urlRE.ReplaceAllString(text, linkSubstitution)

			return template.HTML(text)
		},
		"profile_header_url": func(p *twitwoo.Profile) string {
			if p.Header == "" {
				return ""
			}
			return fmt.Sprintf("/data/profile_media/%d-%s.jpg", m.UserInfo.AccountID, path.Base(p.Header))
		},
		"profile_avatar_url": func(p *twitwoo.Profile) string {
			if p.Avatar == "" {
				return ""
			}
			return fmt.Sprintf("/data/profile_media/%d-%s", m.UserInfo.AccountID, path.Base(p.Avatar))
		},

	}
}
