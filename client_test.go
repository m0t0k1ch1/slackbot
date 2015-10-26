package slackbot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient() *Client {
	return NewClient("token")
}

func TestSetName(t *testing.T) {
	c := newTestClient()
	if c.name != "" {
		t.Errorf("default name is not \"\" - name: \"%s\"", c.name)
	}

	c.SetName("bot")
	if c.name != "bot" {
		t.Errorf("name is not \"bot\" - name: \"%s\"", c.name)
	}
}

func TestSetAsUser(t *testing.T) {
	c := newTestClient()
	if c.asUser {
		t.Errorf("default asUser is not false - asUser: %t", c.asUser)
	}

	c.SetAsUser(true)
	if !c.asUser {
		t.Errorf("asUser is not true - asUser: %t", c.asUser)
	}
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	res := &response{
		Ok: true,
	}
	body, _ := json.Marshal(res)
	w.Write(body)
}

func TestSendMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandler))
	defer server.Close()

	c := newTestClient()
	c.BaseURL = server.URL

	if err := c.SendMessage("#channel", "message"); err != nil {
		t.Error(err)
	}
}
