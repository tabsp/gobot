package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gobot/pkg/log"
)

func subCmd(update tgbotapi.Update) {
	url := getURLFromMessage(update.Message)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	if url == "" {
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
		}
		//msg.Text = "请回复要订阅的 URL"
		msg.Text = "开发中..."
	} else {
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "订阅成功！！！"
	}
	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func replySubCmd(update tgbotapi.Update) {

}
