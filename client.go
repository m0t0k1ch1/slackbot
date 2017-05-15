package slackposter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultUri    = "https://slack.com/api/chat.postMessage"
	defaultAsUser = true
)

type response struct {
	OK    bool
	Error string
}

type Client struct {
	uri    string
	token  string
	name   string
	asUser bool
}

func NewClient(token string) *Client {
	return &Client{
		uri:    defaultUri,
		token:  token,
		name:   "",
		asUser: defaultAsUser,
	}
}

func (client *Client) SetUri(uri string) {
	client.uri = uri
}

func (client *Client) SetToken(token string) {
	client.token = token
}

func (client *Client) SetName(name string) {
	client.name = name
}

func (client *Client) SetAsUser(b bool) {
	client.asUser = b
}

func (client *Client) SendMessage(channel, message string) error {
	v := url.Values{}
	v.Set("token", client.token)
	v.Set("channel", channel)
	v.Set("text", message)
	if client.asUser {
		v.Set("as_user", "true")
	} else {
		v.Set("username", client.name)
	}

	res, err := http.PostForm(client.uri, v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body := &response{}
	if err := json.NewDecoder(res.Body).Decode(body); err != nil {
		return err
	}
	if !body.OK {
		return fmt.Errorf(body.Error)
	}

	return nil
}
