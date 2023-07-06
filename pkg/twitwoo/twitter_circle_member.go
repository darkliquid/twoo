package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterCircleMemberDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterCircleMember",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterCircleMember is the structure of the data/twitter-circle-member.js file.
type TwitterCircleMember struct {
    // Fields go here
}

func (t *TwitterCircleMember) decode(el jsoniter.Any) {
	el = el.Get("twitterCircleMember")
	el.ToVal(t)
}

// TwitterCircleMembers returns all the TwitterCircleMember items.
func (d *Data) TwitterCircleMembers() ([]*TwitterCircleMember, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterCircleMember](d, m.DataTypes.TwitterCircleMember)
}

// EachTwitterCircleMember calls fn for each TwitterCircleMember item.
func (d *Data) EachTwitterCircleMember(fn func(*TwitterCircleMember) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterCircleMember](d, m.DataTypes.TwitterCircleMember, fn)
}
