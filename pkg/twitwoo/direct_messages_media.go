package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessagesMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessagesMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessagesMedia is the structure of the data/direct-messages-media.js file.
type DirectMessagesMedia struct {
    // Fields go here
}

func (d *DirectMessagesMedia) decode(el jsoniter.Any) {
	el = el.Get("directMessagesMedia")
	el.ToVal(d)
}

// DirectMessagesMedias returns all the DirectMessagesMedia items.
func (d *Data) DirectMessagesMedias() ([]*DirectMessagesMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessagesMedia](d, m.DataTypes.DirectMessagesMedia)
}

// EachDirectMessagesMedia calls fn for each DirectMessagesMedia item.
func (d *Data) EachDirectMessagesMedia(fn func(*DirectMessagesMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessagesMedia](d, m.DataTypes.DirectMessagesMedia, fn)
}
