package handler

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/kelseyhightower/envconfig"
	"github.com/ookkoouu/twiutil"
	"golang.org/x/oauth2/clientcredentials"
)

type twitterEnv struct {
	TwitterApiKey    string `required:"true" split_words:"true"`
	TwitterApiSecret string `required:"true" split_words:"true"`
}

var twClient *twitter.Client

func initTwitter() {
	var twenv twitterEnv
	err := envconfig.Process("", &twenv)
	if err != nil {
		log.Fatalln("twitter.go: Env is not found.")
		panic(err)
	}

	conf := &clientcredentials.Config{
		ClientID:     twenv.TwitterApiKey,
		ClientSecret: twenv.TwitterApiSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := conf.Client(context.Background())
	twClient = twitter.NewClient(httpClient)
}

func getTweet(id int64) (tweet *twitter.Tweet, err error) {
	tweet, _, err = twClient.Statuses.Show(id, nil)
	return
}

func ExpandTwitter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	tweetIDs := twiutil.FindIdAll(m.Content)
	if len(tweetIDs) > 0 {
		log.Println("ExpandTwitter called")
		for _, id := range tweetIDs {
			tweet, err := getTweet(id)
			if err != nil {
				return
			}
			medias := twiutil.GetMediaUrlsString(*tweet)
			s.ChannelMessageSend(m.ChannelID, strings.Join(medias[1:], "\n"))

			log.Println(id, medias)
		}
	}
}
