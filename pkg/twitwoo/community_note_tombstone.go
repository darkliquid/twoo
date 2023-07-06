package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCommunityNoteTombstoneDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CommunityNoteTombstone",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CommunityNoteTombstone is the structure of the data/community-note-tombstone.js file.
type CommunityNoteTombstone struct {
    // Fields go here
}

func (c *CommunityNoteTombstone) decode(el jsoniter.Any) {
	el = el.Get("communityNoteTombstone")
	el.ToVal(c)
}

// CommunityNoteTombstones returns all the CommunityNoteTombstone items.
func (d *Data) CommunityNoteTombstones() ([]*CommunityNoteTombstone, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CommunityNoteTombstone](d, m.DataTypes.CommunityNoteTombstone)
}

// EachCommunityNoteTombstone calls fn for each CommunityNoteTombstone item.
func (d *Data) EachCommunityNoteTombstone(fn func(*CommunityNoteTombstone) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CommunityNoteTombstone](d, m.DataTypes.CommunityNoteTombstone, fn)
}
