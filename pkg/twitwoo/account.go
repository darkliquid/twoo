package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAccountDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Account",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Account is the structure of the data/account.js file.
type Account struct {
    // Fields go here
}

func (a *Account) decode(el jsoniter.Any) {
	el = el.Get("account")
	el.ToVal(a)
}

// Accounts returns all the Account items.
func (d *Data) Accounts() ([]*Account, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Account](d, m.DataTypes.Account)
}

// EachAccount calls fn for each Account item.
func (d *Data) EachAccount(fn func(*Account) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Account](d, m.DataTypes.Account, fn)
}
