package handler

import (
	"emb/cmd/env"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	regSpaceStr = `[\s\v\x{00a0}\x{1680}\x{180e}\x{2000}-\x{200b}\x{2028}\x{2029}\x{202f}\x{205f}\x{3000}\x{feff}]`
	filename    = "./ignore.txt"
)

var (
	regStartWithMention = regexp.MustCompile(`^<@!?` + env.Env.BotId + `>` + regSpaceStr + `*`)
	regSpace            = regexp.MustCompile(regSpaceStr)
	ignoreChannels      = make(map[string]int)
)

func init() {
	handles = append(handles, ignore)
	destructs = append(destructs, exitIgnore)
	readIgnore()
}

func readIgnore() {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fp, err := os.Create(filename)
		if err != nil {
			log.Println("Failed to read ignore")
			panic(err)
		}
		fp.Close()
	}

	for _, v := range strings.Split(string(text), "\n") {
		ignoreChannels[v] = 0
	}
	delete(ignoreChannels, "")
}

func saveIgnore() {
	ids := make([]string, 0)
	for id := range ignoreChannels {
		ids = append(ids, id)
	}

	text := strings.Join(ids, "\n")
	err := ioutil.WriteFile(filename, []byte(text), 0664)
	if err != nil {
		log.Println("Failed to save ignore")
		log.Println(ignoreChannels)
		log.Println(err)
	}
}

func addIgnore(id string) {
	ignoreChannels[id] = 0
	log.Println("Add ignore:", id)
}

func removeIgnore(id string) {
	delete(ignoreChannels, id)
	log.Println("Delete ignore:", id)
}

func isIgnoreChannel(channelID string) bool {
	for id := range ignoreChannels {
		if channelID == id {
			return true
		}
	}
	return false
}

func exitIgnore() {
	saveIgnore()
}

func ignore(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !regStartWithMention.MatchString(m.Content) {
		return
	}

	args := regSpace.Split(regStartWithMention.ReplaceAllString(m.Content, ""), -1)
	switch strings.ToLower(args[0]) {
	case "off":
		addIgnore(m.ChannelID)
	case "on":
		removeIgnore(m.ChannelID)
	}

	saveIgnore()
}
