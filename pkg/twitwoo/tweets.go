package twitwoo

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

func registerTweetDecoders() {
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
}

type entities struct {
	Hashtags []struct {
		Text string `json:"text"`
	} `json:"hashtags"`
	URLs []struct {
		URL         string `json:"url"`
		ExpandedURL string `json:"expanded_url"`
	} `json:"urls"`
}

// Tweet represents a single tweet.
type Tweet struct {
	CreatedAt          time.Time         `json:"created_at"`
	URLMap             map[string]string `json:"urlmap"`
	ID                 string            `json:"id_str"`
	InReplyToUserID    string            `json:"in_reply_to_user_id_str"`
	InReplayToStatusID string            `json:"in_reply_to_status_id_str"`
	FullText           string            `json:"full_text"`
	Hashtags           []string          `json:"hashtags"`
	RetweetCount       int64             `json:"retweet_count"`
	FavoriteCount      int64             `json:"favorite_count"`
}

func (t *Tweet) decode(el jsoniter.Any) {
	el = el.Get("tweet")
	el.ToVal(t)

	var e entities
	el.Get("entities").ToVal(&e)

	// extract hashtags to something useful
	if len(e.Hashtags) > 0 {
		t.Hashtags = make([]string, 0, len(e.Hashtags))
	}
	for _, v := range e.Hashtags {
		t.Hashtags = append(t.Hashtags, v.Text)
	}

	// extract urls into a map of url -> expanded url
	if len(e.URLs) > 0 {
		t.URLMap = make(map[string]string, len(e.URLs))
	}
	for _, v := range e.URLs {
		t.URLMap[v.URL] = v.ExpandedURL
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
