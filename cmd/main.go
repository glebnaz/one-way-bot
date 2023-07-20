package main

import (
	"github.com/glebnaz/one-way-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"os"
)

type config struct {
	Token string `envconfig:"TG_TOKEN"`
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
}

func main() {
	var cfg config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Error process config: %s", err)
	}

	log.Info("bot started")

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatalf("Error create bot: %s", err)
	}
	bot.Debug = false

	botService := service.NewBot(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			err := botService.ProcessUpdate(update)
			if err != nil {
				log.Errorf("Error process update: %s", err)
				continue
			}
			log.Debug("message process successfully")
		}
	}
}
