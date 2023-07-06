package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPhoneNumberDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PhoneNumber",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PhoneNumber is the structure of the data/phone-number.js file.
type PhoneNumber struct {
    // Fields go here
}

func (p *PhoneNumber) decode(el jsoniter.Any) {
	el = el.Get("phoneNumber")
	el.ToVal(p)
}

// PhoneNumbers returns all the PhoneNumber items.
func (d *Data) PhoneNumbers() ([]*PhoneNumber, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PhoneNumber](d, m.DataTypes.PhoneNumber)
}

// EachPhoneNumber calls fn for each PhoneNumber item.
func (d *Data) EachPhoneNumber(fn func(*PhoneNumber) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PhoneNumber](d, m.DataTypes.PhoneNumber, fn)
}
