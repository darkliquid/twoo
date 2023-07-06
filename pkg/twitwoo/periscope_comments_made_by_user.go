package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeCommentsMadeByUserDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeCommentsMadeByUser",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeCommentsMadeByUser is the structure of the data/periscope-comments-made-by-user.js file.
type PeriscopeCommentsMadeByUser struct {
    // Fields go here
}

func (p *PeriscopeCommentsMadeByUser) decode(el jsoniter.Any) {
	el = el.Get("periscopeCommentsMadeByUser")
	el.ToVal(p)
}

// PeriscopeCommentsMadeByUsers returns all the PeriscopeCommentsMadeByUser items.
func (d *Data) PeriscopeCommentsMadeByUsers() ([]*PeriscopeCommentsMadeByUser, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeCommentsMadeByUser](d, m.DataTypes.PeriscopeCommentsMadeByUser)
}

// EachPeriscopeCommentsMadeByUser calls fn for each PeriscopeCommentsMadeByUser item.
func (d *Data) EachPeriscopeCommentsMadeByUser(fn func(*PeriscopeCommentsMadeByUser) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeCommentsMadeByUser](d, m.DataTypes.PeriscopeCommentsMadeByUser, fn)
}
