package chatnotify

import (
	"fmt"
	"testing"

	"github.com/jare19/chatnotify/chatops/hipchat"
	"github.com/jare19/chatnotify/chatops/slack"
)

var slackPayload = []*slack.Payload{
	{
		Client: &slack.Client{
			WebHookURL: "https://hooks.fake.com/something/",
		},
		Message: &slack.MessageEnvelope{
			IconEmoji: ":gopher_dance:",
			Username:  "Morty",
			Attachments: []*slack.MessageAttachments{
				{
					Color:   "green",
					Pretext: "This is a Test",
					Fields: []*slack.AttachmentFields{
						{
							Title: "Super Important",
							Value: "Very HIGH",
							Short: false,
						},
					},
				},
			},
		},
	},

	{
		Client: &slack.Client{
			WebHookURL: "",
		},
		Message: &slack.MessageEnvelope{
			IconEmoji: ":gopher_dance:",
			Username:  "Morty",
			Attachments: []*slack.MessageAttachments{
				{
					Color:   "green",
					Pretext: "This is a Test",
					Fields: []*slack.AttachmentFields{
						{
							Title: "Super Important",
							Value: "Very HIGH",
							Short: false,
						},
					},
				},
			},
		},
	},
	{
		Client: &slack.Client{
			WebHookURL: "https://hooks.fake.com/something/",
		},
		Message: &slack.MessageEnvelope{
			Text:      "--- Results ---",
			IconEmoji: ":morty:",
			Username:  "Drone user",
			Attachments: []*slack.MessageAttachments{
				{
					Fallback:   "--- Results --- ",
					AuthorName: "Drone - droneusername",
					AuthorLink: "http://www.somelink.com",
					AuthorIcon: "https://avatars2.githubusercontent.com/u/2181346?s=200&v=4",
					Color:      "danger",
					Title:      "Violations Detected!",
					Text:       "*this is not approved.*",
					Fields: []*slack.AttachmentFields{
						{
							Title: "Total Violations",
							Value: "8",
							Short: true,
						},
						{
							Title: "Id",
							Value: "9",
							Short: true,
						},
						{
							Title: "Registry",
							Value: "fake Registry",
							Short: true,
						},
						{
							Title: "Repo",
							Value: "some repo",
							Short: true,
						},
						{
							Title: "Tag",
							Value: "latest",
							Short: true,
						},
						{
							Title: "Digest",
							Value: "some sha",
							Short: true,
						},
						{
							Title: "First Seen",
							Value: "date&time",
							Short: true,
						},
						{
							Title: "Last Seen",
							Value: "date&time",
							Short: true,
						},
					},
				},
			},
		},
	},
}
var hcPayload = []*hipchat.Payload{
	&hipchat.Payload{
		URL:   "fakeURL",
		Room:  "464564645646",
		Token: "456456456456456",
		Message: &hipchat.MessageEnvelope{
			Color:   "red",
			From:    "Results",
			Message: "<b>%Violations Detected!</b>",
		},
	},
}

func TestSlack(t *testing.T) {
	for _, sp := range slackPayload {
		mrkdwn := sp.Message.Attachments

		mrkdwn[0].Color = "good"
		mrkdwn[0].Title = "this has been changed \n and this is the newline"

		err := Send(sp)
		if err != nil {
			fmt.Printf("error is %v", err)
		}
	}
}

func TestHipChat(t *testing.T) {
	for _, hc := range hcPayload {
		err := Send(hc)
		if err != nil {
			fmt.Printf("error is %v", err)
		}

	}
}
