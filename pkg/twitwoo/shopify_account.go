package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerShopifyAccountDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ShopifyAccount",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ShopifyAccount is the structure of the data/shopify-account.js file.
type ShopifyAccount struct {
    // Fields go here
}

func (s *ShopifyAccount) decode(el jsoniter.Any) {
	el = el.Get("shopifyAccount")
	el.ToVal(s)
}

// ShopifyAccounts returns all the ShopifyAccount items.
func (d *Data) ShopifyAccounts() ([]*ShopifyAccount, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ShopifyAccount](d, m.DataTypes.ShopifyAccount)
}

// EachShopifyAccount calls fn for each ShopifyAccount item.
func (d *Data) EachShopifyAccount(fn func(*ShopifyAccount) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ShopifyAccount](d, m.DataTypes.ShopifyAccount, fn)
}
