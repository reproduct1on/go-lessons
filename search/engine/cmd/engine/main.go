package main

import (
	"bufio"
	"fmt"
	"os"
	"search/crawler/pkg/spider"
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
	scanner := bufio.NewScanner(os.Stdin)
	for str != "exit" {
		fmt.Println(`Введите слово для поиска или "exit" для выхода: `)
		scanner.Scan()
		str = scanner.Text()
		if str != "exit" && str != "" {
			fmt.Println("Найдено совпадение: ", search(data, str))
		}
	}
}

// Предлагает пользователю ввести слово, считывает его и осуществляет поиск совпадений
// Выводит совпадения
func search(data map[string]string, str string) (result string) {
	var flg bool
	for key, val := range data {
		if strings.Contains(strings.ToLower(val), strings.ToLower(str)) || strings.Contains(strings.ToLower(key), strings.ToLower(str)) {
			result = fmt.Sprintf("%s - %s", key, val)
			flg = true
		}
	}
	if !flg {
		result = "Ничего не найдено"
	}
	return result

}
