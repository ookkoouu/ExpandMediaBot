package handler

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/kelseyhightower/envconfig"
	"github.com/ookkoouu/twiutil/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type twitterEnv struct {
	TwitterApiKey    string `required:"true" split_words:"true"`
	TwitterApiSecret string `required:"true" split_words:"true"`
}

var twClient *twitter.Client

func init() {
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

	handles = append(handles, ExpandTwitter)
}

func getTweet(id int64) (tweet *twitter.Tweet, err error) {
	tweet, _, err = twClient.Statuses.Show(id, &twitter.StatusShowParams{TweetMode: "extended"})
	return
}

func getTextMedia(tweet *twitter.Tweet) (text string) {
	urls := twiutil.GetMediaUrls(tweet)
	mtypes := twiutil.GetMediaTypes(tweet)

	if urls == nil {
		return
	}
	if mtypes[0] == "photo" {
		if len(urls) < 2 {
			return
		}
		text = strings.Join(urls[1:], "\n") + "\n"
	} else {
		text = strings.Join(urls, "\n") + "\n"
	}
	return
}

func getTextQuoted(tweet *twitter.Tweet) (text string) {
	quoted := twiutil.GetQuotedTweetUrl(tweet)
	if quoted != "" {
		text = "引用RT\n" + quoted + "\n"
	}
	return
}

func ExpandTwitter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	tweetIDs := twiutil.FindIdAll(m.Content)
	if len(tweetIDs) == 0 {
		return
	}

	for _, id := range tweetIDs {
		postText := ""
		tweet, err := getTweet(id)
		if err != nil {
			log.Printf("\x1b[33m%s\x1b[0m", err)
			return
		}

		// メディア
		postText += getTextMedia(tweet)

		// 引用RT
		postText += getTextQuoted(tweet)

		if postText != "" {
			postText = strings.TrimRight(postText, "\n")
			s.ChannelMessageSend(m.ChannelID, postText)
			log.Println(id, strings.ReplaceAll(postText, "\n", `\n`))
		}
	}
}
