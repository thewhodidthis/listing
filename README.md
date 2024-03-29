## about

Helps create static directory listings in HTML by default.

## setup

Fetch latest from GitHub:

```
go get github.com/thewhodidthis/listing
```

## usage

Pass in a list of files to index and optionally a title and/or a path to a custom template using the `-t` and `-x` flags respectively:

```
# Override default title
ls -A1 *.mp3 | listing -t Tracks > index.html
```

Generate an Atom feed to match:

```
# Override default template
ls -d */ | listing -x feed.xml.tmpl > feed.xml
```

Given a feed.xml.tmpl along the lines of:

```html
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{.Title}}</title>
    <link>https://foo.bar/</link>
    <description>About the foo bar listing feed</description>
    <language>en</language>
    <atom:link href="https://foo.bar/feed.xml" rel="self" type="application/rss+xml"/>
    {{- range .List}}
    {{- $url := printf "https://foo.bar/%s" .Name}}
    <item>
      <title>{{.Name}}</title>
      <pubDate>{{.ModTime.Format "Mon, 02 Jan 2006 03:04:05 -0700"}}</pubDate>
      <link>{{$url}}</link>
      <guid>{{$url}}</guid>
    </item>
    {{- end}}
  </channel>
</rss>
```
