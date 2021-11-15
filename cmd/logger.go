package main

import (
	"bytes"
	"emb/cmd/env"
	"encoding/json"
	"log"
	"net/http"
)

type discordWebhook struct {
	UserName  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Content   string `json:"content"`
	TTS       bool   `json:"tts"`
}

func logDiscord(s string) {
	url := env.Env.DcLogWebhook
	if url == "" {
		return
	}

	msg := new(discordWebhook)
	msg.Content = s
	body, _ := json.Marshal(msg)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("webhook", err)
	}
	defer res.Body.Close()
}
