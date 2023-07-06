package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAdMobileConversionsUnattributedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AdMobileConversionsUnattributed",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AdMobileConversionsUnattributed is the structure of the data/ad-mobile-conversions-unattributed.js file.
type AdMobileConversionsUnattributed struct {
    // Fields go here
}

func (a *AdMobileConversionsUnattributed) decode(el jsoniter.Any) {
	el = el.Get("adMobileConversionsUnattributed")
	el.ToVal(a)
}

// AdMobileConversionsUnattributeds returns all the AdMobileConversionsUnattributed items.
func (d *Data) AdMobileConversionsUnattributeds() ([]*AdMobileConversionsUnattributed, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AdMobileConversionsUnattributed](d, m.DataTypes.AdMobileConversionsUnattributed)
}

// EachAdMobileConversionsUnattributed calls fn for each AdMobileConversionsUnattributed item.
func (d *Data) EachAdMobileConversionsUnattributed(fn func(*AdMobileConversionsUnattributed) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AdMobileConversionsUnattributed](d, m.DataTypes.AdMobileConversionsUnattributed, fn)
}
