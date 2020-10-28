package main

import (
	"search/crawler/pkg/spider"
	"testing"
)

func Test_search(t *testing.T) {
	s := new(spider.LocScan)
	data, _ := scan(s, []string{})
	got := search(data, "почта")
	want := "https://www.mail.ru - Mail.ru: почта, поиск в интернете, новости, игры"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}
