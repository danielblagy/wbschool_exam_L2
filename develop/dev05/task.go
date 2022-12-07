package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func Find(strs []string, val string) bool {
	for _, item := range strs {
		if item == val {
			return true
		} else {
			continue
		}
	}
	return false
}

func Grep(phrase string, text []string, flags *flags) []string {
	// for count flag
	var stringCount = 0
	var result []string
	// comparison condition
	var condition bool

	// manage after, before, context flags
	if flags.context != 0 && flags.context > flags.after && flags.context > flags.before {
		flags.after = flags.context
		flags.before = flags.context
	} else if flags.context != 0 && (flags.context > flags.after || flags.context > flags.before) {
		if flags.context > flags.after {
			flags.after = flags.context
		} else if flags.context > flags.before {
			flags.before = flags.context
		}
	}

	for i, str := range text {
		if flags.ignoreCase {
			str = strings.ToLower(str)
			phrase = strings.ToLower(phrase)
		}

		// if exact match
		if flags.fixed {
			condition = phrase == str
		} else {
			condition = strings.Contains(str, phrase)
		}

		if flags.invert {
			condition = !condition
		}

		if condition {
			stringCount++

			if flags.before != 0 {
				if i <= flags.before-1 {
					for j := i; j >= 0; j-- {
						if !Find(result, text[i-j]) {
							if flags.lineNum {
								result = append(result, strconv.Itoa(i-j+1)+" "+text[i-j])
							} else {
								result = append(result, text[i-j])
							}
						} else {
							continue
						}
					}
				} else {
					for j := flags.before; j >= 0; j-- {
						if !Find(result, text[i-j]) {
							if flags.lineNum {
								result = append(result, strconv.Itoa(i-j+1)+" "+text[i-j])
							} else {
								result = append(result, text[i-j])
							}
						} else {
							continue
						}
					}
				}
			}

			if flags.after != 0 {
				if i > len(text)-1-flags.after {
					for j := 0; j < len(text)-i+1; j++ {
						if !Find(result, text[i]) {
							if flags.lineNum {
								result = append(result, strconv.Itoa(i+1)+" "+text[i])
							} else {
								result = append(result, text[i])
							}
							i++
						} else {
							i++
						}
					}
				} else {
					for j := 0; j < flags.after+1; j++ {
						if !Find(result, text[i]) {
							if flags.lineNum {
								result = append(result, strconv.Itoa(i+1)+" "+text[i])
							} else {
								result = append(result, text[i])
							}
							i++
						} else {
							i++
						}

					}
				}
			}

			if flags.after == 0 && flags.before == 0 {
				if flags.lineNum {
					result = append(result, strconv.Itoa(i+1)+" "+text[i])
				} else {
					result = append(result, text[i])
				}
			}
		}
	}

	if flags.count {
		count := strconv.Itoa(stringCount)
		result = []string{count}
	}

	// add a newline to the final result
	finalResult := make([]string, 0, len(result))
	for _, s := range result {
		finalResult = append(finalResult, s+"\n")
	}
	return finalResult
}

func main() {
	// parse flags
	flags := flags{}
	flag.IntVar(&flags.after, "A", 0, `print +N strings after a match`)
	flag.IntVar(&flags.before, "B", 0, `print +N strings before a match`)
	flag.IntVar(&flags.context, "C", 0, "print ±N strings around a match")
	flag.BoolVar(&flags.count, "c", false, "show line count")
	flag.BoolVar(&flags.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&flags.invert, "v", false, "exclude matches")
	flag.BoolVar(&flags.fixed, "F", false, "exact match")
	flag.BoolVar(&flags.lineNum, "n", false, "print line number")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		log.Fatalln("usage: [flags] [pattern or string] [file]")
	}

	phrases := args[:len(args)-1]
	phrase := strings.Join(phrases, " ")

	file, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	splitString := strings.Split(string(file), "\n")

	fmt.Println(Grep(phrase, splitString, &flags))
}
