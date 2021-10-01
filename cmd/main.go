package main

import (
	"emb/cmd/handler"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	DcTokenDev string `required:"true" split_words:"true"`
}

func main() {
	var env env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalln("main.go: Env is not found.")
		panic(err)
	}

	// ログイン
	dcs, err := discordgo.New("Bot " + env.DcTokenDev)
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

	handler.Init()

	dcs.AddHandler(handler.ExpandTwitter)

	log.Printf("\x1b[32m%s\x1b[0m", "Bot started...")
}
