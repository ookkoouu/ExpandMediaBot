package handler

import (
	"emb/applog"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	handles = append(handles, logJoinGuild)
	handles = append(handles, logLeaveGuild)
}

func logJoinGuild(s *discordgo.Session, g *discordgo.GuildCreate) {
	joined, err := g.JoinedAt.Parse()
	if err != nil {
		return
	}
	if math.Abs(time.Since(joined).Seconds()) < 30 {
		log.Printf("Guild joined to %s (total %d)", g.ID, len(s.State.Guilds))
		text := fmt.Sprintf(":new: %s に追加されました (%d)", g.ID, len(s.State.Guilds))
		applog.Discord(text)
	}
}

func logLeaveGuild(s *discordgo.Session, g *discordgo.GuildDelete) {
	text := fmt.Sprintf(":free: %s から退出しました (%d)", g.ID, len(s.State.Guilds))
	applog.Discord(text)
}
