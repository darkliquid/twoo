package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerScreenNameChangeDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ScreenNameChange",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ScreenNameChange is the structure of the data/screen-name-change.js file.
type ScreenNameChange struct {
    // Fields go here
}

func (s *ScreenNameChange) decode(el jsoniter.Any) {
	el = el.Get("screenNameChange")
	el.ToVal(s)
}

// ScreenNameChanges returns all the ScreenNameChange items.
func (d *Data) ScreenNameChanges() ([]*ScreenNameChange, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ScreenNameChange](d, m.DataTypes.ScreenNameChange)
}

// EachScreenNameChange calls fn for each ScreenNameChange item.
func (d *Data) EachScreenNameChange(fn func(*ScreenNameChange) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ScreenNameChange](d, m.DataTypes.ScreenNameChange, fn)
}
