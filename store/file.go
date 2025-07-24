package store

import (
	"bufio"
	"log"
	"os"
)

const sentItemsFile = "sent_items.txt"

func LoadSentItems() map[string]bool {
	sentItems := make(map[string]bool)
	file, err := os.Open(sentItemsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return sentItems
		}
		log.Println("Error loading sent items:", err)
		return sentItems
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing sent items file:", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sentItems[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading sent items:", err)
	}

	return sentItems
}

func MarkSent(link string) {
	file, err := os.OpenFile(sentItemsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Println("Error opening sent items file:", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing sent items file:", err)
		}
	}()

	if _, err := file.WriteString(link + "\n"); err != nil {
		log.Println("Error writing to sent items file:", err)
	}
}
