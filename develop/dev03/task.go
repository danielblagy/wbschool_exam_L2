package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
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
	strs := []string{
		"1233 Go 45",
		"b4w Bravo 203",
		"738ks Gopher 15000",
		"048na Alpha 7",
		"sjdhd Grin 88",
		"hello Delta 65",
		"048na Alpha 7",
		"sjdhd Grin 88",
		"hello Delta 65",
	}

	sortedStrs, err := SortStrings(strs, SortOptions{k: 2, n: true, r: true, u: true})
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range sortedStrs {
		fmt.Println(str)
	}
}

type SortOptions struct {
	k int
	n bool
	r bool
	u bool
}

func SortStrings(strs []string, options SortOptions) ([]string, error) {
	// the number of columns each string must have is determined by the first string
	// the validation won't be done on each string to save time, in the event a string
	// in the slice has fewer columns, 'out of bounds' error might be thrown
	strsColLen := len(strings.Split(strs[0], " "))
	if options.k < 0 || options.k >= strsColLen {
		return nil, errors.New("invalid SortOptions: k is out of bounds")
	}

	result := make([]string, 0, len(strs))

	// due to the need of converting string to float64 in the event we sort by numbers
	// the decision has been made to split the implementation
	if options.n {
		result = sortNumbers(strs, result, &options)
	} else {
		result = sortStrings(strs, result, &options)
	}

	if options.u {
		allKeys := make(map[string]struct{})
		uniqueStrs := make([]string, 0, len(result))
		for _, s := range result {
			if _, ok := allKeys[s]; !ok {
				allKeys[s] = struct{}{}
				uniqueStrs = append(uniqueStrs, s)
			}
		}
		result = uniqueStrs
	}

	return result, nil
}

func sortStrings(strs []string, result []string, options *SortOptions) []string {
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

	if options.r {
		sort.Sort(sort.Reverse(sort.StringSlice(result)))
	}

	return result
}

func sortNumbers(strs []string, result []string, options *SortOptions) []string {
	pairs := make(map[float64]int, len(strs))
	column := make([]float64, 0, len(strs))

	for i, str := range strs {
		strCols := strings.Split(str, " ")
		rowNumber, _ := strconv.ParseFloat(strCols[options.k], 64)
		pairs[rowNumber] = i
		column = append(column, rowNumber)
	}

	sort.Float64s(column)

	if options.r {
		sort.Sort(sort.Reverse(sort.Float64Slice(column)))
	}

	for _, row := range column {
		result = append(result, strs[pairs[row]])
	}

	return result
}
