package index

import (
	"regexp"
	"sort"
	"strings"
)

//Storage - структура хранения документа
type Storage struct {
	Index int
	Title string
	URL   string
}

//NewStorage - конструктор
//Возвращает массив отсортированных по индексу документов
func NewStorage(data map[string]string) *[]Storage {
	tmp := []Storage{}
	i := 0
	for key, val := range data {
		tmp = append(tmp, Storage{
			Index: i,
			Title: val,
			URL:   key,
		})
		i++
	}
	sort.Slice(tmp[:], func(i, j int) bool {
		return tmp[i].Index < tmp[j].Index
	})
	return &tmp
}

//Create - создает хранилище и обратный индекс для документов
//
func Create(data map[string]string) (st *[]Storage, index map[string][]int) {
	st = NewStorage(data)
	index = make(map[string][]int)
	for _, val := range *st {
		re := regexp.MustCompile(`[A-Za-z']+|\p{Cyrillic}+|-[*?()$.,!]`)
		words := re.FindAllString(val.Title, -1)
		for _, el := range words {
			index[strings.ToLower(el)] = appendUnic(index[strings.ToLower(el)], val.Index)
		}
	}
	return st, index
}

func appendUnic(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

//BSearch - Выполняет бинарный поиск
//Возвращает нужный элемент
func BSearch(st []Storage, start, end, index int) (result Storage) {
	if start == end {
		if st[start].Index == index {
			result = st[start]
		}
	}
	if start != end {
		center := start + (end-start)/2
		if index <= st[center].Index {
			return BSearch(st, start, center, index)
		}
		if index > st[center].Index {
			return BSearch(st, center+1, end, index)
		}
	}
	return result
}
