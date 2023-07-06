package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerMuteDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Mute",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Mute is the structure of the data/mute.js file.
type Mute struct {
    // Fields go here
}

func (m *Mute) decode(el jsoniter.Any) {
	el = el.Get("mute")
	el.ToVal(m)
}

// Mutes returns all the Mute items.
func (d *Data) Mutes() ([]*Mute, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Mute](d, m.DataTypes.Mute)
}

// EachMute calls fn for each Mute item.
func (d *Data) EachMute(fn func(*Mute) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Mute](d, m.DataTypes.Mute, fn)
}
