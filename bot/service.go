package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func getURLFromMessage(m *tgbotapi.Message) (url string) {
	for _, entity := range *m.Entities {
		if entity.Type == "url" {
			if url == "" {
				url = m.Text[entity.Offset : entity.Offset+entity.Length]
			}
		}
	}
	return
}

func isURLMessage(m *tgbotapi.Message) bool {
	if m.Entities == nil || len(*m.Entities) == 0 {
		return false
	}
	entity := (*m.Entities)[0]
	return entity.Offset == 0 && entity.Type == "url"
}

func getFirstDateOfWeek() string {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := weekStartDate.Format("2006-01-02")
	return weekMonday
}
