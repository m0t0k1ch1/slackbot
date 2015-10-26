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
	name    string
	asUser  bool
	BaseURL string
}

func NewClient(token string) *Client {
	return &Client{
		token:   token,
		name:    "",
		asUser:  false,
		BaseURL: defaultBaseURL,
	}
}

func (c *Client) SetName(name string) {
	c.name = name
}

func (c *Client) SetAsUser(b bool) {
	c.asUser = b
}

func (c *Client) SendMessage(channel, message string) error {
	v := url.Values{}
	v.Set("token", c.token)
	v.Set("channel", channel)
	v.Set("text", message)
	if c.asUser {
		v.Set("as_user", "true")
	} else {
		v.Set("username", c.name)
	}

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
