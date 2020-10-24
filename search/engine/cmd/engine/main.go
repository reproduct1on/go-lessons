package main

import (
	"bufio"
	"fmt"
	"os"
	"search/crawler/pkg/spider"
	"strings"
)

func main() {
	var urls = []string{"https://www.ya.ru", "https://www.mail.ru", "https://go.dev/"}
	var data []map[string]string
	for _, url := range urls {
		tmp, err := spider.Scan(url, 2)
		if err != nil {
			fmt.Printf("Ошибка при сканировании сайта: %s ", url)
			return
		}
		data = append(data, tmp)
	}
	if len(data) == 0 {
		fmt.Println("Нет отсканируемых данных")
		return
	}
	fmt.Println("Успешное сканирование")
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	for str != "exit" {
		fmt.Println(`Введите слово для поиска или "exit" для выхода: `)
		scanner.Scan()
		str = scanner.Text()
		if str != "exit" && str != "" {
			search(data, str)
		}
	}
}

// Предлагает пользователю ввести слово, считывает его и осуществляет поиск совпадений
// Выводит совпадения
func search(data []map[string]string, str string) {
	var flg bool
	for _, el := range data {
		for key, val := range el {
			if strings.Contains(strings.ToLower(val), strings.ToLower(str)) || strings.Contains(strings.ToLower(key), strings.ToLower(str)) {
				fmt.Println("Найдено совпадение: ", key, "-", val)
				flg = true
			}
		}
	}
	if !flg {
		fmt.Println("Ничего не найдено")
	}
	return

}
