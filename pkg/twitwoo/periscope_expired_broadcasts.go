package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeExpiredBroadcastsDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeExpiredBroadcasts",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeExpiredBroadcasts is the structure of the data/periscope-expired-broadcasts.js file.
type PeriscopeExpiredBroadcasts struct {
    // Fields go here
}

func (p *PeriscopeExpiredBroadcasts) decode(el jsoniter.Any) {
	el = el.Get("periscopeExpiredBroadcasts")
	el.ToVal(p)
}

// PeriscopeExpiredBroadcastss returns all the PeriscopeExpiredBroadcasts items.
func (d *Data) PeriscopeExpiredBroadcastss() ([]*PeriscopeExpiredBroadcasts, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeExpiredBroadcasts](d, m.DataTypes.PeriscopeExpiredBroadcasts)
}

// EachPeriscopeExpiredBroadcasts calls fn for each PeriscopeExpiredBroadcasts item.
func (d *Data) EachPeriscopeExpiredBroadcasts(fn func(*PeriscopeExpiredBroadcasts) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeExpiredBroadcasts](d, m.DataTypes.PeriscopeExpiredBroadcasts, fn)
}
