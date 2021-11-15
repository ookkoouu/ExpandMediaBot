package handler

import "github.com/bwmarrin/discordgo"

var handles = make([]interface{}, 0)
var destructs = make([]func(), 0)

func Regist(dcs *discordgo.Session) {
	for _, f := range handles {
		dcs.AddHandler(f)
	}
}

func Exit() {
	for _, f := range destructs {
		f()
	}
}
