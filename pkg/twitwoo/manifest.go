package twitwoo

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func registerManifestDecoders() {
	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.UserInfo",
		"AccountID",
		stringToInt64("decode account id"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ArchiveInfo",
		"SizeBytes",
		stringToInt64("decode size bytes"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ArchiveInfo",
		"MaxPartSizeBytes",
		stringToInt64("decode max part size bytes"),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.ArchiveInfo",
		"GenerationDate",
		stringToTime("decode generation date", time.RFC3339),
	)

	jsoniter.RegisterFieldDecoderFunc(
		"twitwoo.DataFile",
		"Count",
		stringToInt64("decode count"),
	)
}

// UserInfo is the structure of the user info section of the manifest.json file.
type UserInfo struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	AccountID   int64  `json:"accountId"`
}

// ArchiveInfo is the structure of the archive info section of the manifest.json file.
type ArchiveInfo struct {
	GenerationDate   time.Time `json:"generationDate"`
	SizeBytes        int64     `json:"sizeBytes"`
	MaxPartSizeBytes int64     `json:"maxPartSizeBytes"`
	IsPartialArchive bool      `json:"isPartialArchive"`
}

// DataFile is the structure of the data file sections of the manifest.json file.
type DataFile struct {
	Name     string `json:"fileName"`
	Preamble string `json:"globalName"`
	Count    int64  `json:"count"`
}

// DataType is the structure of the data type sections of the manifest.json file.
type DataType struct {
	MediaDir string     `json:"mediaDirectory,omitempty"`
	Files    []DataFile `json:"files"`
}

// Manifest is the structure of the manifest.json file.
type Manifest struct {
	DataTypes struct {
		Account                         DataType `json:"account"`
		AccountCreationIP               DataType `json:"accountCreationIp"`
		AccountLabel                    DataType `json:"accountLabel"`
		AccountSuspension               DataType `json:"accountSuspension"`
		AccountTimezone                 DataType `json:"accountTimezone"`
		AdEngagements                   DataType `json:"adEngagements"`
		AdImpressions                   DataType `json:"adImpressions"`
		AdMobileConversionsUnattributed DataType `json:"adMobileConversionsUnattributed"`
		AdMobileConversionsAttributed   DataType `json:"adMobileConversionsAttributed"`
		AdOnlineConversionsUnattributed DataType `json:"adOnlineConversionsUnattributed"`
		AdOnlineConversionsAttributed   DataType `json:"adOnlineConversionsAttributed"`
		AgeInfo                         DataType `json:"ageInfo"`
		App                             DataType `json:"app"`
		Block                           DataType `json:"block"`
		BranchLinks                     DataType `json:"branchLinks"`
		CatalogItem                     DataType `json:"catalogItem"`
		CommerceCatalog                 DataType `json:"commerceCatalog"`
		CommunityNote                   DataType `json:"communityNote"`
		CommunityNoteRating             DataType `json:"communityNoteRating"`
		CommunityNoteTombstone          DataType `json:"communityNoteTombstone"`
		CommunityTweet                  DataType `json:"communityTweet"`
		ConnectedApplication            DataType `json:"connectedApplication"`
		Contact                         DataType `json:"contact"`
		DeletedNoteTweet                DataType `json:"deletedNoteTweet"`
		DeletedTweetHeaders             DataType `json:"deletedTweetHeaders"`
		DeletedTweets                   DataType `json:"deletedTweets"`
		DeviceToken                     DataType `json:"deviceToken"`
		DirectMessageGroupHeaders       DataType `json:"directMessageGroupHeaders"`
		DirectMessageHeaders            DataType `json:"directMessageHeaders"`
		DirectMessageMute               DataType `json:"directMessageMute"`
		DirectMessages                  DataType `json:"directMessages"`
		DirectMessagesGroup             DataType `json:"directMessagesGroup"`
		EmailAddressChange              DataType `json:"emailAddressChange"`
		Follower                        DataType `json:"follower"`
		Following                       DataType `json:"following"`
		IPAudit                         DataType `json:"ipAudit"`
		KeyRegistry                     DataType `json:"keyRegistry"`
		Like                            DataType `json:"like"`
		ListsCreated                    DataType `json:"listsCreated"`
		ListsMember                     DataType `json:"listsMember"`
		ListsSubscribed                 DataType `json:"listsSubscribed"`
		Moment                          DataType `json:"moment"`
		Mute                            DataType `json:"mute"`
		NIDevices                       DataType `json:"niDevices"`
		NoteTweet                       DataType `json:"noteTweet"`
		PeriscopeAccountInformation     DataType `json:"periscopeAccountInformation"`
		PeriscopeBanInformation         DataType `json:"periscopeBanInformation"`
		PeriscopeBroadcastMetadata      DataType `json:"periscopeBroadcastMetadata"`
		PeriscopeCommentsMadeByUser     DataType `json:"periscopeCommentsMadeByUser"`
		PeriscopeExpiredBroadcasts      DataType `json:"periscopeExpiredBroadcasts"`
		PeriscopeFollowers              DataType `json:"periscopeFollowers"`
		PeriscopeProfileDescription     DataType `json:"periscopeProfileDescription"`
		Personalization                 DataType `json:"personalization"`
		PhoneNumber                     DataType `json:"phoneNumber"`
		ProductDrop                     DataType `json:"productDrop"`
		ProductSet                      DataType `json:"productSet"`
		ProfessionalData                DataType `json:"professionalData"`
		Profile                         DataType `json:"profile"`
		ProtectedHistory                DataType `json:"protectedHistory"`
		ReplyPrompt                     DataType `json:"replyPrompt"`
		SavedSearch                     DataType `json:"savedSearch"`
		ScreenNameChange                DataType `json:"screenNameChange"`
		ShopModule                      DataType `json:"shopModule"`
		ShopifyAccount                  DataType `json:"shopifyAccount"`
		Smartblock                      DataType `json:"smartblock"`
		SpacesMetadata                  DataType `json:"spacesMetadata"`
		SSO                             DataType `json:"sso"`
		TweetHeaders                    DataType `json:"tweetHeaders"`
		TweetDeck                       DataType `json:"tweetdeck"`
		Tweets                          DataType `json:"tweets"`
		TwitterArticle                  DataType `json:"twitterArticle"`
		TwitterArticleMetadata          DataType `json:"twitterArticleMetadata"`
		TwitterCircle                   DataType `json:"twitterCircle"`
		TwitterCircleMember             DataType `json:"twitterCircleMember"`
		TwitterCircleTweet              DataType `json:"twitterCircleTweet"`
		TwitterShop                     DataType `json:"twitterShop"`
		UserLinkClicks                  DataType `json:"userLinkClicks"`
		Verified                        DataType `json:"verified"`
	} `json:"dataTypes"`
	UserInfo    UserInfo    `json:"userInfo"`
	ArchiveInfo ArchiveInfo `json:"archiveInfo"`
}

// Manifest is the manifest.json file in the archive.
func (d *Data) Manifest() (*Manifest, error) {
	if d.manifestData != nil {
		return d.manifestData, nil
	}

	f, err := d.readDataFile(&DataFile{Name: "data/manifest.js", Preamble: "__THAR_CONFIG"})
	if err != nil {
		return nil, fmt.Errorf("failed to load manifest: %w", err)
	}
	defer f.Close()

	// Manifest is never going to be that large, so just load it all in.
	d.manifestData = &Manifest{}
	if err = jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(f).Decode(d.manifestData); err != nil {
		return nil, fmt.Errorf("failed to decode manifest: %w", err)
	}

	return d.manifestData, nil
}
