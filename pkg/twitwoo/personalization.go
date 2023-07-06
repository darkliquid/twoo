package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPersonalizationDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Personalization",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Personalization is the structure of the data/personalization.js file.
type Personalization struct {
    // Fields go here
}

func (p *Personalization) decode(el jsoniter.Any) {
	el = el.Get("personalization")
	el.ToVal(p)
}

// Personalizations returns all the Personalization items.
func (d *Data) Personalizations() ([]*Personalization, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Personalization](d, m.DataTypes.Personalization)
}

// EachPersonalization calls fn for each Personalization item.
func (d *Data) EachPersonalization(fn func(*Personalization) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Personalization](d, m.DataTypes.Personalization, fn)
}
