package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessagesDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessages",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessages is the structure of the data/direct-messages.js file.
type DirectMessages struct {
    // Fields go here
}

func (d *DirectMessages) decode(el jsoniter.Any) {
	el = el.Get("directMessages")
	el.ToVal(d)
}

// DirectMessagess returns all the DirectMessages items.
func (d *Data) DirectMessagess() ([]*DirectMessages, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessages](d, m.DataTypes.DirectMessages)
}

// EachDirectMessages calls fn for each DirectMessages item.
func (d *Data) EachDirectMessages(fn func(*DirectMessages) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessages](d, m.DataTypes.DirectMessages, fn)
}
