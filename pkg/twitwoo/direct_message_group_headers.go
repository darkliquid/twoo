package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessageGroupHeadersDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessageGroupHeaders",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessageGroupHeaders is the structure of the data/direct-message-group-headers.js file.
type DirectMessageGroupHeaders struct {
    // Fields go here
}

func (d *DirectMessageGroupHeaders) decode(el jsoniter.Any) {
	el = el.Get("directMessageGroupHeaders")
	el.ToVal(d)
}

// DirectMessageGroupHeaderss returns all the DirectMessageGroupHeaders items.
func (d *Data) DirectMessageGroupHeaderss() ([]*DirectMessageGroupHeaders, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessageGroupHeaders](d, m.DataTypes.DirectMessageGroupHeaders)
}

// EachDirectMessageGroupHeaders calls fn for each DirectMessageGroupHeaders item.
func (d *Data) EachDirectMessageGroupHeaders(fn func(*DirectMessageGroupHeaders) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessageGroupHeaders](d, m.DataTypes.DirectMessageGroupHeaders, fn)
}
