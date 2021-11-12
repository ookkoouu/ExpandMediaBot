package handler

import "github.com/bwmarrin/discordgo"

var handles = make([]interface{}, 1)

func Regist(dcs *discordgo.Session) {
	for _, f := range handles {
		dcs.AddHandler(f)
	}
}
