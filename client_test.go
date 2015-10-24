package slackbot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testClient *Client
)

func init() {
	testClient = NewClient("token")
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	res := &response{
		Ok: true,
	}
	body, _ := json.Marshal(res)
	rw.Write(body)
}

func TestSendMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandler))
	defer server.Close()

	testClient.BaseURL = server.URL

	if err := testClient.SendMessage("#channel", "message"); err != nil {
		t.Error(err)
	}
}
