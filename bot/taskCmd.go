package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gobot/models"
	"gobot/pkg/log"
	"strconv"
	"time"
)

func taskCmd(update tgbotapi.Update) {
	task := update.Message.CommandArguments()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	switch task {
	case "":
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
		}
		msg.Text = "请在 /task 后跟上要记录的任务内容。"
	case "weekly":
		msg.ParseMode = "markdown"
		msg.Text = weeklyReport()
	default:
		_, message := recordTask(task)
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = message
	}
	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func recordTask(task string) (res bool, msg string) {
	return models.CreateTask(task)
}

func weeklyReport() string {
	start := getFirstDateOfWeek()
	end := time.Now().Format("2006-01-02")
	tasks := models.QueryTaskByUpdateTime(start+" 00:00:00", end+" 23:59:59")
	weeklyReport := "本周（" + start + "～" + end + "）周报：\n"
	for i, t := range tasks {
		weeklyReport += strconv.Itoa(i+1) + ". " + t.Content + " \n"
	}
	return weeklyReport
}
