package conf

import (
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Sv struct {
		Timeout int64
		Port    string
		Debug   bool
	}
	Gcp struct {
		ProjectID string `mapstructure:"project_id"`
	}
	API struct {
		PostURL    string `mapstructure:"post_url"`
		AuthURL    string `mapstructure:"auth_url"`
		ProfileURL string `mapstructure:"profile_url"`
		ChatURL    string `mapstructure:"chat_url"`
	}
	Auth struct {
		PubJwtkey string `mapstructure:"pub_jwtkey"`
	}
	Graphql struct {
		Playground string
		Endpoint   string
	}
}

var C config

func init() {

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalln(err)
	}

	spew.Dump(C)
}
