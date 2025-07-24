package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/andev0x/tele-news-bot/config"
	"github.com/andev0x/tele-news-bot/feed"
	"github.com/andev0x/tele-news-bot/notify"
	"github.com/andev0x/tele-news-bot/store"
	tele "gopkg.in/telebot.v3"
)

func main() {
	cfg := config.LoadConfig()
	sentItems := store.LoadSentItems()

	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Bot started! I will now check for news every 10 minutes.")
	})

	go func() {
		for {
			var wg sync.WaitGroup
			for _, feedURL := range cfg.Feeds {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					items := feed.FetchFeed(url)
					for _, item := range items {
						if !sentItems[item.Link] {
							message := fmt.Sprintf("%s\n%s", item.Title, item.Link)
							notify.SendMessage(b, cfg.ChatID, message)
							store.MarkSent(item.Link)
							sentItems[item.Link] = true
						}
					}
				}(feedURL)
			}
			wg.Wait()
			time.Sleep(10 * time.Minute)
		}
	}()

	b.Start()
}
