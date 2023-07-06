package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeProfileDescriptionDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeProfileDescription",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeProfileDescription is the structure of the data/periscope-profile-description.js file.
type PeriscopeProfileDescription struct {
    // Fields go here
}

func (p *PeriscopeProfileDescription) decode(el jsoniter.Any) {
	el = el.Get("periscopeProfileDescription")
	el.ToVal(p)
}

// PeriscopeProfileDescriptions returns all the PeriscopeProfileDescription items.
func (d *Data) PeriscopeProfileDescriptions() ([]*PeriscopeProfileDescription, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeProfileDescription](d, m.DataTypes.PeriscopeProfileDescription)
}

// EachPeriscopeProfileDescription calls fn for each PeriscopeProfileDescription item.
func (d *Data) EachPeriscopeProfileDescription(fn func(*PeriscopeProfileDescription) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeProfileDescription](d, m.DataTypes.PeriscopeProfileDescription, fn)
}
