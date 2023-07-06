package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterCircleTweetMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterCircleTweetMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterCircleTweetMedia is the structure of the data/twitter-circle-tweet-media.js file.
type TwitterCircleTweetMedia struct {
    // Fields go here
}

func (t *TwitterCircleTweetMedia) decode(el jsoniter.Any) {
	el = el.Get("twitterCircleTweetMedia")
	el.ToVal(t)
}

// TwitterCircleTweetMedias returns all the TwitterCircleTweetMedia items.
func (d *Data) TwitterCircleTweetMedias() ([]*TwitterCircleTweetMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterCircleTweetMedia](d, m.DataTypes.TwitterCircleTweetMedia)
}

// EachTwitterCircleTweetMedia calls fn for each TwitterCircleTweetMedia item.
func (d *Data) EachTwitterCircleTweetMedia(fn func(*TwitterCircleTweetMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterCircleTweetMedia](d, m.DataTypes.TwitterCircleTweetMedia, fn)
}
