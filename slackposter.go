package slackposter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultUri    = "https://slack.com/api/chat.postMessage"
	defaultAsUser = true
)

type response struct {
	OK    bool
	Error string
}

type Config struct {
	Uri    string
	Token  string
	Name   string
	AsUser bool
}

type Client struct {
	*http.Client
	config *Config
}

func NewClient(token string) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &Config{
			Uri:    defaultUri,
			Token:  token,
			Name:   "",
			AsUser: defaultAsUser,
		},
	}
}

func (client *Client) SetUri(uri string) {
	client.config.Uri = uri
}

func (client *Client) SetToken(token string) {
	client.config.Token = token
}

func (client *Client) SetName(name string) {
	client.config.Name = name
}

func (client *Client) SetAsUser(b bool) {
	client.config.AsUser = b
}

func (client *Client) SendMessage(ctx context.Context, channel, message string) error {
	v := url.Values{}
	v.Add("token", client.config.Token)
	v.Add("channel", channel)
	v.Add("text", message)
	if client.config.AsUser {
		v.Add("as_user", "true")
	} else {
		v.Add("username", client.config.Name)
	}

	req, err := http.NewRequest(http.MethodPost, client.config.Uri, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res := &response{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return err
	}
	if !res.OK {
		return fmt.Errorf(res.Error)
	}

	return nil
}
