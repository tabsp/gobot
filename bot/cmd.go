package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gobot/pkg/log"
)

const (
	Help = "help"
	Sub  = "sub"
	Task = "task"
)

func helpCmd(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I understand /task.")
	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func unknownCmd(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "I don't know that command"
	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}
