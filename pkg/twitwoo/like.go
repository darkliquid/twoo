package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerLikeDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Like",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Like is the structure of the data/like.js file.
type Like struct {
    // Fields go here
}

func (l *Like) decode(el jsoniter.Any) {
	el = el.Get("like")
	el.ToVal(l)
}

// Likes returns all the Like items.
func (d *Data) Likes() ([]*Like, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Like](d, m.DataTypes.Like)
}

// EachLike calls fn for each Like item.
func (d *Data) EachLike(fn func(*Like) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Like](d, m.DataTypes.Like, fn)
}
