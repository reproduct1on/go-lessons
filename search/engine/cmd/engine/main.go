package main

import (
	"log"

	"search/engine/pkg/finder"
)

func main() {
	var urls = []string{"https://www.ya.ru","https://www.mail.ru","https://go.dev/"}
	err := finder.Find(urls)
	if err != nil {
		log.Println(err)
	}
}