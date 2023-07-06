package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDirectMessagesGroupDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DirectMessagesGroup",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DirectMessagesGroup is the structure of the data/direct-messages-group.js file.
type DirectMessagesGroup struct {
    // Fields go here
}

func (d *DirectMessagesGroup) decode(el jsoniter.Any) {
	el = el.Get("directMessagesGroup")
	el.ToVal(d)
}

// DirectMessagesGroups returns all the DirectMessagesGroup items.
func (d *Data) DirectMessagesGroups() ([]*DirectMessagesGroup, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DirectMessagesGroup](d, m.DataTypes.DirectMessagesGroup)
}

// EachDirectMessagesGroup calls fn for each DirectMessagesGroup item.
func (d *Data) EachDirectMessagesGroup(fn func(*DirectMessagesGroup) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DirectMessagesGroup](d, m.DataTypes.DirectMessagesGroup, fn)
}
