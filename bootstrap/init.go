package bootstrap

import (
	"gobot/bot"
	"gobot/models"
	"gobot/pkg/log"
	"gobot/pkg/settings"
)

func Init() {
	settings.Init()
	log.Init()
	bot.Init()
	models.Init()
	log.Info("资源初始化完成...")
}
