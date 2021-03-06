package slackposter

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&response{
			Ok: true,
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}))
	defer ts.Close()

	client := NewClient("token")
	client.SetUri(ts.URL)

	if err := client.SendMessage(context.Background(), "#channel", "text", nil); err != nil {
		t.Errorf("should not be fail: %v", err)
	}
}
