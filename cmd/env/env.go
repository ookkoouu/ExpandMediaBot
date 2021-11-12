package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppEnv struct {
	DcToken string `required:"true" split_words:"true"`
	BotID   string `required:"true" split_words:"true"`
	OwnerID string `required:"true" split_words:"true"`

	TwitterApiKey    string `required:"true" split_words:"true"`
	TwitterApiSecret string `required:"true" split_words:"true"`
}

var Env AppEnv

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		log.Fatalln("Env is not fulfilled.")
		panic(err)
	}
}
