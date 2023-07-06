package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAccountLabelDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AccountLabel",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AccountLabel is the structure of the data/account-label.js file.
type AccountLabel struct {
    // Fields go here
}

func (a *AccountLabel) decode(el jsoniter.Any) {
	el = el.Get("accountLabel")
	el.ToVal(a)
}

// AccountLabels returns all the AccountLabel items.
func (d *Data) AccountLabels() ([]*AccountLabel, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AccountLabel](d, m.DataTypes.AccountLabel)
}

// EachAccountLabel calls fn for each AccountLabel item.
func (d *Data) EachAccountLabel(fn func(*AccountLabel) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AccountLabel](d, m.DataTypes.AccountLabel, fn)
}
