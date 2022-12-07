package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	fields    int
	delimiter string
	separated bool
}

func getInput(flags *flags) [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	var words [][]string
	fmt.Println("Enter 'quit' to stop")
	for {
		ok := scanner.Scan()
		if !ok {
			log.Fatal(scanner.Err())
		}

		line := scanner.Text()
		if line == "quit" {
			break
		}

		// is separated flag is true and the line doesn't contain delimiter
		if flags.separated && !strings.Contains(line, flags.delimiter) {
			continue
		}

		words = append(words, strings.Split(line, flags.delimiter))
	}

	return words
}

func Cat(flags *flags, words [][]string) {
	if flags.fields < 0 {
		log.Fatal(errors.New("invalid fields flag"))
	}

	// if flags is not defined
	if flags.fields != 0 {
		var columns []string
		for _, s := range words {
			columns = append(columns, s[flags.fields])
		}
		fmt.Println(columns)
	} else {
		for _, s := range words {
			for _, word := range s {
				fmt.Print(word + flags.delimiter)
			}
			fmt.Println()
		}

	}
}

func main() {
	flags := flags{}

	flag.IntVar(&flags.fields, "f", 1, "выбрать поля(колонки)")
	flag.StringVar(&flags.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&flags.separated, "s", false, "только строки с разделителем")

	flag.Parse()

	words := getInput(&flags)
	Cat(&flags, words)
}
