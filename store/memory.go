package store

var SendItems = make(map[string]bool)

func AlreadySent(link string) bool {
	return SendItems[link]
}

func MarkSent(link string) {
	SendItems[link] = true
}
