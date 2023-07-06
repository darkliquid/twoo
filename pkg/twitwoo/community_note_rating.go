package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCommunityNoteRatingDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CommunityNoteRating",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CommunityNoteRating is the structure of the data/community-note-rating.js file.
type CommunityNoteRating struct {
    // Fields go here
}

func (c *CommunityNoteRating) decode(el jsoniter.Any) {
	el = el.Get("communityNoteRating")
	el.ToVal(c)
}

// CommunityNoteRatings returns all the CommunityNoteRating items.
func (d *Data) CommunityNoteRatings() ([]*CommunityNoteRating, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CommunityNoteRating](d, m.DataTypes.CommunityNoteRating)
}

// EachCommunityNoteRating calls fn for each CommunityNoteRating item.
func (d *Data) EachCommunityNoteRating(fn func(*CommunityNoteRating) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CommunityNoteRating](d, m.DataTypes.CommunityNoteRating, fn)
}
