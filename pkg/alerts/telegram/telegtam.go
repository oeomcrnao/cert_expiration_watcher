package telegram

import (
	"log"
	"net/http"
	"net/url"

	"github.com/oeomcrnao/domain_expiration_watcher/pkg/tools/environment"
)

func SendMessage(message string) bool {

	var token, _ = environment.GetStringFromEnv("TELEGRAM_BOT_TOKEN", "")
	if len(token) == 0 {
		return false
	}
	var chatId, _ = environment.GetStringFromEnv("TELEGRAM_CHAT_ID", "")
	if len(token) == 0 {
		return false
	}

	var apiString = "https://api.telegram.org/bot" + token + "/sendMessage"

	_, err := http.PostForm(
		apiString,
		url.Values{
			"chat_id": {chatId},
			"text":    {message},
		})
	if err != nil {
		log.Fatal(err)
	}
	return true

}
