package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerListsSubscribedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ListsSubscribed",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ListsSubscribed is the structure of the data/lists-subscribed.js file.
type ListsSubscribed struct {
    // Fields go here
}

func (l *ListsSubscribed) decode(el jsoniter.Any) {
	el = el.Get("listsSubscribed")
	el.ToVal(l)
}

// ListsSubscribeds returns all the ListsSubscribed items.
func (d *Data) ListsSubscribeds() ([]*ListsSubscribed, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ListsSubscribed](d, m.DataTypes.ListsSubscribed)
}

// EachListsSubscribed calls fn for each ListsSubscribed item.
func (d *Data) EachListsSubscribed(fn func(*ListsSubscribed) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ListsSubscribed](d, m.DataTypes.ListsSubscribed, fn)
}
