package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerPeriscopeBroadcastMetadataDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.PeriscopeBroadcastMetadata",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// PeriscopeBroadcastMetadata is the structure of the data/periscope-broadcast-metadata.js file.
type PeriscopeBroadcastMetadata struct {
    // Fields go here
}

func (p *PeriscopeBroadcastMetadata) decode(el jsoniter.Any) {
	el = el.Get("periscopeBroadcastMetadata")
	el.ToVal(p)
}

// PeriscopeBroadcastMetadatas returns all the PeriscopeBroadcastMetadata items.
func (d *Data) PeriscopeBroadcastMetadatas() ([]*PeriscopeBroadcastMetadata, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*PeriscopeBroadcastMetadata](d, m.DataTypes.PeriscopeBroadcastMetadata)
}

// EachPeriscopeBroadcastMetadata calls fn for each PeriscopeBroadcastMetadata item.
func (d *Data) EachPeriscopeBroadcastMetadata(fn func(*PeriscopeBroadcastMetadata) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*PeriscopeBroadcastMetadata](d, m.DataTypes.PeriscopeBroadcastMetadata, fn)
}
