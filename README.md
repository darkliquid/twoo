# TWOO - TWitter Offline Online

A tool to take a twitter data archive and host it as a website (or turn it into static files).

## Commands

`twoo serve [--cache dir] [--bind host:port] archive|extracted archive dir`

The `serve` command spins up a webserver that hosts all your tweets from your archive.
The tweets are served in the order they are in the archive - **NOT** chronological
order.

`twoo extract [-f format_template] {data type}`

The `extract` command returns data from the archive, formatted using the
`text/template` language to render the output.
