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

`twoo generate [i-e -r -t -o -p -k -s -v -c] archive|extracted_archive_dir`

The `generate` command works like `serve`, but instead of hosting a site, it
builds it statically on disk. This also allows for more powerful options, like
sorting tweets chronologically or including replies or retweets.

Flags for `extract` include:

  | Short Flags | Long Flags         | Description                                             |
  | ----------- | ------------------ | ------------------------------------------------------- |
  | -e          | --extract-only     | only extract the tweets, don't build the static site    |
  | -h          | --help             | help for generate                                       |
  | -r          | --include-replies  | include replies in the output                           |
  | -t          | --include-retweets | include retweets in the output                          |
  | -o          | --out string       | where to write the static site to (default ".")         |
  | -p          | --page-size int    | how many tweets to include per page (default 20)        |
  | -k          | --skip-extract     | skip the extraction step and only build the static site |
  | -s          | --sort string      | sort order for tweets (asc or desc) (default "desc")    |
  | -v          | --verbose          | enable verbose output                                   |
  | -c          | --skip-cleanup     | skip cleaning up temporary json files                   |
  | -m          | --template-dir     | look in this directory for override templates           |

You can list the overridable templates with `twoo generate templates`

`twoo completion`

You can use the completion command to generate shell completions.

## TODO

 - [x] Enable specifying custom generation templates
 - [x] Extract default styles into stylesheet
 - [ ] Add some kind of static JavaScript search index?
 - [ ] Add browse by hashtag/date
