package main

import (
	"fmt"
	"testing"
)

func TestYouGo(t *testing.T) {
	const expected = "Go You!"
	if got := youGo(); got != expected {
		t.Errorf("youGo() returned %s instead of expected %s", got, expected)
	} else {
		fmt.Printf("youGo() got %s and expected %s\n", got, expected)
	}
}
