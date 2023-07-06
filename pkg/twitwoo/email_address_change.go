package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerEmailAddressChangeDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.EmailAddressChange",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// EmailAddressChange is the structure of the data/email-address-change.js file.
type EmailAddressChange struct {
    // Fields go here
}

func (e *EmailAddressChange) decode(el jsoniter.Any) {
	el = el.Get("emailAddressChange")
	el.ToVal(e)
}

// EmailAddressChanges returns all the EmailAddressChange items.
func (d *Data) EmailAddressChanges() ([]*EmailAddressChange, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*EmailAddressChange](d, m.DataTypes.EmailAddressChange)
}

// EachEmailAddressChange calls fn for each EmailAddressChange item.
func (d *Data) EachEmailAddressChange(fn func(*EmailAddressChange) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*EmailAddressChange](d, m.DataTypes.EmailAddressChange, fn)
}
