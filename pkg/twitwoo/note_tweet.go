package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerNoteTweetDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.NoteTweet",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// NoteTweet is the structure of the data/note-tweet.js file.
type NoteTweet struct {
    // Fields go here
}

func (n *NoteTweet) decode(el jsoniter.Any) {
	el = el.Get("noteTweet")
	el.ToVal(n)
}

// NoteTweets returns all the NoteTweet items.
func (d *Data) NoteTweets() ([]*NoteTweet, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*NoteTweet](d, m.DataTypes.NoteTweet)
}

// EachNoteTweet calls fn for each NoteTweet item.
func (d *Data) EachNoteTweet(fn func(*NoteTweet) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*NoteTweet](d, m.DataTypes.NoteTweet, fn)
}
