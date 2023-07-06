package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerMomentDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Moment",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Moment is the structure of the data/moment.js file.
type Moment struct {
    // Fields go here
}

func (m *Moment) decode(el jsoniter.Any) {
	el = el.Get("moment")
	el.ToVal(m)
}

// Moments returns all the Moment items.
func (d *Data) Moments() ([]*Moment, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Moment](d, m.DataTypes.Moment)
}

// EachMoment calls fn for each Moment item.
func (d *Data) EachMoment(fn func(*Moment) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Moment](d, m.DataTypes.Moment, fn)
}
