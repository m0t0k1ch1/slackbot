package slackposter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Uri    = "https://slack.com/api/chat.postMessage"
	AsUser = true
)

type request struct {
	Channel     string      `json:"channel"`
	Text        string      `json:"text"`
	Attachments interface{} `json:"attachments"`
	AsUser      bool        `json:"as_user"`
}

type response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type Config struct {
	Uri   string
	Token string
}

type Client struct {
	*http.Client
	config *Config
}

func NewClient(token string) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &Config{
			Uri:   Uri,
			Token: token,
		},
	}
}

func (client *Client) SetUri(uri string) {
	client.config.Uri = uri
}

func (client *Client) SetToken(token string) {
	client.config.Token = token
}

func (client *Client) SendMessage(ctx context.Context, channel, text string, attachments interface{}) error {
	req := &request{
		Channel:     channel,
		Text:        text,
		Attachments: attachments,
		AsUser:      AsUser,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	return client.doApi(ctx, b)
}

func (client *Client) doApi(ctx context.Context, b []byte) error {
	req, err := http.NewRequest(http.MethodPost, client.config.Uri, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+client.config.Token)
	req.Header.Set("Content-type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}
	if !res.Ok {
		return fmt.Errorf(res.Error)
	}

	return nil
}
