package main

import (
	"reflect"
	"testing"
)

func TestGetAnagrams(t *testing.T) {
	words := []string{
		"тяпка", "пятак", "листок", "пятка", "тяпка", "слиток", "остр", "столик", "пятка", "слиток", "л", "рост", "торс"}

	expectedResult := map[string][]string{
		"акптя":  []string{"пятак", "пятка", "тяпка"},
		"иклост": []string{"листок", "слиток", "столик"},
		"орст":   []string{"остр", "рост", "торс"},
	}

	result := getAnagrams(words)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("getAnagrams returned an unexpected result")
	}
}
