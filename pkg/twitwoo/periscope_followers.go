package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeFollowersDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeFollowers",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeFollowers is the structure of the data/periscope-followers.js file.
type PeriscopeFollowers struct {
    // Fields go here
}

func (p *PeriscopeFollowers) decode(el jsoniter.Any) {
	el = el.Get("periscopeFollowers")
	el.ToVal(p)
}

// PeriscopeFollowerss returns all the PeriscopeFollowers items.
func (d *Data) PeriscopeFollowerss() ([]*PeriscopeFollowers, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeFollowers](d, m.DataTypes.PeriscopeFollowers)
}

// EachPeriscopeFollowers calls fn for each PeriscopeFollowers item.
func (d *Data) EachPeriscopeFollowers(fn func(*PeriscopeFollowers) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeFollowers](d, m.DataTypes.PeriscopeFollowers, fn)
}
