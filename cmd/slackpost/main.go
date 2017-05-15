package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	slackposter "github.com/m0t0k1ch1/go-slack-poster"
)

const (
	DefaultTimeout = 5 // sec
)

func main() {
	os.Exit(run())
}

func run() int {
	var (
		token   string
		channel string
		message string
		timeout int
	)

	flag.StringVar(&token, "token", "", "token")
	flag.StringVar(&channel, "channel", "", "channel")
	flag.StringVar(&message, "message", "", "message")
	flag.IntVar(&timeout, "timeout", DefaultTimeout, "HTTP request timeout (sec)")
	flag.Parse()

	if len(token) == 0 {
		fmt.Println("no token")
		return 1
	}
	if len(message) == 0 {
		fmt.Println("no message")
		return 1
	}

	client := slackposter.NewClient(token)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	doneCh := make(chan bool, 1)
	errCh := make(chan error, 1)

	go func() {
		if err := client.SendMessage(ctx, channel, message); err != nil {
			errCh <- err
			return
		}
		doneCh <- true
	}()

	select {
	case <-doneCh:
		fmt.Println("success")
		return 0
	case err := <-errCh:
		fmt.Println(err)
		return 1
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return 1
	}
}
