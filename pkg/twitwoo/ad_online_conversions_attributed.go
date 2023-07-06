package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAdOnlineConversionsAttributedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AdOnlineConversionsAttributed",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AdOnlineConversionsAttributed is the structure of the data/ad-online-conversions-attributed.js file.
type AdOnlineConversionsAttributed struct {
    // Fields go here
}

func (a *AdOnlineConversionsAttributed) decode(el jsoniter.Any) {
	el = el.Get("adOnlineConversionsAttributed")
	el.ToVal(a)
}

// AdOnlineConversionsAttributeds returns all the AdOnlineConversionsAttributed items.
func (d *Data) AdOnlineConversionsAttributeds() ([]*AdOnlineConversionsAttributed, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AdOnlineConversionsAttributed](d, m.DataTypes.AdOnlineConversionsAttributed)
}

// EachAdOnlineConversionsAttributed calls fn for each AdOnlineConversionsAttributed item.
func (d *Data) EachAdOnlineConversionsAttributed(fn func(*AdOnlineConversionsAttributed) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AdOnlineConversionsAttributed](d, m.DataTypes.AdOnlineConversionsAttributed, fn)
}
