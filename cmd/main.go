package main

import (
	"emb/cmd/env"
	"emb/cmd/handler"
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

	// ハンドラ設定
	handler.Regist(dcs)

	// ws接続
	err = dcs.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		handler.Exit()
		dcs.Close()
		logDiscord(":no_entry: 停止しました")
	}()

	log.Printf("\x1b[32m%s\x1b[0m", "Bot started...")
	logDiscord(":white_check_mark: 起動しました")
}
