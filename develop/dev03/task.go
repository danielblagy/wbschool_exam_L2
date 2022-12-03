package main

import (
	"errors"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

}

type SortOptions struct {
	k int
	n bool
	r bool
	u bool
}

func SortStrings(strs []string, options SortOptions) ([]string, error) {
	// the number of columns each string must have is determined by the first string
	strsColLen := len(strings.Split(strs[0], " "))
	if options.k < 0 || options.k >= strsColLen {
		return nil, errors.New("invalid SortOptions: k is out of bounds")
	}

	result := make([]string, 0, len(strs))

	pairs := make(map[string]int, len(strs))
	column := make([]string, 0, len(strs))

	for i, str := range strs {
		strCols := strings.Split(str, " ")
		pairs[strCols[options.k]] = i
		column = append(column, strCols[options.k])
	}

	sort.Strings(column)

	for _, row := range column {
		result = append(result, strs[pairs[row]])
	}

	return result, nil
}
