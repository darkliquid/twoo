package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTwitterArticleMetadataDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TwitterArticleMetadata",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TwitterArticleMetadata is the structure of the data/twitter-article-metadata.js file.
type TwitterArticleMetadata struct {
    // Fields go here
}

func (t *TwitterArticleMetadata) decode(el jsoniter.Any) {
	el = el.Get("twitterArticleMetadata")
	el.ToVal(t)
}

// TwitterArticleMetadatas returns all the TwitterArticleMetadata items.
func (d *Data) TwitterArticleMetadatas() ([]*TwitterArticleMetadata, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TwitterArticleMetadata](d, m.DataTypes.TwitterArticleMetadata)
}

// EachTwitterArticleMetadata calls fn for each TwitterArticleMetadata item.
func (d *Data) EachTwitterArticleMetadata(fn func(*TwitterArticleMetadata) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TwitterArticleMetadata](d, m.DataTypes.TwitterArticleMetadata, fn)
}
