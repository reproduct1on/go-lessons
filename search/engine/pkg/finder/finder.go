// Package finder реализует поисковый движок для сайтов.
// Позволяет искать необходимые страницы сайтов по ключевым словам.
package finder

import (
	"fmt"
	"os"
	"strings"

	"search/crawler/pkg/spider"
)

// Вызывает метод Scan пакета spider для каждого сайта из массива urls
// Результаты сохраняет в слайс, где каждый элемент результат обхода по одному сайту
// После сохранения данных вызывает функцию search
func Find(urls []string) ( err error) {
	var data []map[string]string
	for _, url := range urls{
		tmp, err := spider.Scan(url, 2)
		if err != nil{
			return  fmt.Errorf("Ошибка при сканировании сайта: %s ", url )
		}
		data = append(data,tmp)
	}
	if len(data) > 0{
		fmt.Println("Успешное сканирование")
		search(data)
	}

	return  nil
}

// Предлагает пользователю ввести слово, считывает его и осуществляет поиск совпадений
// Вызывает себя же, чтобы позволить пользователю еще раз провести поиск 
func search(data []map[string]string) {
	var str string
	fmt.Println(`Введите слово для поиска или "exit" для выхода: `)
	fmt.Fscan(os.Stdin, &str)
	if str == "exit"{
		return
	}
	var flg bool
	for _, el := range data{
		for key,val := range el {
			if strings.Contains(strings.ToLower(val),strings.ToLower(str))||strings.Contains(strings.ToLower(key),strings.ToLower(str)) {
				fmt.Println("Найдено совпадение: ",key,"-",val)
				flg = true
			}
		}
	}
	if !flg {
		fmt.Println("Ничего не найдено")
	}

	search(data)
	return

}