package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerBlockDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.Block",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// Block is the structure of the data/block.js file.
type Block struct {
    // Fields go here
}

func (b *Block) decode(el jsoniter.Any) {
	el = el.Get("block")
	el.ToVal(b)
}

// Blocks returns all the Block items.
func (d *Data) Blocks() ([]*Block, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Block](d, m.DataTypes.Block)
}

// EachBlock calls fn for each Block item.
func (d *Data) EachBlock(fn func(*Block) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Block](d, m.DataTypes.Block, fn)
}
