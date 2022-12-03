package main

import (
	"reflect"
	"testing"
)

func TestSortStringsKOutOfBounds(t *testing.T) {
	strs := []string{"1 2 3 4 5 6 7 8 9 10 11 12 13 14 15"}

	if _, err := SortStrings(strs, SortOptions{k: 2}); err != nil {
		t.Errorf("unexpected result: k is within bounds")
	}

	if _, err := SortStrings(strs, SortOptions{k: -2}); err == nil {
		t.Errorf("unexpected result: k is out of bounds, but the error is not returned")
	}

	if _, err := SortStrings(strs, SortOptions{k: 15}); err == nil {
		t.Errorf("unexpected result: k is out of bounds, but the error is not returned")
	}
}

func TestSortStringsNumbersReverseUnique(t *testing.T) {
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

	expectedResult := []string{
		"738ks Gopher 15000",
		"b4w Bravo 203",
		"sjdhd Grin 88",
		"hello Delta 65",
		"1233 Go 45",
		"048na Alpha 7",
	}

	sortedStrs, err := SortStrings(strs, SortOptions{k: 2, n: true, r: true, u: true})
	if err != nil {
		t.Errorf("the error should be nil")
	}

	if !reflect.DeepEqual(sortedStrs, expectedResult) {
		t.Errorf("sortedStrs is not the expected result")
	}
}

func TestSortStrings(t *testing.T) {
	strs := []string{
		"b",
		"f",
		"fb",
		"fa",
		"n",
		"a",
		"c3",
		"j0",
		"z",
	}

	expectedResult := []string{
		"a",
		"b",
		"c3",
		"f",
		"fa",
		"fb",
		"j0",
		"n",
		"z",
	}

	sortedStrs, err := SortStrings(strs, SortOptions{k: 0})
	if err != nil {
		t.Errorf("the error should be nil")
	}

	if !reflect.DeepEqual(sortedStrs, expectedResult) {
		t.Errorf("sortedStrs is not the expected result")
	}
}
