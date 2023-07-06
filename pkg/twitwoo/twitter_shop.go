package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterShopDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterShop",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterShop is the structure of the data/twitter-shop.js file.
type TwitterShop struct {
    // Fields go here
}

func (t *TwitterShop) decode(el jsoniter.Any) {
	el = el.Get("twitterShop")
	el.ToVal(t)
}

// TwitterShops returns all the TwitterShop items.
func (d *Data) TwitterShops() ([]*TwitterShop, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterShop](d, m.DataTypes.TwitterShop)
}

// EachTwitterShop calls fn for each TwitterShop item.
func (d *Data) EachTwitterShop(fn func(*TwitterShop) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterShop](d, m.DataTypes.TwitterShop, fn)
}
