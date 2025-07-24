package notify

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func SendMessage(botToken, chatID, message string) {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	resp, err := http.PostForm(endpoint, url.Values{
		"chat_id": {chatID},
		"text":    {message},
	})
	if err != nil {
		log.Println("Telegram error:", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("error closing telegram response body:", err)
		}
	}()
}
