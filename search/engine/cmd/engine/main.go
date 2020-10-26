package main

import (
	"bufio"
	"fmt"
	"os"
	"search/crawler/pkg/spider"
	"search/indexer/pkg/index"
	"strings"
)

func main() {
	ws := spider.NewWebScan([]string{"https://www.ya.ru", "https://www.mail.ru", "https://go.dev/"})
	data, _ := ws.Scan(2)
	if len(data) == 0 {
		fmt.Println("Нет отсканируемых данных")
		return
	}
	fmt.Println("Успешное сканирование")
	var str string
	tmp, indexFile := index.Create(data)
	scanner := bufio.NewScanner(os.Stdin)
	for str != "exit" {
		fmt.Println(`Введите слово для поиска или "exit" для выхода: `)
		scanner.Scan()
		str = scanner.Text()
		if str != "exit" && str != "" {
			matches := indexFile[strings.ToLower(str)]
			if len(matches) == 0 {
				fmt.Println("Ничего не найдено")
			}
			if len(matches) > 0 {
				for _, m := range matches {
					tmp := index.BSearch(*tmp, 0, len(*tmp), m)
					fmt.Println("Найдено совпадение: ", tmp.Title, "-", tmp.URL)
				}
			}
		}
	}
}
