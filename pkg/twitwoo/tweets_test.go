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

	require.Equal(t, []string{"dungeon23", "City23", "setting23", "finishit23"}, tw[7].Hashtags)
	require.Equal(t, map[string]string{
		"https://t.co/UeQo1oR1hd": "http://dice.camp/@darkliquid",
	}, tw[7].URLMap)
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
	require.Equal(t, map[string]string{
		"https://t.co/UeQo1oR1hd": "http://dice.camp/@darkliquid",
	}, tw[7].URLMap)
}
