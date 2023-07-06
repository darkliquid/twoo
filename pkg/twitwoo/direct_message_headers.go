package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessageHeadersDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessageHeaders",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessageHeaders is the structure of the data/direct-message-headers.js file.
type DirectMessageHeaders struct {
    // Fields go here
}

func (d *DirectMessageHeaders) decode(el jsoniter.Any) {
	el = el.Get("directMessageHeaders")
	el.ToVal(d)
}

// DirectMessageHeaderss returns all the DirectMessageHeaders items.
func (d *Data) DirectMessageHeaderss() ([]*DirectMessageHeaders, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessageHeaders](d, m.DataTypes.DirectMessageHeaders)
}

// EachDirectMessageHeaders calls fn for each DirectMessageHeaders item.
func (d *Data) EachDirectMessageHeaders(fn func(*DirectMessageHeaders) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessageHeaders](d, m.DataTypes.DirectMessageHeaders, fn)
}
