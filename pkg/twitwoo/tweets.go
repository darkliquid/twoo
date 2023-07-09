package twitwoo

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

func registerTweetDecoders() {
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"ID",
		stringToInt64("decode id"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"FavoriteCount",
		stringToInt64("decode favourite count"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"RetweetCount",
		stringToInt64("decode retweet count"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Tweet",
		"CreatedAt",
		stringToTime("decode created at", time.RubyDate),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Mention",
		"ID",
		stringToInt64("decode id"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Media",
		"ID",
		stringToInt64("decode id"),
	)
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
	ExpandedURL string `json:"expanded_url"`
	URL         string `json:"url"`
	MediaURL    string `json:"media_url"`
	Type        string `json:"type"`
	DisplayURL  string `json:"display_url"`
	ID          int64  `json:"id"`
}

// Tweet represents a single tweet.
type Tweet struct {
	CreatedAt         time.Time       `json:"created_at"`
	URLMap            map[string]Link `json:"urlmap"`
	InReplyToUserID   string          `json:"in_reply_to_user_id_str"`
	InReplyToStatusID string          `json:"in_reply_to_status_id_str"`
	FullText          string          `json:"full_text"`
	Hashtags          []string        `json:"hashtags"`
	Mentions          []Mention       `json:"mentions"`
	Media             []Media         `json:"media"`
	ID                int64           `json:"id"`
	RetweetCount      int64           `json:"retweet_count"`
	FavoriteCount     int64           `json:"favorite_count"`
}

func (t *Tweet) decode(el jsoniter.Any) {
	el = el.Get("tweet")
	el.ToVal(t)

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
	el.Get("entities", "hashtags").ToVal(&hashtags)

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
	el.Get("entities", "urls").ToVal(&urls)

	// extract urls into a map of url -> expanded url.
	if len(urls) > 0 {
		t.URLMap = make(map[string]Link, len(urls))
	}
	for _, v := range urls {
		t.URLMap[v.URL] = v.Link
	}
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
