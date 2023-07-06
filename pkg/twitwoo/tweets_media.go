package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTweetsMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TweetsMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TweetsMedia is the structure of the data/tweets-media.js file.
type TweetsMedia struct {
    // Fields go here
}

func (t *TweetsMedia) decode(el jsoniter.Any) {
	el = el.Get("tweetsMedia")
	el.ToVal(t)
}

// TweetsMedias returns all the TweetsMedia items.
func (d *Data) TweetsMedias() ([]*TweetsMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TweetsMedia](d, m.DataTypes.TweetsMedia)
}

// EachTweetsMedia calls fn for each TweetsMedia item.
func (d *Data) EachTweetsMedia(fn func(*TweetsMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TweetsMedia](d, m.DataTypes.TweetsMedia, fn)
}
