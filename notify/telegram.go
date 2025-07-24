package notify

import (
	"log"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func SendMessage(b *tele.Bot, chatID, message string) {
	i, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		log.Println("Error parsing chatID", err)
	}
	chat := &tele.Chat{ID: i}

	_, err = b.Send(chat, message)
	if err != nil {
		log.Println("Telegram error:", err)
		return
	}
}
