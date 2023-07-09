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

func FuncMap(m *twitwoo.Manifest) template.FuncMap {
	return template.FuncMap{
		"fancy_tweet": func(t *twitwoo.Tweet) template.HTML {
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

			if len(t.Media) > 0 {
				text += "<ul>"
				for _, media := range t.Media {
					text += "<li>"
					file := fmt.Sprintf("/data/tweets_media/%d-%s", t.ID, path.Base(media.MediaURL))
					switch media.Type {
					case "photo":
						text += fmt.Sprintf(`<img src="%s">`, file)
					case "video":
						text += fmt.Sprintf(`<video controls><source src="%s" type="video/mp4"></video>`, file)
					}
					text += "</li>"
				}
				text += "</ul>"
			}

			return template.HTML(text) //nolint:gosec // input is trusted
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
