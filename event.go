package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	Type      string  `json:"type"`
	EventTS   float64 `json:"event_ts"`
	User      string  `json:"user"`
	Timestamp string  `json:"ts"`
	Item      string  `json:"item"`
}

func eventsHandler(c *gin.Context) {
	var req event
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		fmt.Println("event handler error:", err)
		return
	}
	defer c.Request.Body.Close()
	fmt.Printf("event received: %+v\n", req)

	if req.Challenge != "" {
		c.JSON(http.StatusOK, struct {
			Challenge string `json:"challenge"`
		}{req.Challenge})
	}

	c.JSON(http.StatusOK, nil)
}
