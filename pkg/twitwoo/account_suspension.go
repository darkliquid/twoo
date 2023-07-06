package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAccountSuspensionDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AccountSuspension",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AccountSuspension is the structure of the data/account-suspension.js file.
type AccountSuspension struct {
    // Fields go here
}

func (a *AccountSuspension) decode(el jsoniter.Any) {
	el = el.Get("accountSuspension")
	el.ToVal(a)
}

// AccountSuspensions returns all the AccountSuspension items.
func (d *Data) AccountSuspensions() ([]*AccountSuspension, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AccountSuspension](d, m.DataTypes.AccountSuspension)
}

// EachAccountSuspension calls fn for each AccountSuspension item.
func (d *Data) EachAccountSuspension(fn func(*AccountSuspension) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AccountSuspension](d, m.DataTypes.AccountSuspension, fn)
}
