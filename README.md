# go-slack-poster

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/slackbot?status.svg)](https://godoc.org/github.com/m0t0k1ch1/slackbot) [![wercker status](https://app.wercker.com/status/9b663f5536c8d7b8147b238613b336e3/s/master "wercker status")](https://app.wercker.com/project/bykey/9b663f5536c8d7b8147b238613b336e3)

Slack chat.postMessage API client for golang

## Example

``` go
package main

import (
	"log"

	slackposter "github.com/m0t0k1ch1/go-slack-poster"
)

func main() {
	client := slackposter.NewClient("your token")
	if err := client.SendMessage("#channel", "message"); err != nil {
		log.Fatal(err)
	}
}
```
