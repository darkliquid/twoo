package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerProductSetDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ProductSet",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ProductSet is the structure of the data/product-set.js file.
type ProductSet struct {
    // Fields go here
}

func (p *ProductSet) decode(el jsoniter.Any) {
	el = el.Get("productSet")
	el.ToVal(p)
}

// ProductSets returns all the ProductSet items.
func (d *Data) ProductSets() ([]*ProductSet, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ProductSet](d, m.DataTypes.ProductSet)
}

// EachProductSet calls fn for each ProductSet item.
func (d *Data) EachProductSet(fn func(*ProductSet) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ProductSet](d, m.DataTypes.ProductSet, fn)
}
