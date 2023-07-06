package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerKeyRegistryDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.KeyRegistry",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// KeyRegistry is the structure of the data/key-registry.js file.
type KeyRegistry struct {
    // Fields go here
}

func (k *KeyRegistry) decode(el jsoniter.Any) {
	el = el.Get("keyRegistry")
	el.ToVal(k)
}

// KeyRegistrys returns all the KeyRegistry items.
func (d *Data) KeyRegistrys() ([]*KeyRegistry, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*KeyRegistry](d, m.DataTypes.KeyRegistry)
}

// EachKeyRegistry calls fn for each KeyRegistry item.
func (d *Data) EachKeyRegistry(fn func(*KeyRegistry) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*KeyRegistry](d, m.DataTypes.KeyRegistry, fn)
}
