package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeAccountInformationDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeAccountInformation",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeAccountInformation is the structure of the data/periscope-account-information.js file.
type PeriscopeAccountInformation struct {
    // Fields go here
}

func (p *PeriscopeAccountInformation) decode(el jsoniter.Any) {
	el = el.Get("periscopeAccountInformation")
	el.ToVal(p)
}

// PeriscopeAccountInformations returns all the PeriscopeAccountInformation items.
func (d *Data) PeriscopeAccountInformations() ([]*PeriscopeAccountInformation, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeAccountInformation](d, m.DataTypes.PeriscopeAccountInformation)
}

// EachPeriscopeAccountInformation calls fn for each PeriscopeAccountInformation item.
func (d *Data) EachPeriscopeAccountInformation(fn func(*PeriscopeAccountInformation) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeAccountInformation](d, m.DataTypes.PeriscopeAccountInformation, fn)
}
