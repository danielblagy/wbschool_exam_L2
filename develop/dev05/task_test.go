package main

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	phrase := "whacked"
	text := []string{
		"cat whacked the tent",
		"dog stirred the dust",
		"baby cried",
		"the lizard stayed",
	}
	flags := flags{}
	expectedResult := []string{"cat whacked the tent\n"}

	result := Grep(phrase, text, &flags)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Grep returned an unexpected result")
	}
}
