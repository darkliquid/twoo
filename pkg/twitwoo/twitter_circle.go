package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterCircleDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterCircle",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterCircle is the structure of the data/twitter-circle.js file.
type TwitterCircle struct {
    // Fields go here
}

func (t *TwitterCircle) decode(el jsoniter.Any) {
	el = el.Get("twitterCircle")
	el.ToVal(t)
}

// TwitterCircles returns all the TwitterCircle items.
func (d *Data) TwitterCircles() ([]*TwitterCircle, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterCircle](d, m.DataTypes.TwitterCircle)
}

// EachTwitterCircle calls fn for each TwitterCircle item.
func (d *Data) EachTwitterCircle(fn func(*TwitterCircle) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterCircle](d, m.DataTypes.TwitterCircle, fn)
}
