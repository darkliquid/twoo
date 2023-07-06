package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerMomentsMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.MomentsMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// MomentsMedia is the structure of the data/moments-media.js file.
type MomentsMedia struct {
    // Fields go here
}

func (m *MomentsMedia) decode(el jsoniter.Any) {
	el = el.Get("momentsMedia")
	el.ToVal(m)
}

// MomentsMedias returns all the MomentsMedia items.
func (d *Data) MomentsMedias() ([]*MomentsMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*MomentsMedia](d, m.DataTypes.MomentsMedia)
}

// EachMomentsMedia calls fn for each MomentsMedia item.
func (d *Data) EachMomentsMedia(fn func(*MomentsMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*MomentsMedia](d, m.DataTypes.MomentsMedia, fn)
}
