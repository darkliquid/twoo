package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeBanInformationDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeBanInformation",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeBanInformation is the structure of the data/periscope-ban-information.js file.
type PeriscopeBanInformation struct {
    // Fields go here
}

func (p *PeriscopeBanInformation) decode(el jsoniter.Any) {
	el = el.Get("periscopeBanInformation")
	el.ToVal(p)
}

// PeriscopeBanInformations returns all the PeriscopeBanInformation items.
func (d *Data) PeriscopeBanInformations() ([]*PeriscopeBanInformation, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeBanInformation](d, m.DataTypes.PeriscopeBanInformation)
}

// EachPeriscopeBanInformation calls fn for each PeriscopeBanInformation item.
func (d *Data) EachPeriscopeBanInformation(fn func(*PeriscopeBanInformation) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeBanInformation](d, m.DataTypes.PeriscopeBanInformation, fn)
}
