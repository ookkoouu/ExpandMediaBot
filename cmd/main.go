package main

import (
	"emb/cmd/handler"
	"emb/cmd/env"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// ログイン
	dcs, err := discordgo.New("Bot " + env.Env.DcToken)
	if err != nil {
		panic(err)
	}

	// ws接続
	err = dcs.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		dcs.Close()
	}()

	handler.Regist(dcs)

	log.Printf("\x1b[32m%s\x1b[0m", "Bot started...")
}
