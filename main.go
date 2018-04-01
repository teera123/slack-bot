package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)

	bot.Hear("(?i)(หวัดดี|ดีจ้า|สวัสดี).*").MessageHandler(helloHandler)
	go bot.Run()

	// heroku requires the process to bind port or it is killed
	r := gin.New()
	r.POST("/events", eventsHandler)
	r.Run(":" + os.Getenv("PORT"))
}

func eventsHandler(c *gin.Context) {
	var req struct {
		Token     string `json:"token"`
		Challenge string `json:"challenge"`
		Type      string `json:"type"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		log.Println("event handler error:", err)
		return
	}
	log.Printf("event received: %+v\n", req)

	res := struct {
		Challenge string `json:"challenge"`
	}{req.Challenge}
	c.JSON(http.StatusOK, res)
}

func helloHandler(_ context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "สวัสดีค่ะ ヾ（〃＾∇＾）ﾉ♪", slackbot.WithTyping)
}
