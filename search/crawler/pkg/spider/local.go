// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

// LocScan Заглушка для тестов.
type LocScan struct {
}

// Scan возвращает ассоциативный массив: ссылок-описание
func (l *LocScan) Scan(Urls []string, depth int) (data map[string]string, err error) {
	data = map[string]string{
		"https://www.mail.ru": "Mail.ru: почта, поиск в интернете, новости, игры",
	}
	return data, nil
}
