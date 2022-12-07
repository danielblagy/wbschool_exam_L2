package main

import (
	"errors"
	"os"
	"testing"
)

func TestWget(t *testing.T) {
	Wget()

	if _, err := os.Stat("www.google.com.html"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Wget returned an unexpected result")
	}
}
