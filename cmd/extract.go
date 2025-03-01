package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/darkliquid/twoo/cmd/util"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

var extractTypes = []string{
	"account",
	"accountCreationIp",
	"accountLabel",
	"accountSuspension",
	"accountTimezone",
	"adengagements",
	"adImpressions",
	"adMobileConversionsUnattributed",
	"adMobileConversionsAttributed",
	"adOnlineConversionsUnattributed",
	"adOnlineConversionsAttributed",
	"ageInfo",
	"app",
	"block",
	"branchLinks",
	"catalogItem",
	"commerceCatalog",
	"communityNote",
	"communityNoteRating",
	"communityNoteTombstone",
	"communityTweet",
	"connectedApplication",
	"contact",
	"deletedNoteTweet",
	"deletedTweetHeaders",
	"deletedTweets",
	"deviceToken",
	"directMessagegroupheaders",
	"directMessageHeaders",
	"directMessageMute",
	"directMessages",
	"directMessagesGroup",
	"emailAddressChange",
	"follower",
	"following",
	"ipAudit",
	"keyRegistry",
	"like",
	"listsCreated",
	"listsMember",
	"listsSubscribed",
	"manifest",
	"moment",
	"mute",
	"niDevices",
	"noteTweet",
	"periscopeAccountInformation",
	"periscopeBanInformation",
	"periscopeBroadcastMetadata",
	"periscopeCommentsMadeByUser",
	"periscopeExpiredBroadcasts",
	"periscopeFollowers",
	"periscopeProfileDescription",
	"personalization",
	"phoneNumber",
	"productDrop",
	"productSet",
	"professionalData",
	"profile",
	"protectedHistory",
	"replyPrompt",
	"savedSearch",
	"screenNameChange",
	"shopModule",
	"shopifyAccount",
	"smartblock",
	"spacesMetadata",
	"sso",
	"tweetHeaders",
	"tweetdeck",
	"tweets",
	"twitterArticle",
	"twitterArticleMetadata",
	"twitterCircle",
	"twitterCircleMember",
	"twitterCircleTweet",
	"twitterShop",
	"userLinkClicks",
	"verified",
}

// dig taken from sprig library: github.com/Masterminds/sprig.
func dig(ps ...interface{}) (interface{}, error) {
	if len(ps) < 3 { //nolint:mnd // documented below
		return nil, errors.New("dig needs at least three arguments")
	}
	dict, ok := ps[len(ps)-1].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf(
			"dig needs a map as its last argument, got %T",
			ps[len(ps)-1],
		)
	}
	def := ps[len(ps)-2]
	ks := make([]string, len(ps)-2) //nolint:mnd // just stop it
	for i := 0; i < len(ks); i++ {
		ks[i], ok = ps[i].(string)
		if !ok {
			return nil, fmt.Errorf(
				"dig needs string arguments, got %T",
				ps[i],
			)
		}
	}

	return digFromDict(dict, def, ks)
}

// digFromDict taken from sprig library: github.com/Masterminds/sprig.
func digFromDict(dict map[string]interface{}, d interface{}, ks []string) (interface{}, error) {
	k, ns := ks[0], ks[1:]
	step, has := dict[k]
	if !has {
		return d, nil
	}
	if len(ns) == 0 {
		return step, nil
	}
	return digFromDict(step.(map[string]interface{}), d, ns) //nolint:errcheck // not my code
}

// toPrettyJSON taken from sprig library: github.com/Masterminds/sprig.
func toPrettyJSON(v interface{}) string {
	output, _ := json.MarshalIndent(v, "", "  ")
	return string(output)
}

var extractFuncMap = template.FuncMap{
	"get":  dig,
	"json": toPrettyJSON,
}

var extractFormat string

const defaultExtractFormat = "\n{{ . | json }}\n"

// extractCmd represents the extract command.
var extractCmd = &cobra.Command{
	Use:       "extract DATATYPE FILE|DIR",
	Short:     "extract data",
	Long:      `extract the data included in the archive`,
	Args:      util.SubcommandExactArgs(extractTypes, 1),
	ValidArgs: extractTypes,
	RunE: func(_ *cobra.Command, args []string) error {
		fs, closer, err := util.Open(args[1])
		if err != nil {
			return err
		}
		defer closer() //nolint:errcheck // nothing we can do about this

		tmpl, err := template.New("extract").Funcs(extractFuncMap).Parse(extractFormat)
		if err != nil {
			return err
		}

		data := twitwoo.New(fs)
		manifest, err := data.Manifest()
		if err != nil {
			return err
		}

		if args[0] == "manifest" {
			return tmpl.Execute(os.Stdout, manifest)
		}

		el := reflect.TypeOf(manifest.DataTypes)
		val := reflect.ValueOf(manifest.DataTypes)
		for i := 0; i < el.NumField(); i++ {
			field := el.Field(i)
			jTag := field.Tag.Get("json")
			if jTag == args[0] {
				dt, ok := val.Field(i).Interface().(twitwoo.DataType)
				if !ok {
					return fmt.Errorf("invalid datatype entry: %#v", val.Field(i).Interface())
				}
				return twitwoo.EachRaw(data, dt, func(m map[string]any) error {
					return tmpl.Execute(os.Stdout, m)
				})
			}
		}
		return fmt.Errorf("failed to find data type called %s", args[0])
	},
}

func init() {
	extractCmd.Flags().StringVarP(&extractFormat, "format", "f", defaultExtractFormat, "format of extracted data type")
	extractCmd.AddCommand(&cobra.Command{
		Use:  "datatypes",
		Long: extractCmd.UsageString() + "\nAvailable data types:\n  " + strings.Join(extractTypes, "\n  "),
		Args: cobra.NoArgs,
	})
	rootCmd.AddCommand(extractCmd)
}
