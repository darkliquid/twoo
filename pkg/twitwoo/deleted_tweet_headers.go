package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDeletedTweetHeadersDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DeletedTweetHeaders",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DeletedTweetHeaders is the structure of the data/deleted-tweet-headers.js file.
type DeletedTweetHeaders struct {
    // Fields go here
}

func (d *DeletedTweetHeaders) decode(el jsoniter.Any) {
	el = el.Get("deletedTweetHeaders")
	el.ToVal(d)
}

// DeletedTweetHeaderss returns all the DeletedTweetHeaders items.
func (d *Data) DeletedTweetHeaderss() ([]*DeletedTweetHeaders, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DeletedTweetHeaders](d, m.DataTypes.DeletedTweetHeaders)
}

// EachDeletedTweetHeaders calls fn for each DeletedTweetHeaders item.
func (d *Data) EachDeletedTweetHeaders(fn func(*DeletedTweetHeaders) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DeletedTweetHeaders](d, m.DataTypes.DeletedTweetHeaders, fn)
}
