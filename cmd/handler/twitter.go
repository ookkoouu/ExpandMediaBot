package handler

import (
	"context"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/kelseyhightower/envconfig"
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

func getVideoUrl(media twitter.MediaEntity) (url string) {
	variants := make([]twitter.VideoVariant, 1, 4)
	for _, v := range media.VideoInfo.Variants {
		if v.ContentType == "video/mp4" {
			variants = append(variants, v)
		}
	}
	sort.Slice(variants, func(i, j int) bool { return variants[i].Bitrate > variants[j].Bitrate })
	url = variants[0].URL
	return
}

func getTweet(id int64) (tweet *twitter.Tweet, err error) {
	tweet, _, err = twClient.Statuses.Show(id, nil)
	return
}

func getMediaUrls(id int64) (urls []string) {
	tweet, err := getTweet(id)
	if err != nil || len(tweet.ExtendedEntities.Media) == 0 {
		return
	}

	medias := tweet.ExtendedEntities.Media
	if medias[0].Type == "photo" {
		urls = make([]string, 0, 4)
		for _, e := range medias {
			urls = append(urls, e.MediaURLHttps)
		}
		urls = urls[1:]
	} else if medias[0].Type == "video" || medias[0].Type == "animated_gif" {
		urls = append(urls, getVideoUrl(medias[0]))
	}
	return
}

func getIdFromUrl(twUrl string) (id int64) {
	u, err := url.Parse(twUrl)
	if err != nil {
		return 0
	}
	idStr := u.Path[strings.LastIndex(u.Path, "/")+1:]
	id, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return
}

func getTwitterUrl(content string) (urls []string) {
	regexTwitterUrl := regexp.MustCompile(`https?://twitter\.com(/\w+)?/status(es)?/\d+`)
	return regexTwitterUrl.FindAllString(content, -1)
}

func ExpandTwitter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	
	twUrls := getTwitterUrl(m.Content)
	if len(twUrls) > 0 {
		log.Println("ExpandTwitter called")
		for _, url := range twUrls {
			id := getIdFromUrl(url)
			medias := getMediaUrls(id)
			s.ChannelMessageSend(m.ChannelID, strings.Join(medias, "\n"))

			log.Println(id, medias)
		}
	}
}
