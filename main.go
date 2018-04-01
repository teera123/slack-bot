package main

import (
	"os"
	"context"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func main()  {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)

	bot.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)
	bot.Run()
}

func helloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "สวัสดีค่ะ ヾ（〃＾∇＾）ﾉ♪", slackbot.WithTyping)
}
