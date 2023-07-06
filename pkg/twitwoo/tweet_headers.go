package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTweetHeadersDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TweetHeaders",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TweetHeaders is the structure of the data/tweet-headers.js file.
type TweetHeaders struct {
    // Fields go here
}

func (t *TweetHeaders) decode(el jsoniter.Any) {
	el = el.Get("tweetHeaders")
	el.ToVal(t)
}

// TweetHeaderss returns all the TweetHeaders items.
func (d *Data) TweetHeaderss() ([]*TweetHeaders, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TweetHeaders](d, m.DataTypes.TweetHeaders)
}

// EachTweetHeaders calls fn for each TweetHeaders item.
func (d *Data) EachTweetHeaders(fn func(*TweetHeaders) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TweetHeaders](d, m.DataTypes.TweetHeaders, fn)
}
