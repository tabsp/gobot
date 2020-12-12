package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gobot/pkg/log"
	"gobot/pkg/settings"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	bot *tgbotapi.BotAPI
)

func Init() {
	token := os.Getenv("TELEGRAM_API_TOKEN")
	if token == "" {
		log.Panic("Token 未设置，通过环境变量 TELEGRAM_API_TOKEN 设置")
	}
	botapi, err := tgbotapi.NewBotAPIWithClient(token, getHttpClient())
	if err != nil {
		log.Panicf("bot API 创建失败", err)
	}
	bot = botapi
	if strings.ToLower(settings.Settings.Log.Level) == "debug" {
		bot.Debug = true
	}
	log.Infof("Token 认证成功！Bot username @%s", bot.Self.UserName)
}

func getHttpClient() *http.Client {
	var transport *http.Transport
	if settings.Settings.Bot.UseProxy {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://127.0.0.1:7890")
		}
		transport = &http.Transport{Proxy: proxy}
	} else {
		transport = &http.Transport{}
	}
	return &http.Client{Transport: transport}
}

func Start() {
	updateHandler()
}

func updateHandler() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if update.Message.IsCommand() { // command Messages
			commandHandler(update)
		}
		if isURLMessage(update.Message) { // command Messages

		}
	}
}

func commandHandler(update tgbotapi.Update) {
	switch strings.ToLower(update.Message.Command()) {
	case Help:
		handle(helpCmd, update)
	case Sub:
		handle(subCmd, update)
	case Task:
		handle(taskCmd, update)
	default:
		handle(unknownCmd, update)
	}
}

func urlHandler(update tgbotapi.Update) {

}

func handle(handler func(u tgbotapi.Update), u tgbotapi.Update) {
	handler(u)
}
