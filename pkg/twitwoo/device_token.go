package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerDeviceTokenDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DeviceToken",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// DeviceToken is the structure of the data/device-token.js file.
type DeviceToken struct {
    // Fields go here
}

func (d *DeviceToken) decode(el jsoniter.Any) {
	el = el.Get("deviceToken")
	el.ToVal(d)
}

// DeviceTokens returns all the DeviceToken items.
func (d *Data) DeviceTokens() ([]*DeviceToken, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*DeviceToken](d, m.DataTypes.DeviceToken)
}

// EachDeviceToken calls fn for each DeviceToken item.
func (d *Data) EachDeviceToken(fn func(*DeviceToken) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*DeviceToken](d, m.DataTypes.DeviceToken, fn)
}
