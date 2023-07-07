package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

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

func (p *Profile) decode(el jsoniter.Any) {
	el = el.Get("profile")
	el.ToVal(p)
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
