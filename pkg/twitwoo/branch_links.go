package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerBranchLinksDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.BranchLinks",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// BranchLinks is the structure of the data/branch-links.js file.
type BranchLinks struct {
    // Fields go here
}

func (b *BranchLinks) decode(el jsoniter.Any) {
	el = el.Get("branchLinks")
	el.ToVal(b)
}

// BranchLinkss returns all the BranchLinks items.
func (d *Data) BranchLinkss() ([]*BranchLinks, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*BranchLinks](d, m.DataTypes.BranchLinks)
}

// EachBranchLinks calls fn for each BranchLinks item.
func (d *Data) EachBranchLinks(fn func(*BranchLinks) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*BranchLinks](d, m.DataTypes.BranchLinks, fn)
}
