package slackbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://slack.com/api"
)

type response struct {
	Ok    bool
	Error string
}

type Client struct {
	token   string
	BaseURL string
}

func NewClient(token string) *Client {
	return &Client{
		token:   token,
		BaseURL: defaultBaseURL,
	}
}

func (c *Client) SendMessage(channel, message string) error {
	v := url.Values{}
	v.Set("token", c.token)
	v.Set("channel", channel)
	v.Set("text", message)
	v.Set("as_user", "true")

	res, err := http.PostForm(fmt.Sprintf("%s/chat.postMessage", c.BaseURL), v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody := &response{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(resBody); err != nil {
		return err
	}

	if !resBody.Ok {
		return errors.New(resBody.Error)
	}

	return nil
}
