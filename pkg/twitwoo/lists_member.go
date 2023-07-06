package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerListsMemberDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ListsMember",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ListsMember is the structure of the data/lists-member.js file.
type ListsMember struct {
    // Fields go here
}

func (l *ListsMember) decode(el jsoniter.Any) {
	el = el.Get("listsMember")
	el.ToVal(l)
}

// ListsMembers returns all the ListsMember items.
func (d *Data) ListsMembers() ([]*ListsMember, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ListsMember](d, m.DataTypes.ListsMember)
}

// EachListsMember calls fn for each ListsMember item.
func (d *Data) EachListsMember(fn func(*ListsMember) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ListsMember](d, m.DataTypes.ListsMember, fn)
}
