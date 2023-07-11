package twitwoo_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
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
	require.Equal(t, int64(1541752634348019715), tw[27].ID)
	require.Equal(t, "http://pbs.twimg.com/media/FWVnxzyWIAIEZU3.png", tw[27].Media[0].MediaURL)

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
		if len(tw) >= 28 {
			return twitwoo.ErrBreak
		}
		return nil
	})

	require.NoError(t, err)
	require.Len(t, tw, 28)
	require.Equal(t, int64(1541752634348019715), tw[27].ID)

	require.Equal(t, []string{"dungeon23", "City23", "setting23", "finishit23"}, tw[7].Hashtags)
	require.Equal(t, map[string]twitwoo.Link{
		"https://t.co/UeQo1oR1hd": {
			DisplayURL:  "dice.camp/@darkliquid",
			ExpandedURL: "http://dice.camp/@darkliquid",
		},
	}, tw[7].URLMap)
}

const exampleTweet = `{ "tweet":{
      "edit_info" : {
        "initial" : {
          "editTweetIds" : [
            "585456172267917312"
          ],
          "editableUntil" : "2015-04-07T15:56:43.375Z",
          "editsRemaining" : "5",
          "isEditEligible" : true
        }
      },
      "retweeted" : false,
      "source" : "<a href=\"https://about.twitter.com/products/tweetdeck\" rel=\"nofollow\">TweetDeck</a>",
      "entities" : {
        "user_mentions" : [
          {
            "name" : "Andrew Nesbitt",
            "screen_name" : "teabass",
            "indices" : [
              "3",
              "11"
            ],
            "id_str" : "686803",
            "id" : "686803"
          }
        ],
        "urls" : [ ],
        "symbols" : [ ],
        "media" : [
          {
            "expanded_url" : "https://twitter.com/teabass/status/585436639423569921/photo/1",
            "source_status_id" : "585436639423569921",
            "indices" : [
              "40",
              "62"
            ],
            "url" : "http://t.co/6sCKGymthV",
            "media_url" : "http://pbs.twimg.com/tweet_video_thumb/CB_jhR-WAAAECxl.png",
            "id_str" : "585436637481533440",
            "source_user_id" : "686803",
            "id" : "585436637481533440",
            "media_url_https" : "https://pbs.twimg.com/tweet_video_thumb/CB_jhR-WAAAECxl.png",
            "source_user_id_str" : "686803",
            "sizes" : {
              "small" : {
                "w" : "300",
                "h" : "225",
                "resize" : "fit"
              },
              "medium" : {
                "w" : "300",
                "h" : "225",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "large" : {
                "w" : "300",
                "h" : "225",
                "resize" : "fit"
              }
            },
            "type" : "photo",
            "source_status_id_str" : "585436639423569921",
            "display_url" : "pic.twitter.com/6sCKGymthV"
          }
        ],
        "hashtags" : [ ]
      },
      "display_text_range" : [
        "0",
        "62"
      ],
      "favorite_count" : "0",
      "id_str" : "585456172267917312",
      "truncated" : false,
      "retweet_count" : "0",
      "id" : "585456172267917312",
      "possibly_sensitive" : false,
      "created_at" : "Tue Apr 07 14:56:43 +0000 2015",
      "favorited" : false,
      "full_text" : "RT @teabass: Refactoring in a nutshell: http://t.co/6sCKGymthV",
      "lang" : "en",
      "extended_entities" : {
        "media" : [
          {
            "expanded_url" : "https://twitter.com/teabass/status/585436639423569921/photo/1",
            "source_status_id" : "585436639423569921",
            "indices" : [
              "40",
              "62"
            ],
            "url" : "http://t.co/6sCKGymthV",
            "media_url" : "http://pbs.twimg.com/tweet_video_thumb/CB_jhR-WAAAECxl.png",
            "id_str" : "585436637481533440",
            "video_info" : {
              "aspect_ratio" : [
                "4",
                "3"
              ],
              "variants" : [
                {
                  "bitrate" : "0",
                  "content_type" : "video/mp4",
                  "url" : "https://video.twimg.com/tweet_video/CB_jhR-WAAAECxl.mp4"
                }
              ]
            },
            "source_user_id" : "686803",
            "id" : "585436637481533440",
            "media_url_https" : "https://pbs.twimg.com/tweet_video_thumb/CB_jhR-WAAAECxl.png",
            "source_user_id_str" : "686803",
            "sizes" : {
              "small" : {
                "w" : "300",
                "h" : "225",
                "resize" : "fit"
              },
              "medium" : {
                "w" : "300",
                "h" : "225",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "large" : {
                "w" : "300",
                "h" : "225",
                "resize" : "fit"
              }
            },
            "type" : "animated_gif",
            "source_status_id_str" : "585436639423569921",
            "display_url" : "pic.twitter.com/6sCKGymthV"
          }
        ]
      }
    }
}`

//nolint:lll  // This JSON is gona be long
const exampleTweet2 = `{
    "tweet" : {
      "edit_info" : {
        "initial" : {
          "editTweetIds" : [
            "1444613894736130055"
          ],
          "editableUntil" : "2021-10-03T11:42:46.790Z",
          "editsRemaining" : "5",
          "isEditEligible" : true
        }
      },
      "retweeted" : false,
      "source" : "<a href=\"http://twitter.com/download/android\" rel=\"nofollow\">Twitter for Android</a>",
      "entities" : {
        "user_mentions" : [
          {
            "name" : "Secret Attic",
            "screen_name" : "secretattic2020",
            "indices" : [
              "3",
              "19"
            ],
            "id_str" : "1627395997688074241",
            "id" : "1627395997688074241"
          }
        ],
        "urls" : [
          {
            "url" : "https://t.co/Dn5fStuX9f",
            "expanded_url" : "https://www.secret-attic.co.uk/p/short-story-contest.html",
            "display_url" : "secret-attic.co.uk/p/short-story-…",
            "indices" : [
              "81",
              "104"
            ]
          }
        ],
        "symbols" : [ ],
        "media" : [
          {
            "expanded_url" : "https://twitter.com/secretattic2020/status/1443876132160524310/photo/1",
            "source_status_id" : "1443876132160524310",
            "indices" : [
              "105",
              "128"
            ],
            "url" : "https://t.co/Mkq93EVirc",
            "media_url" : "http://pbs.twimg.com/tweet_video_thumb/FAmtr4VX0AAeJyJ.jpg",
            "id_str" : "1443876124916961280",
            "source_user_id" : "103290954",
            "id" : "1443876124916961280",
            "media_url_https" : "https://pbs.twimg.com/tweet_video_thumb/FAmtr4VX0AAeJyJ.jpg",
            "source_user_id_str" : "103290954",
            "sizes" : {
              "small" : {
                "w" : "480",
                "h" : "322",
                "resize" : "fit"
              },
              "large" : {
                "w" : "480",
                "h" : "322",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "medium" : {
                "w" : "480",
                "h" : "322",
                "resize" : "fit"
              }
            },
            "type" : "photo",
            "source_status_id_str" : "1443876132160524310",
            "display_url" : "pic.twitter.com/Mkq93EVirc"
          }
        ],
        "hashtags" : [ ]
      },
      "display_text_range" : [
        "0",
        "128"
      ],
      "favorite_count" : "0",
      "id_str" : "1444613894736130055",
      "truncated" : false,
      "retweet_count" : "0",
      "id" : "1444613894736130055",
      "possibly_sensitive" : false,
      "created_at" : "Sun Oct 03 10:42:46 +0000 2021",
      "favorited" : false,
      "full_text" : "RT @secretattic2020: The Monthly Short Story contest for October is now open...\n\nhttps://t.co/Dn5fStuX9f https://t.co/Mkq93EVirc",
      "lang" : "en",
      "extended_entities" : {
        "media" : [
          {
            "expanded_url" : "https://twitter.com/secretattic2020/status/1443876132160524310/photo/1",
            "source_status_id" : "1443876132160524310",
            "indices" : [
              "105",
              "128"
            ],
            "url" : "https://t.co/Mkq93EVirc",
            "media_url" : "http://pbs.twimg.com/tweet_video_thumb/FAmtr4VX0AAeJyJ.jpg",
            "id_str" : "1443876124916961280",
            "video_info" : {
              "aspect_ratio" : [
                "240",
                "161"
              ],
              "variants" : [
                {
                  "bitrate" : "0",
                  "content_type" : "video/mp4",
                  "url" : "https://video.twimg.com/tweet_video/FAmtr4VX0AAeJyJ.mp4"
                }
              ]
            },
            "source_user_id" : "103290954",
            "id" : "1443876124916961280",
            "media_url_https" : "https://pbs.twimg.com/tweet_video_thumb/FAmtr4VX0AAeJyJ.jpg",
            "source_user_id_str" : "103290954",
            "sizes" : {
              "small" : {
                "w" : "480",
                "h" : "322",
                "resize" : "fit"
              },
              "large" : {
                "w" : "480",
                "h" : "322",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "medium" : {
                "w" : "480",
                "h" : "322",
                "resize" : "fit"
              }
            },
            "type" : "animated_gif",
            "source_status_id_str" : "1443876132160524310",
            "display_url" : "pic.twitter.com/Mkq93EVirc"
          }
        ]
      }
    }
  }`

//nolint:lll  // This JSON is gona be long
const exampleTweet3 = `{
    "tweet" : {
      "edit_info" : {
        "initial" : {
          "editTweetIds" : [
            "1292062131647635456"
          ],
          "editableUntil" : "2020-08-08T12:36:32.195Z",
          "editsRemaining" : "5",
          "isEditEligible" : true
        }
      },
      "retweeted" : false,
      "source" : "<a href=\"http://twitter.com/download/android\" rel=\"nofollow\">Twitter for Android</a>",
      "entities" : {
        "user_mentions" : [
          {
            "name" : "Simpsons Against DevOps",
            "screen_name" : "SimpsonsOps",
            "indices" : [
              "3",
              "15"
            ],
            "id_str" : "1279015319701315584",
            "id" : "1279015319701315584"
          }
        ],
        "urls" : [ ],
        "symbols" : [ ],
        "media" : [
          {
            "expanded_url" : "https://twitter.com/SimpsonsOps/status/1291924030757277697/video/1",
            "source_status_id" : "1291924030757277697",
            "indices" : [
              "53",
              "76"
            ],
            "url" : "https://t.co/q2usg7ygsh",
            "media_url" : "http://pbs.twimg.com/ext_tw_video_thumb/1291923825303490562/pu/img/6e65Wacwl9cB1IDW.jpg",
            "id_str" : "1291923825303490562",
            "source_user_id" : "1279015319701315584",
            "id" : "1291923825303490562",
            "media_url_https" : "https://pbs.twimg.com/ext_tw_video_thumb/1291923825303490562/pu/img/6e65Wacwl9cB1IDW.jpg",
            "source_user_id_str" : "1279015319701315584",
            "sizes" : {
              "small" : {
                "w" : "680",
                "h" : "510",
                "resize" : "fit"
              },
              "medium" : {
                "w" : "986",
                "h" : "740",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "large" : {
                "w" : "986",
                "h" : "740",
                "resize" : "fit"
              }
            },
            "type" : "photo",
            "source_status_id_str" : "1291924030757277697",
            "display_url" : "pic.twitter.com/q2usg7ygsh"
          }
        ],
        "hashtags" : [ ]
      },
      "display_text_range" : [
        "0",
        "76"
      ],
      "favorite_count" : "0",
      "id_str" : "1292062131647635456",
      "truncated" : false,
      "retweet_count" : "0",
      "id" : "1292062131647635456",
      "possibly_sensitive" : false,
      "created_at" : "Sat Aug 08 11:36:32 +0000 2020",
      "favorited" : false,
      "full_text" : "RT @SimpsonsOps: Every ops team trying to \"do agile\" https://t.co/q2usg7ygsh",
      "lang" : "en",
      "extended_entities" : {
        "media" : [
          {
            "expanded_url" : "https://twitter.com/SimpsonsOps/status/1291924030757277697/video/1",
            "source_status_id" : "1291924030757277697",
            "indices" : [
              "53",
              "76"
            ],
            "url" : "https://t.co/q2usg7ygsh",
            "media_url" : "http://pbs.twimg.com/ext_tw_video_thumb/1291923825303490562/pu/img/6e65Wacwl9cB1IDW.jpg",
            "id_str" : "1291923825303490562",
            "video_info" : {
              "aspect_ratio" : [
                "493",
                "370"
              ],
              "duration_millis" : "24800",
              "variants" : [
                {
                  "bitrate" : "832000",
                  "content_type" : "video/mp4",
                  "url" : "https://video.twimg.com/ext_tw_video/1291923825303490562/pu/vid/478x360/5Ru_huxGQzpgBf6E.mp4?tag=10"
                },
                {
                  "bitrate" : "256000",
                  "content_type" : "video/mp4",
                  "url" : "https://video.twimg.com/ext_tw_video/1291923825303490562/pu/vid/358x270/EPOyedoxJ0m-pzGA.mp4?tag=10"
                },
                {
                  "bitrate" : "2176000",
                  "content_type" : "video/mp4",
                  "url" : "https://video.twimg.com/ext_tw_video/1291923825303490562/pu/vid/958x720/Woa8cTl6giO9lfBd.mp4?tag=10"
                },
                {
                  "content_type" : "application/x-mpegURL",
                  "url" : "https://video.twimg.com/ext_tw_video/1291923825303490562/pu/pl/HiFLDK4tEIfoLPMA.m3u8?tag=10"
                }
              ]
            },
            "source_user_id" : "1279015319701315584",
            "additional_media_info" : {
              "monetizable" : false
            },
            "id" : "1291923825303490562",
            "media_url_https" : "https://pbs.twimg.com/ext_tw_video_thumb/1291923825303490562/pu/img/6e65Wacwl9cB1IDW.jpg",
            "source_user_id_str" : "1279015319701315584",
            "sizes" : {
              "small" : {
                "w" : "680",
                "h" : "510",
                "resize" : "fit"
              },
              "medium" : {
                "w" : "986",
                "h" : "740",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "large" : {
                "w" : "986",
                "h" : "740",
                "resize" : "fit"
              }
            },
            "type" : "video",
            "source_status_id_str" : "1291924030757277697",
            "display_url" : "pic.twitter.com/q2usg7ygsh"
          }
        ]
      }
    }
  }`

//nolint:lll  // This JSON is gona be long
const exampleTweet4 = `{
    "tweet" : {
      "edit_info" : {
        "initial" : {
          "editTweetIds" : [
            "231600589707038720"
          ],
          "editableUntil" : "2012-08-04T05:00:58.174Z",
          "editsRemaining" : "5",
          "isEditEligible" : true
        }
      },
      "retweeted" : false,
      "source" : "<a href=\"https://about.twitter.com/products/tweetdeck\" rel=\"nofollow\">TweetDeck</a>",
      "entities" : {
        "user_mentions" : [
          {
            "name" : "SimCity",
            "screen_name" : "simcity",
            "indices" : [
              "3",
              "11"
            ],
            "id_str" : "414379328",
            "id" : "414379328"
          }
        ],
        "urls" : [ ],
        "symbols" : [ ],
        "media" : [
          {
            "expanded_url" : "https://twitter.com/simcity/status/231519535277154304/photo/1",
            "source_status_id" : "231519535277154304",
            "indices" : [
              "102",
              "122"
            ],
            "url" : "http://t.co/Clu3b6g6",
            "media_url" : "http://pbs.twimg.com/media/AzaFym5CIAARZgp.jpg",
            "id_str" : "231519535335874560",
            "source_user_id" : "414379328",
            "id" : "231519535335874560",
            "media_url_https" : "https://pbs.twimg.com/media/AzaFym5CIAARZgp.jpg",
            "source_user_id_str" : "414379328",
            "sizes" : {
              "large" : {
                "w" : "1024",
                "h" : "765",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "medium" : {
                "w" : "1024",
                "h" : "765",
                "resize" : "fit"
              },
              "small" : {
                "w" : "680",
                "h" : "508",
                "resize" : "fit"
              }
            },
            "type" : "photo",
            "source_status_id_str" : "231519535277154304",
            "display_url" : "pic.twitter.com/Clu3b6g6"
          }
        ],
        "hashtags" : [
          {
            "text" : "SimCity",
            "indices" : [
              "86",
              "94"
            ]
          }
        ]
      },
      "display_text_range" : [
        "0",
        "122"
      ],
      "favorite_count" : "0",
      "id_str" : "231600589707038720",
      "truncated" : false,
      "retweet_count" : "0",
      "id" : "231600589707038720",
      "possibly_sensitive" : false,
      "created_at" : "Sat Aug 04 04:00:58 +0000 2012",
      "favorited" : false,
      "full_text" : "RT @simcity: Feeling lucky? Follow us and RT for a random chance to score this snazzy #SimCity shirt! http://t.co/Clu3b6g6",
      "lang" : "en",
      "extended_entities" : {
        "media" : [
          {
            "expanded_url" : "https://twitter.com/simcity/status/231519535277154304/photo/1",
            "source_status_id" : "231519535277154304",
            "indices" : [
              "102",
              "122"
            ],
            "url" : "http://t.co/Clu3b6g6",
            "media_url" : "http://pbs.twimg.com/media/AzaFym5CIAARZgp.jpg",
            "id_str" : "231519535335874560",
            "source_user_id" : "414379328",
            "id" : "231519535335874560",
            "media_url_https" : "https://pbs.twimg.com/media/AzaFym5CIAARZgp.jpg",
            "source_user_id_str" : "414379328",
            "sizes" : {
              "large" : {
                "w" : "1024",
                "h" : "765",
                "resize" : "fit"
              },
              "thumb" : {
                "w" : "150",
                "h" : "150",
                "resize" : "crop"
              },
              "medium" : {
                "w" : "1024",
                "h" : "765",
                "resize" : "fit"
              },
              "small" : {
                "w" : "680",
                "h" : "508",
                "resize" : "fit"
              }
            },
            "type" : "photo",
            "source_status_id_str" : "231519535277154304",
            "display_url" : "pic.twitter.com/Clu3b6g6"
          }
        ]
      }
    }
  }`

func TestTweetDecoding(t *testing.T) {
	// This is a tweet with a single video.
	var tweet twitwoo.Tweet
	require.NoError(t, jsoniter.ConfigDefault.Unmarshal([]byte(exampleTweet), &tweet))
	require.Equal(t, int64(585456172267917312), tweet.ID)
	require.Len(t, tweet.Media, 1)
	require.Equal(t,
		"https://video.twimg.com/tweet_video/CB_jhR-WAAAECxl.mp4",
		tweet.Media[0].MediaURL,
	)

	// We should use a video variant if it exists, as that is what is in the
	// archive.
	var tweet2 twitwoo.Tweet
	require.NoError(t, jsoniter.ConfigDefault.Unmarshal([]byte(exampleTweet2), &tweet2))
	require.Equal(t, int64(1444613894736130055), tweet2.ID)
	require.Len(t, tweet2.Media, 1)
	require.Equal(t,
		"https://video.twimg.com/tweet_video/FAmtr4VX0AAeJyJ.mp4",
		tweet2.Media[0].MediaURL,
	)

	// The media url should be for the highest quality variant as that's what
	// the archive contains.
	var tweet3 twitwoo.Tweet
	require.NoError(t, jsoniter.ConfigDefault.Unmarshal([]byte(exampleTweet3), &tweet3))
	require.Equal(t, int64(1292062131647635456), tweet3.ID)
	require.Len(t, tweet3.Media, 1)
	require.Equal(t,
		"https://video.twimg.com/ext_tw_video/1291923825303490562/pu/vid/958x720/Woa8cTl6giO9lfBd.mp4",
		tweet3.Media[0].MediaURL,
	)

	// The archive doesn't include media for retweets, but we still have the url in case
	// we want to try fetching it.
	var tweet4 twitwoo.Tweet
	require.NoError(t, jsoniter.ConfigDefault.Unmarshal([]byte(exampleTweet4), &tweet4))
	require.Equal(t, int64(231600589707038720), tweet4.ID)
	require.Len(t, tweet4.Media, 1)
	require.Equal(t,
		"http://pbs.twimg.com/media/AzaFym5CIAARZgp.jpg",
		tweet4.Media[0].MediaURL,
	)
	require.NotEqual(t, tweet4.ID, tweet4.Media[0].SourceStatusID)
}
