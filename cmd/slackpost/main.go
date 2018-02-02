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
		text    string
		timeout int
	)

	flag.StringVar(&token, "token", "", "token")
	flag.StringVar(&channel, "channel", "", "channel")
	flag.StringVar(&text, "text", "", "text")
	flag.IntVar(&timeout, "timeout", DefaultTimeout, "HTTP request timeout (sec)")
	flag.Parse()

	if len(token) == 0 {
		return fail(fmt.Errorf("no token"))
	}
	if len(channel) == 0 {
		return fail(fmt.Errorf("no channel"))
	}
	if len(text) == 0 {
		return fail(fmt.Errorf("no text"))
	}

	client := slackposter.NewClient(token)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		errCh <- client.SendMessage(ctx, channel, text, nil)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return fail(err)
		}
		return success()
	case <-ctx.Done():
		return fail(ctx.Err())
	}
}

func success() int {
	return 0
}

func fail(err error) int {
	printlnStdout(err)
	return 1
}

func printlnStdout(v ...interface{}) {
	fmt.Fprintln(os.Stdout, v...)
}
