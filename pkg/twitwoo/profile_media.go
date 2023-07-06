package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerProfileMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ProfileMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ProfileMedia is the structure of the data/profile-media.js file.
type ProfileMedia struct {
    // Fields go here
}

func (p *ProfileMedia) decode(el jsoniter.Any) {
	el = el.Get("profileMedia")
	el.ToVal(p)
}

// ProfileMedias returns all the ProfileMedia items.
func (d *Data) ProfileMedias() ([]*ProfileMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ProfileMedia](d, m.DataTypes.ProfileMedia)
}

// EachProfileMedia calls fn for each ProfileMedia item.
func (d *Data) EachProfileMedia(fn func(*ProfileMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ProfileMedia](d, m.DataTypes.ProfileMedia, fn)
}
