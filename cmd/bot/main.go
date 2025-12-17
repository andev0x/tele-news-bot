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
	store.LoadSubscribers()

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
		store.AddSubscriber(c.Chat().ID)
		log.Printf("New subscriber: %d", c.Chat().ID)
		return c.Send("You have subscribed to news updates.")
	})

	b.Handle("/stop", func(c tele.Context) error {
		store.RemoveSubscriber(c.Chat().ID)
		log.Printf("Unsubscribed: %d", c.Chat().ID)
		return c.Send("You have unsubscribed from news updates.")
	})

	go func() {
		for {
			var wg sync.WaitGroup
			for _, feedURL := range cfg.Feeds {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					items := feed.FetchFeed(url)
					log.Printf("Found %d news from %s", len(items), url)
					for _, item := range items {
						if !sentItems[item.Link] {
							log.Printf("Sending new item: %s", item.Title)
							message := fmt.Sprintf("%s\n%s", item.Title, item.Link)
							subscribers := store.GetSubscribers()
							for _, subscriber := range subscribers {
								notify.SendMessage(b, subscriber, message)
							}
							store.MarkSent(item.Link)
							sentItems[item.Link] = true
						}
					}
				}(feedURL)
			}
			wg.Wait()
			log.Println("Waiting for next check...")
			time.Sleep(10 * time.Minute)
		}
	}()

	log.Println("Bot started! I will now check for news every 10 minutes.")
	b.Start()
}
