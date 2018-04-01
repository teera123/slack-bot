package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	//toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	//toMe.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)

	bot.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)
	go bot.Run()

	// heroku requires the process to bind port or it is killed
	r := gin.New()
	r.POST("/events", eventsHandler)
	r.Run(":" + os.Getenv("PORT"))
}

func helloHandler(_ context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	msg := []string{
		"สวัสดีค่ะ ヾ（〃＾∇＾）ﾉ♪",
		"เราชื่อพิมฐายินดีที่ได้รู้จัก ヽ(;^o^ヽ)",
		"หวัดดีจ้าาาา",
		"สวัสดีค่ะ เราชื่อพิมฐา ฝากเนื้อฝากตัวด้วยนะคะ 。（＞ω＜。）",
	}

	bot.Reply(evt, msg[random(0, len(msg))], slackbot.WithTyping)
}

func random(min, max int) int {
	return min + rand.Intn(max-min)
}
