package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppEnv struct {
	DcToken      string `required:"true" split_words:"true"`
	DcLogWebhook string `split_words:"true"`

	TwitterApiKey    string `required:"true" split_words:"true"`
	TwitterApiSecret string `required:"true" split_words:"true"`
}

var Env AppEnv

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		log.Println("Env is not fulfilled.")
		panic(err)
	}
}
