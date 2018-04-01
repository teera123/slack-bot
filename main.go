package main

import (
	"context"
	"encoding/json"
	"fmt"
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

type event struct {
	Token      string    `json:"token"`
	TeamID     string    `json:"team_id"`
	APIAppID   string    `json:"api_app_id"`
	Event      eventType `json:"event"`
	Type       string    `json:"type"`
	AuthedUser []string  `json:"authed_user"`
	EventID    string    `json:"event_id"`
	EventTime  int       `json:"event_time"`
	Challenge  string    `json:"challenge"`
}

type eventType struct {
	Type      string `json:"type"`
	EventTS   string `json:"event_ts"`
	User      string `json:"user"`
	Timestamp string `json:"ts"`
	Item      string `json:"item"`
}

func eventsHandler(c *gin.Context) {
	var req event
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		fmt.Println("event handler error:", err)
		return
	}
	fmt.Printf("event received: %+v\n", req)

	if req.Challenge != "" {
		c.JSON(http.StatusOK, struct {
			Challenge string `json:"challenge"`
		}{req.Challenge})
	}

	c.JSON(http.StatusOK, nil)
}

func helloHandler(_ context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "สวัสดีค่ะ ヾ（〃＾∇＾）ﾉ♪", slackbot.WithTyping)
}
