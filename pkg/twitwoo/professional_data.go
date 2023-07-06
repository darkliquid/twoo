package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerProfessionalDataDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ProfessionalData",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ProfessionalData is the structure of the data/professional-data.js file.
type ProfessionalData struct {
    // Fields go here
}

func (p *ProfessionalData) decode(el jsoniter.Any) {
	el = el.Get("professionalData")
	el.ToVal(p)
}

// ProfessionalDatas returns all the ProfessionalData items.
func (d *Data) ProfessionalDatas() ([]*ProfessionalData, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ProfessionalData](d, m.DataTypes.ProfessionalData)
}

// EachProfessionalData calls fn for each ProfessionalData item.
func (d *Data) EachProfessionalData(fn func(*ProfessionalData) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ProfessionalData](d, m.DataTypes.ProfessionalData, fn)
}
