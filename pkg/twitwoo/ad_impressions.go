package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAdImpressionsDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AdImpressions",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AdImpressions is the structure of the data/ad-impressions.js file.
type AdImpressions struct {
    // Fields go here
}

func (a *AdImpressions) decode(el jsoniter.Any) {
	el = el.Get("adImpressions")
	el.ToVal(a)
}

// AdImpressionss returns all the AdImpressions items.
func (d *Data) AdImpressionss() ([]*AdImpressions, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AdImpressions](d, m.DataTypes.AdImpressions)
}

// EachAdImpressions calls fn for each AdImpressions item.
func (d *Data) EachAdImpressions(fn func(*AdImpressions) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AdImpressions](d, m.DataTypes.AdImpressions, fn)
}
