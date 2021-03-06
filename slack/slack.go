package slack

import (
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/webhook"
)

// deprecated, use CustomizedLog instead 
func Log(from, message string) {
	slackurl := "https://hooks.slack.com/services/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"
	var slackChannel = webhook.New(slackurl)
	attachments := make([]*chat.Attachment, 1)
	attachments = append(attachments, &chat.Attachment{
		Title: "MIGRATION ERROR",
		Color: "#FF2D00",
		Text:  message,
	})
	slack_msg := "*Message from 👉* \n> " + from
	m := &chat.Message{Text: slack_msg, Attachments: attachments}
	m.Send(slackChannel)
}

// more customized slack message for debugging purpose
func CustomizedLog(from, title, message, slackurl string) {
	var slackChannel = webhook.New(slackurl)
	attachments := make([]*chat.Attachment, 1)
	attachments = append(attachments, &chat.Attachment{
		Title: title,
		Color: "#FF2D00",
		Text:  message,
	})
	slack_msg := "*Message from 👉* \n> " + from
	m := &chat.Message{Text: slack_msg, Attachments: attachments}
	m.Send(slackChannel)
}
