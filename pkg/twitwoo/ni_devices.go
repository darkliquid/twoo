package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerNIDevicesDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.NIDevices",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// NIDevices is the structure of the data/ni-devices.js file.
type NIDevices struct {
    // Fields go here
}

func (n *NIDevices) decode(el jsoniter.Any) {
	el = el.Get("nIdevices")
	el.ToVal(n)
}

// NIDevicess returns all the NIDevices items.
func (d *Data) NIDevicess() ([]*NIDevices, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*NIDevices](d, m.DataTypes.NIDevices)
}

// EachNIDevices calls fn for each NIDevices item.
func (d *Data) EachNIDevices(fn func(*NIDevices) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*NIDevices](d, m.DataTypes.NIDevices, fn)
}
