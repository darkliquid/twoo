package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerConnectedApplicationDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ConnectedApplication",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ConnectedApplication is the structure of the data/connected-application.js file.
type ConnectedApplication struct {
    // Fields go here
}

func (c *ConnectedApplication) decode(el jsoniter.Any) {
	el = el.Get("connectedApplication")
	el.ToVal(c)
}

// ConnectedApplications returns all the ConnectedApplication items.
func (d *Data) ConnectedApplications() ([]*ConnectedApplication, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ConnectedApplication](d, m.DataTypes.ConnectedApplication)
}

// EachConnectedApplication calls fn for each ConnectedApplication item.
func (d *Data) EachConnectedApplication(fn func(*ConnectedApplication) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ConnectedApplication](d, m.DataTypes.ConnectedApplication, fn)
}
