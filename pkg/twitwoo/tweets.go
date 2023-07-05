package twitwoo

import (
	"io"
	"time"

	jsoniter "github.com/json-iterator/go"
)

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

func (d *Data) readTweets() (io.ReadCloser, error) {
	m, err := d.manifest()
	if err != nil {
		return nil, err
	}

	files := make([]io.Reader, len(m.DataTypes.Tweets.Files))
	for i, df := range m.DataTypes.Tweets.Files {
		r, err := d.readDataFile(&df)
		if err != nil {
			return nil, err
		}
		files[i] = r
	}

	return NewMultiReadCloser(files...), nil
}

// Tweets returns a slice of tweets.
func (d *Data) Tweets() ([]Tweet, error) {
	r, err := d.readTweets()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	tweets := make([]Tweet, 0)
	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, 1024)
	for iter.ReadArray() {
		var tweet Tweet
		decode(iter.ReadAny().Get("tweet"), &tweet)
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

// EachTweet calls fn for each tweet.
func (d *Data) EachTweet(fn func(Tweet) error) error {
	r, err := d.readTweets()
	if err != nil {
		return err
	}
	defer r.Close()

	iter := jsoniter.Parse(jsoniter.ConfigFastest, r, 1024)
	for iter.ReadArray() {
		var tweet Tweet
		decode(iter.ReadAny().Get("tweet"), &tweet)
		if err := fn(tweet); err != nil {
			return err
		}
	}

	return nil
}
