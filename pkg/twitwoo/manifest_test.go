package twitwoo_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/darkliquid/twoo/pkg/twitwoo"
)

func TestManifest(t *testing.T) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), "./testdata")
	data := twitwoo.New(fs)

	m, err := data.Manifest()
	require.NoError(t, err)
	require.Equal(t, "darkliquid", m.UserInfo.UserName)
	require.Equal(t, twitwoo.DataType{
		MediaDir: "data/tweets_media",
		Files: []twitwoo.DataFile{
			{
				Name:     "data/tweets.js",
				Preamble: "YTD.tweets.part0",
				Count:    8769,
			},
		},
	}, m.DataTypes.Tweets)
}
