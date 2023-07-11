package twitwoo

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func registerProfileDecoders() {
	jsoniter.RegisterTypeDecoderFunc(
		"twitwoo.Profile",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			p := ((*Profile)(ptr))
			el := iter.ReadAny().Get("profile")
			el.Get("description").ToVal(&p.Description)
			p.Avatar = el.Get("avatarMediaUrl").ToString()
			p.Header = el.Get("headerMediaUrl").ToString()
		},
	)
}

// Profile is the structure of the data/profile.js file.
type Profile struct {
	Description struct {
		Bio      string `json:"bio"`
		Website  string `json:"website"`
		Location string `json:"location"`
	} `json:"description"`
	Avatar string `json:"avatarMediaUrl"`
	Header string `json:"headerMediaUrl"`
}

// Profiles returns all the Profile items.
func (d *Data) Profiles() ([]*Profile, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*Profile](d, m.DataTypes.Profile)
}

// EachProfile calls fn for each Profile item.
func (d *Data) EachProfile(fn func(*Profile) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*Profile](d, m.DataTypes.Profile, fn)
}
