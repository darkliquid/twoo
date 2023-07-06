package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerListsCreatedDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ListsCreated",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// ListsCreated is the structure of the data/lists-created.js file.
type ListsCreated struct {
    // Fields go here
}

func (l *ListsCreated) decode(el jsoniter.Any) {
	el = el.Get("listsCreated")
	el.ToVal(l)
}

// ListsCreateds returns all the ListsCreated items.
func (d *Data) ListsCreateds() ([]*ListsCreated, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*ListsCreated](d, m.DataTypes.ListsCreated)
}

// EachListsCreated calls fn for each ListsCreated item.
func (d *Data) EachListsCreated(fn func(*ListsCreated) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*ListsCreated](d, m.DataTypes.ListsCreated, fn)
}
