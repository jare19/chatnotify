package slack

import (
	"fmt"
	"testing"
)

var uRLs = []struct {
	Pay *Payload
	Res string
}{
	{
		&Payload{
			Client: &Client{
				WebHookURL: "https://hooks.fake.com/something/",
			},
		}, "https://hooks.fake.com/something/"}, {

		&Payload{
			Client: &Client{
				WebHookURL: "",
			},
		}, "Webhook required"},
}

func TestNewClient(t *testing.T) {
	for _, p := range uRLs {
		x := p.Pay.NewClientURL()
		if x != p.Res {
			t.Fail()
		}
	}
}

func TestNewMessage(t *testing.T) {
	var nullMessage Payload

	pay := Payload{
		Message: &MessageEnvelope{
			IconEmoji: ":gopher_dance:",
			Username:  "Morty",
		},
	}
	y, err := pay.NewMessage()
	if y == nil || err != nil {
		fmt.Println(err)
	}

	_, err = nullMessage.NewMessage()
	if err == nil {
		t.Fail()
	}

}
func TestSendMessage(t *testing.T) {
	nullMessage := Payload{
		Client: &Client{
			WebHookURL: "https://hooks.fake.com/something/",
		},
	}
	sendMessage := Payload{
		Client: &Client{
			WebHookURL: "https://hooks.fake.com/something/",
		},
		Message: &MessageEnvelope{
			IconEmoji: ":gopher_dance:",
			Username:  "Morty",
			Attachments: []*MessageAttachments{
				&MessageAttachments{
					Color:   "green",
					Pretext: "This is a Test",
					Fields: []*AttachmentFields{
						&AttachmentFields{
							Title: "Super Important",
							Value: "Very HIGH",
							Short: false,
						},
					},
				},
			},
		},
	}
	err := sendMessage.SendMessage()
	if err != nil {
		t.Fail()
	}
	err = nullMessage.SendMessage()
	if err == nil {
		t.Fail()
	}
}
