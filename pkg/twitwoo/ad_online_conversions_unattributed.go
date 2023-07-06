package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAdOnlineConversionsUnattributedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AdOnlineConversionsUnattributed",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AdOnlineConversionsUnattributed is the structure of the data/ad-online-conversions-unattributed.js file.
type AdOnlineConversionsUnattributed struct {
    // Fields go here
}

func (a *AdOnlineConversionsUnattributed) decode(el jsoniter.Any) {
	el = el.Get("adOnlineConversionsUnattributed")
	el.ToVal(a)
}

// AdOnlineConversionsUnattributeds returns all the AdOnlineConversionsUnattributed items.
func (d *Data) AdOnlineConversionsUnattributeds() ([]*AdOnlineConversionsUnattributed, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AdOnlineConversionsUnattributed](d, m.DataTypes.AdOnlineConversionsUnattributed)
}

// EachAdOnlineConversionsUnattributed calls fn for each AdOnlineConversionsUnattributed item.
func (d *Data) EachAdOnlineConversionsUnattributed(fn func(*AdOnlineConversionsUnattributed) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AdOnlineConversionsUnattributed](d, m.DataTypes.AdOnlineConversionsUnattributed, fn)
}
