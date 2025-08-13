package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	subscribers = make(map[int64]bool)
	mutex       = &sync.Mutex{}
)

func LoadSubscribers() {
	file, err := os.OpenFile("subscribers.json", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error opening subscribers file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading subscribers file:", err)
		return
	}

	if len(data) == 0 {
		return
	}

	var ids []int64
	if err := json.Unmarshal(data, &ids); err != nil {
		log.Println("Error unmarshalling subscribers:", err)
		return
	}

	for _, id := range ids {
		subscribers[id] = true
	}
}

func AddSubscriber(chatID int64) {
	mutex.Lock()
	defer mutex.Unlock()

	subscribers[chatID] = true
	saveSubscribers()
}

func RemoveSubscriber(chatID int64) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(subscribers, chatID)
	saveSubscribers()
}

func GetSubscribers() []string {
	mutex.Lock()
	defer mutex.Unlock()

	var ids []string
	for id := range subscribers {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	return ids
}

func saveSubscribers() {
	var ids []int64
	for id := range subscribers {
		ids = append(ids, id)
	}

	data, err := json.Marshal(ids)
	if err != nil {
		log.Println("Error marshalling subscribers:", err)
		return
	}

	if err := ioutil.WriteFile("subscribers.json", data, 0644); err != nil {
		log.Println("Error writing subscribers file:", err)
	}
}
