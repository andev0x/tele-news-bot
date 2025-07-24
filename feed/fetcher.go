package feed

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title string
	Link  string
}

func FetchFeed(url string) []Item {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return nil
	}

	var items []Item
	for _, entry := range feed.Items {
		if entry.PublishedParsed != nil && time.Since(*entry.PublishedParsed) < 2*time.Hour {
			items = append(items, Item{Title: entry.Title, Link: entry.Link})
		}
	}
	return items
}
