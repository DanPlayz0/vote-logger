package config

import (
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		TestMode string `env:"NODE_ENV" envDefault:"development"`

		BotList struct {
			Topgg struct {
				Auth string `env:"AUTH"`
			} `envPrefix:"TOPGG_"`
			Dscbot struct {
				Auth string `env:"AUTH"`
			} `envPrefix:"DSCBOT_"`
			Dlist struct {
				Auth string `env:"AUTH"`
			} `envPrefix:"DLIST_"`
		}
		WebhookURL string `env:"WEBHOOK_URL"`
		Port       string `env:"PORT" envDefault:"3340"`
	}
)

var Conf Config

func Parse() {
	var err error
	if _, err = os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}
}
