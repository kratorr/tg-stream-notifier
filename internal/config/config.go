package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TwitchClientID  string `yaml:"TWITCH_CLIENT_ID" env:"TWITCH_CLIENT_ID"`
	TwitchSecretKey string `yaml:"TWITCH_SECRET_KEY" env:"TWITCH_SECRET_KEY"`
	StreamerLogin   string `yaml:"STREAMER_LOGIN" env:"STREAMER_LOGIN"`
	TgBotToken      string `yaml:"TG_BOT_TOKEN" env:"TG_BOT_TOKEN"`
	TgChannelID     int    `yaml:"TG_CHANNEL_ID" env:"TG_CHANNEL_ID"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("/home/kratorr/GolandProjects/tg-stream-notifier/config.yaml", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
