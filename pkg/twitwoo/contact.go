package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerContactDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Contact",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Contact is the structure of the data/contact.js file.
type Contact struct {
    // Fields go here
}

func (c *Contact) decode(el jsoniter.Any) {
	el = el.Get("contact")
	el.ToVal(c)
}

// Contacts returns all the Contact items.
func (d *Data) Contacts() ([]*Contact, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Contact](d, m.DataTypes.Contact)
}

// EachContact calls fn for each Contact item.
func (d *Data) EachContact(fn func(*Contact) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Contact](d, m.DataTypes.Contact, fn)
}
