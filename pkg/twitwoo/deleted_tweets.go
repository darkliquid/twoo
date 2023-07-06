package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDeletedTweetsDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DeletedTweets",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DeletedTweets is the structure of the data/deleted-tweets.js file.
type DeletedTweets struct {
    // Fields go here
}

func (d *DeletedTweets) decode(el jsoniter.Any) {
	el = el.Get("deletedTweets")
	el.ToVal(d)
}

// DeletedTweetss returns all the DeletedTweets items.
func (d *Data) DeletedTweetss() ([]*DeletedTweets, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DeletedTweets](d, m.DataTypes.DeletedTweets)
}

// EachDeletedTweets calls fn for each DeletedTweets item.
func (d *Data) EachDeletedTweets(fn func(*DeletedTweets) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DeletedTweets](d, m.DataTypes.DeletedTweets, fn)
}
