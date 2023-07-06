package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCommunityTweetDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CommunityTweet",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CommunityTweet is the structure of the data/community-tweet.js file.
type CommunityTweet struct {
    // Fields go here
}

func (c *CommunityTweet) decode(el jsoniter.Any) {
	el = el.Get("communityTweet")
	el.ToVal(c)
}

// CommunityTweets returns all the CommunityTweet items.
func (d *Data) CommunityTweets() ([]*CommunityTweet, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CommunityTweet](d, m.DataTypes.CommunityTweet)
}

// EachCommunityTweet calls fn for each CommunityTweet item.
func (d *Data) EachCommunityTweet(fn func(*CommunityTweet) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CommunityTweet](d, m.DataTypes.CommunityTweet, fn)
}
