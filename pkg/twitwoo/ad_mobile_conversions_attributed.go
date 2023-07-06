package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAdMobileConversionsAttributedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AdMobileConversionsAttributed",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AdMobileConversionsAttributed is the structure of the data/ad-mobile-conversions-attributed.js file.
type AdMobileConversionsAttributed struct {
    // Fields go here
}

func (a *AdMobileConversionsAttributed) decode(el jsoniter.Any) {
	el = el.Get("adMobileConversionsAttributed")
	el.ToVal(a)
}

// AdMobileConversionsAttributeds returns all the AdMobileConversionsAttributed items.
func (d *Data) AdMobileConversionsAttributeds() ([]*AdMobileConversionsAttributed, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AdMobileConversionsAttributed](d, m.DataTypes.AdMobileConversionsAttributed)
}

// EachAdMobileConversionsAttributed calls fn for each AdMobileConversionsAttributed item.
func (d *Data) EachAdMobileConversionsAttributed(fn func(*AdMobileConversionsAttributed) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AdMobileConversionsAttributed](d, m.DataTypes.AdMobileConversionsAttributed, fn)
}
