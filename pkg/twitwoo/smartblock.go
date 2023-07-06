package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerSmartblockDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Smartblock",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Smartblock is the structure of the data/smartblock.js file.
type Smartblock struct {
    // Fields go here
}

func (s *Smartblock) decode(el jsoniter.Any) {
	el = el.Get("smartblock")
	el.ToVal(s)
}

// Smartblocks returns all the Smartblock items.
func (d *Data) Smartblocks() ([]*Smartblock, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Smartblock](d, m.DataTypes.Smartblock)
}

// EachSmartblock calls fn for each Smartblock item.
func (d *Data) EachSmartblock(fn func(*Smartblock) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Smartblock](d, m.DataTypes.Smartblock, fn)
}
