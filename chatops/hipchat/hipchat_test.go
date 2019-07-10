package hipchat

import (
	"fmt"
	"testing"
)

var payload = []*Payload{
	&Payload{
		URL:   "fakeURL",
		Room:  "5727867587",
		Token: "78572678274546",
		Message: &MessageEnvelope{
			Color:   "red",
			From:    "Results",
			Message: "<b>%Violations!</b>",
		},
	},
}

func TestNewMessage(t *testing.T) {
	for _, p := range payload {
		y, err := p.NewMessage()
		if y == nil || err != nil {
			fmt.Println(err)
		}
	}
}

func TestSendMessage(t *testing.T) {
	for _, p := range payload {
		err := p.SendMessage()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestNewClientURL(t *testing.T) {
	for _, p := range payload {
		x := p.NewClientURL()
		if x == "" {
			t.Fail()
		}

	}
}
