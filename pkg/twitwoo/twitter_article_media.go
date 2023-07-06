package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterArticleMediaDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterArticleMedia",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterArticleMedia is the structure of the data/twitter-article-media.js file.
type TwitterArticleMedia struct {
    // Fields go here
}

func (t *TwitterArticleMedia) decode(el jsoniter.Any) {
	el = el.Get("twitterArticleMedia")
	el.ToVal(t)
}

// TwitterArticleMedias returns all the TwitterArticleMedia items.
func (d *Data) TwitterArticleMedias() ([]*TwitterArticleMedia, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterArticleMedia](d, m.DataTypes.TwitterArticleMedia)
}

// EachTwitterArticleMedia calls fn for each TwitterArticleMedia item.
func (d *Data) EachTwitterArticleMedia(fn func(*TwitterArticleMedia) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterArticleMedia](d, m.DataTypes.TwitterArticleMedia, fn)
}
