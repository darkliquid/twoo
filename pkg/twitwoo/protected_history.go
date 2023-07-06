package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerProtectedHistoryDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ProtectedHistory",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ProtectedHistory is the structure of the data/protected-history.js file.
type ProtectedHistory struct {
    // Fields go here
}

func (p *ProtectedHistory) decode(el jsoniter.Any) {
	el = el.Get("protectedHistory")
	el.ToVal(p)
}

// ProtectedHistorys returns all the ProtectedHistory items.
func (d *Data) ProtectedHistorys() ([]*ProtectedHistory, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ProtectedHistory](d, m.DataTypes.ProtectedHistory)
}

// EachProtectedHistory calls fn for each ProtectedHistory item.
func (d *Data) EachProtectedHistory(fn func(*ProtectedHistory) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ProtectedHistory](d, m.DataTypes.ProtectedHistory, fn)
}
