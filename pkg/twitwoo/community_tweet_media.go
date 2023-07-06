package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCommunityTweetMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CommunityTweetMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CommunityTweetMedia is the structure of the data/community-tweet-media.js file.
type CommunityTweetMedia struct {
    // Fields go here
}

func (c *CommunityTweetMedia) decode(el jsoniter.Any) {
	el = el.Get("communityTweetMedia")
	el.ToVal(c)
}

// CommunityTweetMedias returns all the CommunityTweetMedia items.
func (d *Data) CommunityTweetMedias() ([]*CommunityTweetMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CommunityTweetMedia](d, m.DataTypes.CommunityTweetMedia)
}

// EachCommunityTweetMedia calls fn for each CommunityTweetMedia item.
func (d *Data) EachCommunityTweetMedia(fn func(*CommunityTweetMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CommunityTweetMedia](d, m.DataTypes.CommunityTweetMedia, fn)
}
