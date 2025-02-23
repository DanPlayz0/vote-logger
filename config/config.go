package config

import (
	"github.com/caarlos0/env/v10"
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
	}
)

var Conf Config

func Parse() {
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}
}
