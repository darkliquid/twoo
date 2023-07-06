package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerProductDropDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ProductDrop",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ProductDrop is the structure of the data/product-drop.js file.
type ProductDrop struct {
    // Fields go here
}

func (p *ProductDrop) decode(el jsoniter.Any) {
	el = el.Get("productDrop")
	el.ToVal(p)
}

// ProductDrops returns all the ProductDrop items.
func (d *Data) ProductDrops() ([]*ProductDrop, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ProductDrop](d, m.DataTypes.ProductDrop)
}

// EachProductDrop calls fn for each ProductDrop item.
func (d *Data) EachProductDrop(fn func(*ProductDrop) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ProductDrop](d, m.DataTypes.ProductDrop, fn)
}
