# m3u8-decorder

[![CircleCI](https://circleci.com/gh/poccariswet/m3u8-decoder/tree/master.svg?style=svg)](https://circleci.com/gh/poccariswet/m3u8-decoder/tree/master)

`m3u8-decoder` is a Go library base of rfc8216  

## Installation

```
$ go get github.com/poccariswet/m3u8-decorder
```

## Usage

``` go

func main() {
  master, err := os.Open("master.m3u8")
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  playlist, err := m3u8.DecodeFrom(bufio.NewReader(master)) // or ReadFile
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  fmt.Println(playlist)
}
```

## see

[pure go m3u8 decoder](https://medium.com/@poccariswet/pure-go-m3u8-decoder-eea5eb23c197)

## thx

- [master playlist](https://developer.apple.com/documentation/http_live_streaming/example_playlists_for_http_live_streaming/creating_a_master_playlist)
- [media](https://developer.apple.com/documentation/http_live_streaming/example_playlists_for_http_live_streaming/adding_alternate_media_to_a_playlist)
- [m3u](https://tools.ietf.org/html/draft-pantos-http-live-streaming-23)
