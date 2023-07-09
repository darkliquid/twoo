package cmd

import (
	"encoding/json"
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
	"accountcreationip",
	"accountlabel",
	"accountsuspension",
	"accounttimezone",
	"adengagements",
	"adimpressions",
	"admobileconversionsunattributed",
	"admobileconversionsattributed",
	"adonlineconversionsunattributed",
	"adonlineconversionsattributed",
	"ageinfo",
	"app",
	"block",
	"branchlinks",
	"catalogitem",
	"commercecatalog",
	"communitynote",
	"communitynoterating",
	"communitynotetombstone",
	"communitytweet",
	"connectedapplication",
	"contact",
	"deletednotetweet",
	"deletedtweetheaders",
	"deletedtweets",
	"devicetoken",
	"directmessagegroupheaders",
	"directmessageheaders",
	"directmessagemute",
	"directmessages",
	"directmessagesgroup",
	"emailaddresschange",
	"follower",
	"following",
	"ipaudit",
	"keyregistry",
	"like",
	"listscreated",
	"listsmember",
	"listssubscribed",
	"manifest",
	"moment",
	"mute",
	"nidevices",
	"notetweet",
	"periscopeaccountinformation",
	"periscopebaninformation",
	"periscopebroadcastmetadata",
	"periscopecommentsmadebyuser",
	"periscopeexpiredbroadcasts",
	"periscopefollowers",
	"periscopeprofiledescription",
	"personalization",
	"phonenumber",
	"productdrop",
	"productset",
	"professionaldata",
	"profile",
	"protectedhistory",
	"replyprompt",
	"savedsearch",
	"screennamechange",
	"shopmodule",
	"shopifyaccount",
	"smartblock",
	"spacesmetadata",
	"sso",
	"tweetheaders",
	"tweetdeck",
	"tweets",
	"twitterarticle",
	"twitterarticlemetadata",
	"twittercircle",
	"twittercirclemember",
	"twittercircletweet",
	"twittershop",
	"userlinkclicks",
	"verified",
}

// dig taken from sprig library: github.com/Masterminds/sprig
func dig(ps ...interface{}) (interface{}, error) {
	if len(ps) < 3 {
		panic("dig needs at least three arguments")
	}
	dict := ps[len(ps)-1].(map[string]interface{})
	def := ps[len(ps)-2]
	ks := make([]string, len(ps)-2)
	for i := 0; i < len(ks); i++ {
		ks[i] = ps[i].(string)
	}

	return digFromDict(dict, def, ks)
}

// digFromDict taken from sprig library: github.com/Masterminds/sprig
func digFromDict(dict map[string]interface{}, d interface{}, ks []string) (interface{}, error) {
	k, ns := ks[0], ks[1:]
	step, has := dict[k]
	if !has {
		return d, nil
	}
	if len(ns) == 0 {
		return step, nil
	}
	return digFromDict(step.(map[string]interface{}), d, ns)
}

// toPrettyJson taken from sprig library: github.com/Masterminds/sprig
func toPrettyJson(v interface{}) string {
	output, _ := json.MarshalIndent(v, "", "  ")
	return string(output)
}

var extractFuncMap = template.FuncMap{
	"get":  dig,
	"json": toPrettyJson,
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
				return twitwoo.EachRaw(data, val.Field(i).Interface().(twitwoo.DataType), func(m map[string]any) error {
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
}

// Command returns the extract command.
func Command() *cobra.Command {
	return extractCmd
}
