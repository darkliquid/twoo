package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerCatalogItemDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.CatalogItem",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// CatalogItem is the structure of the data/catalog-item.js file.
type CatalogItem struct {
    // Fields go here
}

func (c *CatalogItem) decode(el jsoniter.Any) {
	el = el.Get("catalogItem")
	el.ToVal(c)
}

// CatalogItems returns all the CatalogItem items.
func (d *Data) CatalogItems() ([]*CatalogItem, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*CatalogItem](d, m.DataTypes.CatalogItem)
}

// EachCatalogItem calls fn for each CatalogItem item.
func (d *Data) EachCatalogItem(fn func(*CatalogItem) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*CatalogItem](d, m.DataTypes.CatalogItem, fn)
}
