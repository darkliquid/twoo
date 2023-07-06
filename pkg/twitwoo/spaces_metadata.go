package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerSpacesMetadataDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.SpacesMetadata",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// SpacesMetadata is the structure of the data/spaces-metadata.js file.
type SpacesMetadata struct {
    // Fields go here
}

func (s *SpacesMetadata) decode(el jsoniter.Any) {
	el = el.Get("spacesMetadata")
	el.ToVal(s)
}

// SpacesMetadatas returns all the SpacesMetadata items.
func (d *Data) SpacesMetadatas() ([]*SpacesMetadata, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*SpacesMetadata](d, m.DataTypes.SpacesMetadata)
}

// EachSpacesMetadata calls fn for each SpacesMetadata item.
func (d *Data) EachSpacesMetadata(fn func(*SpacesMetadata) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*SpacesMetadata](d, m.DataTypes.SpacesMetadata, fn)
}
