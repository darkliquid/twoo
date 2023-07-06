package twitwoo

import (
	jsoniter "github.com/json-iterator/go"
)

func registerTweetDeckDecoders() {
    /*
    // Example custom field decoder
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.TweetDeck",
		"FieldName",
		stringToInt64("decode field name"),
	)
    */
}

// TweetDeck is the structure of the data/tweet-deck.js file.
type TweetDeck struct {
    // Fields go here
}

func (t *TweetDeck) decode(el jsoniter.Any) {
	el = el.Get("tweetDeck")
	el.ToVal(t)
}

// TweetDecks returns all the TweetDeck items.
func (d *Data) TweetDecks() ([]*TweetDeck, error) {
	m, err := d.Manifest()
	if err != nil {
		return nil, err
	}
	return All[*TweetDeck](d, m.DataTypes.TweetDeck)
}

// EachTweetDeck calls fn for each TweetDeck item.
func (d *Data) EachTweetDeck(fn func(*TweetDeck) error) error {
	m, err := d.Manifest()
	if err != nil {
		return err
	}
	return Each[*TweetDeck](d, m.DataTypes.TweetDeck, fn)
}
