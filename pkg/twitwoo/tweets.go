package twitwoo

import (
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

//nolint:gocognit // Big function, but not complex - straightforward JSON parsing.
func registerTweetDecoders() {
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Mention",
		"ID",
		stringToInt64("decode id"),
	)

	jsoniter.RegisterTypeDecoderFunc(
		"twitwoo.Tweet",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			t := ((*Tweet)(ptr))
			el := iter.ReadAny()

			el = el.Get("tweet")
			var err error
			t.ID, err = strconv.ParseInt(el.Get("id").ToString(), 10, 64)
			if err != nil {
				iter.ReportError("decode tweet id", err.Error())
			}

			t.CreatedAt, err = time.Parse(time.RubyDate, el.Get("created_at").ToString())
			if err != nil {
				iter.ReportError("decode tweet created_at", err.Error())
			}

			t.FullText = el.Get("full_text").ToString()
			t.RetweetCount, err = strconv.ParseInt(el.Get("retweet_count").ToString(), 10, 64)
			if err != nil {
				iter.ReportError("decode tweet retweet count", err.Error())
			}

			t.FavoriteCount, err = strconv.ParseInt(el.Get("favorite_count").ToString(), 10, 64)
			if err != nil {
				iter.ReportError("decode tweet favourite count", err.Error())
			}

			userReply := el.Get("in_reply_to_user_id")
			if userReply.ValueType() != jsoniter.InvalidValue {
				t.InReplyToUserID, err = strconv.ParseInt(userReply.ToString(), 10, 64)
				if err != nil {
					iter.ReportError("decode tweet in reply to user id", err.Error())
				}
			}

			statusReply := el.Get("in_reply_to_status_id")
			if statusReply.ValueType() != jsoniter.InvalidValue {
				t.InReplyToStatusID, err = strconv.ParseInt(statusReply.ToString(), 10, 64)
				if err != nil {
					iter.ReportError("decode tweet in reply to status id", err.Error())
				}
			}

			media := el.Get("extended_entities", "media")
			if media.Size() > 0 {
				media.ToVal(&t.Media)
			}

			mentions := el.Get("entities", "user_mentions")
			if mentions.Size() > 0 {
				mentions.ToVal(&t.Mentions)
			}

			var hashtags []struct {
				Text string `json:"text"`
			}
			hts := el.Get("entities", "hashtags")
			if hts.Size() > 0 {
				hts.ToVal(&hashtags)
			}

			// extract hashtags to something useful.
			if len(hashtags) > 0 {
				t.Hashtags = make([]string, 0, len(hashtags))
			}
			for _, v := range hashtags {
				t.Hashtags = append(t.Hashtags, v.Text)
			}

			var urls []struct {
				URL string `json:"url"`
				Link
			}
			links := el.Get("entities", "urls")
			if links.Size() > 0 {
				links.ToVal(&urls)
			}

			// extract urls into a map of url -> expanded url.
			if len(urls) > 0 {
				t.URLMap = make(map[string]Link, len(urls))
			}
			for _, v := range urls {
				t.URLMap[v.URL] = v.Link
			}
		},
	)
}

func registerTweetMediaDecoders() {
	jsoniter.RegisterFieldDecoderFunc("twitwoo.Variant", "Bitrate", stringToInt64("decode bitrate"))

	jsoniter.RegisterTypeDecoderFunc(
		"twitwoo.Media",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			var err error
			m := ((*Media)(ptr))

			el := iter.ReadAny()
			m.MediaURL = el.Get("media_url").ToString()
			m.ID, err = strconv.ParseInt(el.Get("id").ToString(), 10, 64)
			if err != nil {
				iter.ReportError("decode media id", err.Error())
				return
			}
			sourceStatusID := el.Get("source_status_id")
			if sourceStatusID.ValueType() != jsoniter.InvalidValue {
				m.SourceStatusID, err = strconv.ParseInt(el.Get("source_status_id").ToString(), 10, 64)
				if err != nil {
					iter.ReportError("decode media source status id", err.Error())
					return
				}
			}

			m.ExpandedURL = el.Get("expanded_url").ToString()
			m.URL = el.Get("url").ToString()
			m.Type = el.Get("type").ToString()
			m.DisplayURL = el.Get("display_url").ToString()

			variants := el.Get("video_info", "variants")
			if variants.Size() > 0 {
				v := make([]Variant, variants.Size())
				variants.ToVal(&v)

				// Only the highest bitrate is needed.
				sort.Slice(v, func(i, j int) bool {
					return v[i].Bitrate > v[j].Bitrate
				})

				m.MediaURL = v[0].URL
				m.MediaURL = strings.Split(m.MediaURL, "?")[0]
			}

			if err != nil {
				iter.ReportError("decode media url", err.Error())
				return
			}
		},
	)
}

// Variant represents a video variant.
type Variant struct {
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
	Bitrate     int64  `json:"bitrate"`
}

// Mention represents a mention of a user in a tweet.
type Mention struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	ID         int64  `json:"id"`
}

// Link represents a link in a tweet.
type Link struct {
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
}

// Media represents a media item in a tweet.
type Media struct {
	ExpandedURL    string `json:"expanded_url"`
	URL            string `json:"url"`
	MediaURL       string `json:"media_url"`
	Type           string `json:"type"`
	DisplayURL     string `json:"display_url"`
	ID             int64  `json:"id"`
	SourceStatusID int64  `json:"source_status_id"`
}

// Tweet represents a single tweet.
type Tweet struct {
	CreatedAt         time.Time       `json:"created_at"`
	URLMap            map[string]Link `json:"urlmap"`
	FullText          string          `json:"full_text"`
	Hashtags          []string        `json:"hashtags"`
	Mentions          []Mention       `json:"mentions"`
	Media             []Media         `json:"media"`
	InReplyToUserID   int64           `json:"in_reply_to_user_id"`
	InReplyToStatusID int64           `json:"in_reply_to_status_id"`
	ID                int64           `json:"id"`
	RetweetCount      int64           `json:"retweet_count"`
	FavoriteCount     int64           `json:"favorite_count"`
}

// Tweets returns a slice of tweets.
func (d *Data) Tweets() ([]*Tweet, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Tweet](d, m.DataTypes.Tweets)
}

// EachTweet calls fn for each tweet.
func (d *Data) EachTweet(fn func(*Tweet) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Tweet](d, m.DataTypes.Tweets, fn)
}
