package twitwoo_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

func TestTweets(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata")
	data := twitwoo.New(fs)

	tw, err := data.Tweets()
	require.NoError(t, err)
	require.Len(t, tw, 28)
	require.Equal(t, "1541752634348019715", tw[27].ID)
	require.Equal(t, "data/tweets_media/1541752634348019715-FWVnxzyWIAIEZU3.png", tw[27].Media[0].File())

	require.Equal(t, []string{"dungeon23", "City23", "setting23", "finishit23"}, tw[7].Hashtags)
	require.Equal(t, map[string]twitwoo.Link{
		"https://t.co/UeQo1oR1hd": {
			DisplayURL:  "dice.camp/@darkliquid",
			ExpandedURL: "http://dice.camp/@darkliquid",
		},
	}, tw[7].URLMap)

	require.Len(t, tw[11].Mentions, 12, "tweet: %#v", tw[11])
}

func TestEachTweet(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata")
	data := twitwoo.New(fs)

	var tw []*twitwoo.Tweet
	err := data.EachTweet(func(t *twitwoo.Tweet) error {
		tw = append(tw, t)
		return nil
	})

	require.NoError(t, err)
	require.Len(t, tw, 28)
	require.Equal(t, "1541752634348019715", tw[27].ID)

	require.Equal(t, []string{"dungeon23", "City23", "setting23", "finishit23"}, tw[7].Hashtags)
	require.Equal(t, map[string]twitwoo.Link{
		"https://t.co/UeQo1oR1hd": {
			DisplayURL:  "dice.camp/@darkliquid",
			ExpandedURL: "http://dice.camp/@darkliquid",
		},
	}, tw[7].URLMap)
}
