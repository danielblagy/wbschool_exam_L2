package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	s, err := unpackString("a4bc2d5e")
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString("abcd")
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString("45")
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString("")
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString(`qwe\4\5`)
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString(`qwe\45`)
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString(`qwe\\5`)
	fmt.Printf("\"%s\" %v\n", s, err)

	s, err = unpackString(`\5\ab8`)
	fmt.Printf("\"%s\" %v\n", s, err)
}

func unpackString(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	runes := []rune(s)
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("invalid string: begins with a number")
	}

	builder := strings.Builder{}
	builder.Grow(len(s))

	for i := 1; i < len(runes); i++ {
		if number, err := strconv.Atoi(string(runes[i])); err == nil {
			if runes[i-1] == '\\' {
				//i++
				continue
			}

			for j := 0; j < number; j++ {
				builder.WriteRune(runes[i-1])
			}
			i++
			continue
		}

		builder.WriteRune(runes[i-1])
	}

	if !unicode.IsDigit(runes[len(runes)-1]) || runes[len(runes)-2] == '\\' {
		builder.WriteRune(runes[len(runes)-1])
	}

	return builder.String(), nil
}
