# go-slack-poster

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/go-slack-poster?status.svg)](https://godoc.org/github.com/m0t0k1ch1/go-slack-poster) [![wercker status](https://app.wercker.com/status/251a9f2059e5668a7d34f07808b2a06f/s/master "wercker status")](https://app.wercker.com/project/byKey/251a9f2059e5668a7d34f07808b2a06f)

Slack chat.postMessage API client for golang

## Examples

### Use as CLI

``` sh
$ go get -u github.com/m0t0k1ch1/go-slack-poster/cmd/slackpost
$ slackpost -token <token> -channel <channel> -text <text>
```

### Use in code

``` go
package main

import (
	"context"
	"log"

	slackposter "github.com/m0t0k1ch1/go-slack-poster"
)

func main() {
	client := slackposter.NewClient("xoxb-1234-56789abcdefghijklmnop")
	if err := client.SendText(context.Background(), "#channel", "message"); err != nil {
		log.Fatal(err)
	}
}
```
