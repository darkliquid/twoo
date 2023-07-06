package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDeletedTweetsMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DeletedTweetsMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DeletedTweetsMedia is the structure of the data/deleted-tweets-media.js file.
type DeletedTweetsMedia struct {
    // Fields go here
}

func (d *DeletedTweetsMedia) decode(el jsoniter.Any) {
	el = el.Get("deletedTweetsMedia")
	el.ToVal(d)
}

// DeletedTweetsMedias returns all the DeletedTweetsMedia items.
func (d *Data) DeletedTweetsMedias() ([]*DeletedTweetsMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DeletedTweetsMedia](d, m.DataTypes.DeletedTweetsMedia)
}

// EachDeletedTweetsMedia calls fn for each DeletedTweetsMedia item.
func (d *Data) EachDeletedTweetsMedia(fn func(*DeletedTweetsMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DeletedTweetsMedia](d, m.DataTypes.DeletedTweetsMedia, fn)
}
