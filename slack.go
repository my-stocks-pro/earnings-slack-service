package main

import (
	"github.com/nlopes/slack"
	"fmt"
	"time"
)

type TypeSlackService struct {
	Host   string
	Port   string
	Config *TypeConfig
	Logger *TypeLogger
	Client *slack.Client
}

func NewService() *TypeSlackService {
	s := &TypeSlackService{}

	s.Logger = NewLogger("earnings-slack-service")
	s.Config = s.LoadConfig()
	s.Host, s.Port = s.GetSevicePath()
	s.Client = slack.New("YOUR_TOKEN_HERE")
	s.Logger.Info.Println("Create slack Client...")

	return s
}

func NewAttachment(title string, imageURL string) []slack.Attachment {
	return []slack.Attachment{
		slack.Attachment{
			Title:    title,
			ImageURL: imageURL,
		},
	}
}

func NewMessage(idi string, imageLink string, location string) string {
	return fmt.Sprintf("Date: %s\nLocation: %s\nID: %s",
		time.Now().Format("2006-01-02"),
		location,
		fmt.Sprintf("<%s|%s>", imageLink, idi))
}

func (s *TypeSlackService) PostMessage(data Data) {
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	s.Client.SetDebug(true)

	message := NewMessage(data.IdI, data.ImageLink, data.Location)

	params := slack.PostMessageParameters{}
	params.Attachments = NewAttachment(data.Description, data.ImageURL)

	channelID, timestamp, err := s.Client.PostMessage(s.Config.Channel, message, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
