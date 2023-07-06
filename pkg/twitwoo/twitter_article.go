package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterArticleDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterArticle",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterArticle is the structure of the data/twitter-article.js file.
type TwitterArticle struct {
    // Fields go here
}

func (t *TwitterArticle) decode(el jsoniter.Any) {
	el = el.Get("twitterArticle")
	el.ToVal(t)
}

// TwitterArticles returns all the TwitterArticle items.
func (d *Data) TwitterArticles() ([]*TwitterArticle, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterArticle](d, m.DataTypes.TwitterArticle)
}

// EachTwitterArticle calls fn for each TwitterArticle item.
func (d *Data) EachTwitterArticle(fn func(*TwitterArticle) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterArticle](d, m.DataTypes.TwitterArticle, fn)
}
