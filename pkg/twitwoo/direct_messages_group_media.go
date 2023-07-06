package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessagesGroupMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessagesGroupMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessagesGroupMedia is the structure of the data/direct-messages-group-media.js file.
type DirectMessagesGroupMedia struct {
    // Fields go here
}

func (d *DirectMessagesGroupMedia) decode(el jsoniter.Any) {
	el = el.Get("directMessagesGroupMedia")
	el.ToVal(d)
}

// DirectMessagesGroupMedias returns all the DirectMessagesGroupMedia items.
func (d *Data) DirectMessagesGroupMedias() ([]*DirectMessagesGroupMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessagesGroupMedia](d, m.DataTypes.DirectMessagesGroupMedia)
}

// EachDirectMessagesGroupMedia calls fn for each DirectMessagesGroupMedia item.
func (d *Data) EachDirectMessagesGroupMedia(fn func(*DirectMessagesGroupMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessagesGroupMedia](d, m.DataTypes.DirectMessagesGroupMedia, fn)
}
