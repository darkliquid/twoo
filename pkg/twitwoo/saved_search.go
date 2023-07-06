package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerSavedSearchDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.SavedSearch",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// SavedSearch is the structure of the data/saved-search.js file.
type SavedSearch struct {
    // Fields go here
}

func (s *SavedSearch) decode(el jsoniter.Any) {
	el = el.Get("savedSearch")
	el.ToVal(s)
}

// SavedSearchs returns all the SavedSearch items.
func (d *Data) SavedSearchs() ([]*SavedSearch, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*SavedSearch](d, m.DataTypes.SavedSearch)
}

// EachSavedSearch calls fn for each SavedSearch item.
func (d *Data) EachSavedSearch(fn func(*SavedSearch) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*SavedSearch](d, m.DataTypes.SavedSearch, fn)
}
