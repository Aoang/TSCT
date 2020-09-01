package tsct

import (
	"encoding/json"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"net/url"
)

var api *tg.BotAPI

func Bot(token string) {
	api, _ = tg.NewBotAPI(token)
}

func Webhook(urls string) {
	u, _ := url.Parse(urls)
	_, _ = api.SetWebhook(tg.WebhookConfig{
		URL: u,
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	var update tg.Update
	_ = json.Unmarshal(bytes, &update)

	d := Find(update.Message.Chat.ID)
	if d == nil {
		return
	}

	_, _ = api.Send(tg.MessageConfig{
		BaseChat: tg.BaseChat{
			ChatID: update.Message.Chat.ID,
		},
		Text:      d.GET().String(),
		ParseMode: "MarkdownV2",
	})
}
