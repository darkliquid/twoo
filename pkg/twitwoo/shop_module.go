package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerShopModuleDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ShopModule",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ShopModule is the structure of the data/shop-module.js file.
type ShopModule struct {
    // Fields go here
}

func (s *ShopModule) decode(el jsoniter.Any) {
	el = el.Get("shopModule")
	el.ToVal(s)
}

// ShopModules returns all the ShopModule items.
func (d *Data) ShopModules() ([]*ShopModule, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ShopModule](d, m.DataTypes.ShopModule)
}

// EachShopModule calls fn for each ShopModule item.
func (d *Data) EachShopModule(fn func(*ShopModule) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ShopModule](d, m.DataTypes.ShopModule, fn)
}
