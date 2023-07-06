package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessageMuteDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessageMute",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessageMute is the structure of the data/direct-message-mute.js file.
type DirectMessageMute struct {
    // Fields go here
}

func (d *DirectMessageMute) decode(el jsoniter.Any) {
	el = el.Get("directMessageMute")
	el.ToVal(d)
}

// DirectMessageMutes returns all the DirectMessageMute items.
func (d *Data) DirectMessageMutes() ([]*DirectMessageMute, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessageMute](d, m.DataTypes.DirectMessageMute)
}

// EachDirectMessageMute calls fn for each DirectMessageMute item.
func (d *Data) EachDirectMessageMute(fn func(*DirectMessageMute) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessageMute](d, m.DataTypes.DirectMessageMute, fn)
}
