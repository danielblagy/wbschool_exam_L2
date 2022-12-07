package main

import "testing"

func TestCat(t *testing.T) {
	flags := flags{}
	words := [][]string{{"hello", "cat"}, {"how u", "doin"}}

	Cat(&flags, words)
}
