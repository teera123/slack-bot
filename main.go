package main

import (
	"os"
	"context"
	"net/http"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func main()  {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	//toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	//toMe.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)

	bot.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)
	bot.Run()

	// heroku requires the process to bind port or it is killed
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}

func helloHandler(_ context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "สวัสดีค่ะ ヾ（〃＾∇＾）ﾉ♪", slackbot.WithTyping)
}
