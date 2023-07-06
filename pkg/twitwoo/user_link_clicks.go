package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerUserLinkClicksDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.UserLinkClicks",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// UserLinkClicks is the structure of the data/user-link-clicks.js file.
type UserLinkClicks struct {
    // Fields go here
}

func (u *UserLinkClicks) decode(el jsoniter.Any) {
	el = el.Get("userLinkClicks")
	el.ToVal(u)
}

// UserLinkClickss returns all the UserLinkClicks items.
func (d *Data) UserLinkClickss() ([]*UserLinkClicks, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*UserLinkClicks](d, m.DataTypes.UserLinkClicks)
}

// EachUserLinkClicks calls fn for each UserLinkClicks item.
func (d *Data) EachUserLinkClicks(fn func(*UserLinkClicks) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*UserLinkClicks](d, m.DataTypes.UserLinkClicks, fn)
}
