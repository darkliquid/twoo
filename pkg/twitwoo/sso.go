package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerSSODecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.SSO",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// SSO is the structure of the data/sso.js file.
type SSO struct {
    // Fields go here
}

func (s *SSO) decode(el jsoniter.Any) {
	el = el.Get("sSo")
	el.ToVal(s)
}

// SSOs returns all the SSO items.
func (d *Data) SSOs() ([]*SSO, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*SSO](d, m.DataTypes.SSO)
}

// EachSSO calls fn for each SSO item.
func (d *Data) EachSSO(fn func(*SSO) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*SSO](d, m.DataTypes.SSO, fn)
}
