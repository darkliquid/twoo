package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerVerifiedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Verified",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Verified is the structure of the data/verified.js file.
type Verified struct {
    // Fields go here
}

func (v *Verified) decode(el jsoniter.Any) {
	el = el.Get("verified")
	el.ToVal(v)
}

// Verifieds returns all the Verified items.
func (d *Data) Verifieds() ([]*Verified, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Verified](d, m.DataTypes.Verified)
}

// EachVerified calls fn for each Verified item.
func (d *Data) EachVerified(fn func(*Verified) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Verified](d, m.DataTypes.Verified, fn)
}
