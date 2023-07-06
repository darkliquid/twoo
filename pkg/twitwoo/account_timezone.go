package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAccountTimezoneDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AccountTimezone",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AccountTimezone is the structure of the data/account-timezone.js file.
type AccountTimezone struct {
    // Fields go here
}

func (a *AccountTimezone) decode(el jsoniter.Any) {
	el = el.Get("accountTimezone")
	el.ToVal(a)
}

// AccountTimezones returns all the AccountTimezone items.
func (d *Data) AccountTimezones() ([]*AccountTimezone, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AccountTimezone](d, m.DataTypes.AccountTimezone)
}

// EachAccountTimezone calls fn for each AccountTimezone item.
func (d *Data) EachAccountTimezone(fn func(*AccountTimezone) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AccountTimezone](d, m.DataTypes.AccountTimezone, fn)
}
