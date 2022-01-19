## about

Static directory indexing on demand.

## setup

Fetch latest from GitHub:

```sh
go get github.com/thewhodidthis/listing
```

## usage

Pass in a list of files to index and optionally a title and/or a path to a custom template using the `-t` and `-x` flags respectively:
```sh
ls -A1 *.mp3 | listing -t Tracks > index.html
```
