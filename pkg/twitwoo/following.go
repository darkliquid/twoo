package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerFollowingDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Following",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Following is the structure of the data/following.js file.
type Following struct {
    // Fields go here
}

func (f *Following) decode(el jsoniter.Any) {
	el = el.Get("following")
	el.ToVal(f)
}

// Followings returns all the Following items.
func (d *Data) Followings() ([]*Following, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Following](d, m.DataTypes.Following)
}

// EachFollowing calls fn for each Following item.
func (d *Data) EachFollowing(fn func(*Following) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Following](d, m.DataTypes.Following, fn)
}
