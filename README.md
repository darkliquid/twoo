# TWOO - TWitter Offline Online

A tool to take a twitter data archive and host it as a website (or turn it into static files).

## Commands

`twoo serve [--cache dir] [--bind host:port] archive|extracted_archive_dir`

The `serve` command spins up a webserver that hosts all your tweets from your archive.
The tweets are served in the order they are in the archive - **NOT** chronological
order.

`twoo extract [-f format_template] data_type archive|extracted_archive_dir`

The `extract` command returns data from the archive, formatted using the
`text/template` language to render the output. By default, it renders out
everything as JSON.

You can list the available datatypes with `twoo extract datatypes`

`twoo generate [-i -e -r -t -o -p -k -s -v -c] archive|extracted_archive_dir`

The `generate` command works like `serve`, but instead of hosting a site, it
builds it statically on disk. This also allows for more powerful options, like
sorting tweets chronologically or including replies or retweets.

Flags for `extract` include:

| Short Flags | Long Flags         | Description                                                           |
| ----------- | ------------------ | --------------------------------------------------------------------- |
| -e          | --extract-only     | only extract the tweets, don't build the static site                  |
| -h          | --help             | help for generate                                                     |
| -r          | --include-replies  | include replies in the output                                         |
| -t          | --include-retweets | include retweets in the output                                        |
| -o          | --out string       | where to write the static site to (default ".")                       |
| -p          | --page-size int    | how many tweets to include per page (default 20)                      |
| -k          | --skip-extract     | skip the extraction step and only build the static site               |
| -s          | --sort string      | sort order for tweets (asc or desc) (default "desc")                  |
| -v          | --verbose          | enable verbose output                                                 |
| -c          | --skip-cleanup     | skip cleaning up temporary json files                                 |
| -m          | --template-dir     | look in this directory for override templates                         |
| -i          | --search-index     | create a search index file and enable search in the default templates |

You can list the overridable templates with `twoo generate templates` and write them out to disk by also passing a directory.

Write the search index only creates an input file for an actual static search generation tool.
To generate the actual search javascript and compressed index file, use [tinysearch][1] like so:

```
tinysearch -p out -o out/search.json
```

Where `out` is the `generate` output dir.

`twoo completion`

You can use the completion command to generate shell completions.

## Templates

The generate command uses 4 templates to render the static site. They are:

-   `header.tmpl` - top half of every page
-   `footer.tmpl` - bottom half of every page
-   `tweet.tmpl` - template for each tweet
-   `stylesheet.tmpl` - template for the stylesheet

The templates all take the same `PageData` data except for `tweet.tmpl` which takes a `Tweet` object. The same functions are available to all templates.

### Data

#### PageData

```yaml
Profile:
    Description:
        Bio: "Description of profile"
        Website: "http://my.url"
        Location: "Someplace, Earth"
    Avatar: "http://path/to/twitter/avatar.jpg"
    Header: "http://path/to/twitter/header.jpg"
PrevPage: 1
NextPage: 3
UserInfo:
    UserName: cooluser
    DisplayName: The Coolest User
    AccountID: 7657865785
Page: 2
PageSize: 10
PageCount: 234
TweetCount: 2340
```

#### Tweet

```yaml
CreatedAt: "2023-04-01T12:34:56Z"
URLMap:
    "http://t.co/abc":
        DisplayURL: "my.link"
        ExpandedURL: "http://my.link/actual/link"
FullText: "complete tweet content"
Hashtags:
    - tag1
    - tag2
Mentions:
    - Name: @user
      ScreenName: A User
      ID: 858758765
Media:
    - ExpandedURL: "http://big/fat/url"
      URL: "http://small/url"
      MediaURL: "http://actual/media/location"
      Type: "video" | "photo" | "animated_gif"
      DisplayURL: "media/locatio..."
      ID: 96987698769
      SourceStatusID: 9869868976
InReplyToUserID: 878757865
InReplyToStatusID: 68587657865
ID: 5465376356
RetweetCount: 1
FavoriteCount: 3
```

### Functions

For examples of usage, see the built-in templates.

#### `fancy_tweet`

Accepts a `Tweet` object and renders it out nicely, including media.

#### `profile_header_url`

Accepts a `Profile` object and returns the url to the asset.

#### `profile_avatar_url`

Accepts a `Profile` object and returns the url to the asset.

#### `tweet_url`

Accepts a `Tweet` objects and returns the url to that specific status.

## TODO

-   [x] Enable specifying custom generation templates
-   [x] Extract default styles into stylesheet
-   [x] Add some kind of static JavaScript search index?
-   [ ] Add browse by hashtag/date

[1]: https://github.com/tinysearch/tinysearch
