package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerAgeInfoDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.AgeInfo",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// AgeInfo is the structure of the data/age-info.js file.
type AgeInfo struct {
    // Fields go here
}

func (a *AgeInfo) decode(el jsoniter.Any) {
	el = el.Get("ageInfo")
	el.ToVal(a)
}

// AgeInfos returns all the AgeInfo items.
func (d *Data) AgeInfos() ([]*AgeInfo, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*AgeInfo](d, m.DataTypes.AgeInfo)
}

// EachAgeInfo calls fn for each AgeInfo item.
func (d *Data) EachAgeInfo(fn func(*AgeInfo) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*AgeInfo](d, m.DataTypes.AgeInfo, fn)
}
