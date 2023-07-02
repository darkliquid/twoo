package twitwoo

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestTweets(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata")
	data := New(fs)

	tw, err := data.Tweets()
	require.NoError(t, err)
	require.Len(t, tw, 28)
	require.Equal(t, "1541752634348019715", tw[27].ID)

	require.Equal(t, []string{"dungeon23", "City23", "setting23", "finishit23"}, tw[7].Hashtags)
	require.Equal(t, map[string]string{
		"https://t.co/UeQo1oR1hd": "http://dice.camp/@darkliquid",
	}, tw[7].URLMap)
}
