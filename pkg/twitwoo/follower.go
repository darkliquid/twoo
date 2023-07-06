package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerFollowerDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Follower",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Follower is the structure of the data/follower.js file.
type Follower struct {
    // Fields go here
}

func (f *Follower) decode(el jsoniter.Any) {
	el = el.Get("follower")
	el.ToVal(f)
}

// Followers returns all the Follower items.
func (d *Data) Followers() ([]*Follower, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Follower](d, m.DataTypes.Follower)
}

// EachFollower calls fn for each Follower item.
func (d *Data) EachFollower(fn func(*Follower) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Follower](d, m.DataTypes.Follower, fn)
}
