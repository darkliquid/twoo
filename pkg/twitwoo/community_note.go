package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCommunityNoteDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CommunityNote",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CommunityNote is the structure of the data/community-note.js file.
type CommunityNote struct {
    // Fields go here
}

func (c *CommunityNote) decode(el jsoniter.Any) {
	el = el.Get("communityNote")
	el.ToVal(c)
}

// CommunityNotes returns all the CommunityNote items.
func (d *Data) CommunityNotes() ([]*CommunityNote, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CommunityNote](d, m.DataTypes.CommunityNote)
}

// EachCommunityNote calls fn for each CommunityNote item.
func (d *Data) EachCommunityNote(fn func(*CommunityNote) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CommunityNote](d, m.DataTypes.CommunityNote, fn)
}
