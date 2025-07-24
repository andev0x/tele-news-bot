package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/andev0x/tele-news-bot/config"
	"github.com/andev0x/tele-news-bot/feed"
	"github.com/andev0x/tele-news-bot/notify"
	"github.com/andev0x/tele-news-bot/store"
)

func main() {
	cfg := config.LoadConfig()
	sentItems := store.LoadSentItems()

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
						notify.SendMessage(cfg.BotToken, cfg.ChatID, message)
						store.MarkSent(item.Link)
						sentItems[item.Link] = true
					}
				}
			}(feedURL)
		}
		wg.Wait()
		time.Sleep(10 * time.Minute)
	}
}
