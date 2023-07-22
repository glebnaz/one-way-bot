package service

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) ProcessUpdate(update tgbotapi.Update) error {
	if update.Message.NewChatMembers != nil {
		err := b.ProcessNewChatMembers(update.Message.NewChatMembers, update.Message.Chat.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) ProcessNewChatMembers(user []tgbotapi.User, chatID int64) error {
	var err error
	for _, person := range user {
		errSend := b.SendWelcomeMessage(person, chatID)
		if errSend != nil {
			if err == nil {
				err = fmt.Errorf("error send welcome message: %s", errSend)
			} else {
				err = errSend
			}
			continue
		}
	}
	return err
}

func (b *Bot) SendWelcomeMessage(person tgbotapi.User, chatID int64) error {
	message := fmt.Sprintf(welcomeMessageFormat, person.UserName)

	tgmsg := tgbotapi.NewMessage(chatID, message)
	tgmsg.DisableNotification = true

	_, err := b.bot.Send(tgbotapi.NewMessage(chatID, message))
	if err != nil {
		log.Errorf("Error send welcome message: %s", err)
		return err
	}
	return nil
}
