package main

import (
	"fmt"
	"time"

	cfg "tg-stream-notifier/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicklaw5/helix/v2"
	"github.com/rs/zerolog/log"
)

const messageTmpl = "https://twitch.tv/%s"

func main() {
	config, err := cfg.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	token, err := Auth(config.TwitchClientID, config.TwitchSecretKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get access token")
	}
	client, err := helix.NewClient(&helix.Options{ClientID: config.TwitchClientID})
	if err != nil {
		panic(err)
	}
	client.SetAppAccessToken(token)

	bot, err := tgbotapi.NewBotAPI(config.TgBotToken)
	if err != nil {
		log.Panic().Err(err).Msg("create tg bot")
	}

	bot.Debug = true

	log.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	ticker := time.NewTimer(time.Second * 30)
	isOnline := false
	for {
		select {
		case <-ticker.C:
			streams, err := client.GetStreams(&helix.StreamsParams{
				UserLogins: []string{config.StreamerLogin},
			})
			if err != nil {
				log.Err(err).Msg("get streams")
			}
			if len(streams.Data.Streams) > 0 && !isOnline {
				isOnline = true
				if _, err := bot.Send(tgbotapi.NewMessage(
					int64(config.TgChannelID),
					fmt.Sprintf(messageTmpl, config.StreamerLogin))); err != nil {
					log.Err(err).Msg("send message to TG channel")
				}
			}
			ticker.Reset(time.Second * 30)
		}
	}
}
