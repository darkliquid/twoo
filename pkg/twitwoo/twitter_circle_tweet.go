package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterCircleTweetDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterCircleTweet",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterCircleTweet is the structure of the data/twitter-circle-tweet.js file.
type TwitterCircleTweet struct {
    // Fields go here
}

func (t *TwitterCircleTweet) decode(el jsoniter.Any) {
	el = el.Get("twitterCircleTweet")
	el.ToVal(t)
}

// TwitterCircleTweets returns all the TwitterCircleTweet items.
func (d *Data) TwitterCircleTweets() ([]*TwitterCircleTweet, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterCircleTweet](d, m.DataTypes.TwitterCircleTweet)
}

// EachTwitterCircleTweet calls fn for each TwitterCircleTweet item.
func (d *Data) EachTwitterCircleTweet(fn func(*TwitterCircleTweet) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterCircleTweet](d, m.DataTypes.TwitterCircleTweet, fn)
}
