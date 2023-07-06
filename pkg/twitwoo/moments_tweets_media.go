package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerMomentsTweetsMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.MomentsTweetsMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// MomentsTweetsMedia is the structure of the data/moments-tweets-media.js file.
type MomentsTweetsMedia struct {
    // Fields go here
}

func (m *MomentsTweetsMedia) decode(el jsoniter.Any) {
	el = el.Get("momentsTweetsMedia")
	el.ToVal(m)
}

// MomentsTweetsMedias returns all the MomentsTweetsMedia items.
func (d *Data) MomentsTweetsMedias() ([]*MomentsTweetsMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*MomentsTweetsMedia](d, m.DataTypes.MomentsTweetsMedia)
}

// EachMomentsTweetsMedia calls fn for each MomentsTweetsMedia item.
func (d *Data) EachMomentsTweetsMedia(fn func(*MomentsTweetsMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*MomentsTweetsMedia](d, m.DataTypes.MomentsTweetsMedia, fn)
}
