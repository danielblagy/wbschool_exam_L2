package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{
		"тяпка", "пятак", "листок", "пятка", "тяпка", "слиток", "остр", "столик", "пятка", "слиток", "л", "рост", "торс"}

	anagrams := getAnagrams(words)
	fmt.Println(anagrams)
}

func getAnagrams(words []string) map[string][]string {
	encounteredWords := map[string]struct{}{} // to make sure the words in an anagram group are unique
	anagrams := map[string][]string{}

	for _, word := range words {
		// to lower case
		word = strings.ToLower(word)

		// if the word has been encountered before
		if _, ok := encounteredWords[word]; ok {
			continue
		}

		// if the word consists of one letter (or is an empty string)
		runes := RuneSlice(word)
		if len(runes) < 2 {
			continue
		}

		encounteredWords[word] = struct{}{}
		// sort the letters
		sort.Sort(runes)
		anagrams[string(runes)] = append(anagrams[string(runes)], word)
	}

	// sort anagram groups in ascending order
	for _, anagramsGroup := range anagrams {
		sort.Strings(anagramsGroup)
	}

	return anagrams
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
