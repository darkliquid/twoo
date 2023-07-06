package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerIPAuditDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.IPAudit",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// IPAudit is the structure of the data/ip-audit.js file.
type IPAudit struct {
    // Fields go here
}

func (i *IPAudit) decode(el jsoniter.Any) {
	el = el.Get("iPaudit")
	el.ToVal(i)
}

// IPAudits returns all the IPAudit items.
func (d *Data) IPAudits() ([]*IPAudit, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*IPAudit](d, m.DataTypes.IPAudit)
}

// EachIPAudit calls fn for each IPAudit item.
func (d *Data) EachIPAudit(fn func(*IPAudit) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*IPAudit](d, m.DataTypes.IPAudit, fn)
}
