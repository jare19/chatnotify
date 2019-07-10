package hipchat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	defaultURL = "https://api.hipchat.com"
	notifyPath = "%s/v2/room/%s/notification?auth_token=%s"
)

// Payload is what is needed to send a HipChat message
type Payload struct {
	URL     string
	Room    string
	Token   string
	Message *MessageEnvelope
}

// MessageEnvelope represents the HipChat notification message.
type MessageEnvelope struct {
	From    string `json:"from"`
	Color   string `json:"color"`
	Message string `json:"message"`
}

// NewClientURL returns a new HipChat Client.
func (p *Payload) NewClientURL() string {
	if p.URL == "" {
		p.URL = defaultURL
	}

	return fmt.Sprintf(notifyPath, p.URL, p.Room, p.Token)

}

// NewMessage formats the MessageEnvelope into a *bytes.Reader for the HTTP POST
func (p *Payload) NewMessage() (*bytes.Reader, error) {
	data, err := json.Marshal(p.Message)
	if err != nil {
		return nil, err
	}
	if len(data) <= 2 {
		return nil, errors.New("message format is not correct")
	}
	body := bytes.NewReader(data)

	return body, err
}

// SendMessage takes a HipChat payload and sends to HipChat
func (p *Payload) SendMessage() error {

	if p.Token == "" {
		return errors.New("Token required")
	}
	clientURL := p.NewClientURL()
	message, err := p.NewMessage()
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post(clientURL, "application/json", message)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		t, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("HipChatError: %d %s", resp.StatusCode, string(t))
	}

	return nil
}
