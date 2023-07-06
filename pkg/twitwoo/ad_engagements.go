package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAdEngagementsDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AdEngagements",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AdEngagements is the structure of the data/ad-engagements.js file.
type AdEngagements struct {
    // Fields go here
}

func (a *AdEngagements) decode(el jsoniter.Any) {
	el = el.Get("adEngagements")
	el.ToVal(a)
}

// AdEngagementss returns all the AdEngagements items.
func (d *Data) AdEngagementss() ([]*AdEngagements, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AdEngagements](d, m.DataTypes.AdEngagements)
}

// EachAdEngagements calls fn for each AdEngagements item.
func (d *Data) EachAdEngagements(fn func(*AdEngagements) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AdEngagements](d, m.DataTypes.AdEngagements, fn)
}
