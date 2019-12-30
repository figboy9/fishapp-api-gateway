package conf

import (
	"fmt"
	"log"
	"os"
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
	Grpc struct {
		PostURL string `mapstructure:"post_url"`
		UserURL string `mapstructure:"user_url"`
	}
	Auth struct {
		PubJwtkey string `mapstructure:"pub_jwtkey"`
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
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
