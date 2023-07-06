package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDeletedNoteTweetDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DeletedNoteTweet",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DeletedNoteTweet is the structure of the data/deleted-note-tweet.js file.
type DeletedNoteTweet struct {
    // Fields go here
}

func (d *DeletedNoteTweet) decode(el jsoniter.Any) {
	el = el.Get("deletedNoteTweet")
	el.ToVal(d)
}

// DeletedNoteTweets returns all the DeletedNoteTweet items.
func (d *Data) DeletedNoteTweets() ([]*DeletedNoteTweet, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DeletedNoteTweet](d, m.DataTypes.DeletedNoteTweet)
}

// EachDeletedNoteTweet calls fn for each DeletedNoteTweet item.
func (d *Data) EachDeletedNoteTweet(fn func(*DeletedNoteTweet) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DeletedNoteTweet](d, m.DataTypes.DeletedNoteTweet, fn)
}
