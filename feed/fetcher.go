package feed

import (
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title string
	Link  string
}

func FetchFeed(url string) []Item {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fp := gofeed.NewParser()
	fp.Client = client

	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Printf("Error fetching feed from %s: %v", url, err)
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
