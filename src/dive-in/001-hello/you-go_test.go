package main

import (
	"fmt"
	"testing"
)

func TestYouGo(t *testing.T) {
	expected := "Go You!"
	if got := YouGo(); got != expected {
		t.Errorf("YouGo() returned %s instead of expected %s", got, expected)
	} else {
		fmt.Printf("YouGo() got %s and expected %s\n", got, expected)
	}
}
