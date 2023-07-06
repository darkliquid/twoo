package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCommerceCatalogDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CommerceCatalog",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CommerceCatalog is the structure of the data/commerce-catalog.js file.
type CommerceCatalog struct {
    // Fields go here
}

func (c *CommerceCatalog) decode(el jsoniter.Any) {
	el = el.Get("commerceCatalog")
	el.ToVal(c)
}

// CommerceCatalogs returns all the CommerceCatalog items.
func (d *Data) CommerceCatalogs() ([]*CommerceCatalog, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CommerceCatalog](d, m.DataTypes.CommerceCatalog)
}

// EachCommerceCatalog calls fn for each CommerceCatalog item.
func (d *Data) EachCommerceCatalog(fn func(*CommerceCatalog) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CommerceCatalog](d, m.DataTypes.CommerceCatalog, fn)
}
